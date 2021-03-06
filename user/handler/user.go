package handler

import (
	"bytes"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/sirupsen/logrus"
	pbbottle "github.com/vegchic/fullbottle/bottle/proto/bottle"
	"github.com/vegchic/fullbottle/common"
	"github.com/vegchic/fullbottle/common/db"
	"github.com/vegchic/fullbottle/common/kv"
	"github.com/vegchic/fullbottle/common/log"
	"github.com/vegchic/fullbottle/config"
	"github.com/vegchic/fullbottle/user/dao"
	pb "github.com/vegchic/fullbottle/user/proto/user"
	"github.com/vegchic/fullbottle/util"
	"github.com/vegchic/fullbottle/weed"
	"io"
	"time"
)

const UserInfoKey = "info:usr_id=%d"
const UserEmailLockKey = "lock:email=%s"

type UserHandler struct{}

func (u *UserHandler) GetUserInfo(ctx context.Context, req *pb.GetUserRequest, resp *pb.GetUserResponse) error {
	uid := req.GetUid()
	user := &dao.User{}
	key := fmt.Sprintf(UserInfoKey, uid)
	if err := kv.GetM(key, user); err != nil {
		user, err = dao.GetUsersById(uid)
		if err != nil {
			return err
		} else if user == nil {
			return errors.New(config.UserSrvName, "User not found", common.NotFoundError)
		}
		_ = kv.SetM(key, user, 24*time.Hour)
	}

	resp.Uid = user.ID
	resp.Username = user.Username
	resp.Email = user.Email
	resp.Role = user.Role
	resp.AvatarFid = user.AvatarFid
	resp.Status = user.Status
	resp.CreateTime = user.CreateTime.Unix()
	resp.UpdateTime = user.UpdateTime.Unix()
	if user.DeleteTime != nil {
		resp.DeleteTime = user.DeleteTime.Unix()
	}

	return nil
}

func (u *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest, resp *pb.CreateUserResponse) error {
	email := req.GetEmail()

	lock, err := kv.Obtain(fmt.Sprintf(UserEmailLockKey, email), 100*time.Millisecond)
	if err != nil {
		return err
	}
	defer lock.Release()

	if u, err := dao.GetUsersByEmail(email); err != nil {
		return err
	} else if u != nil {
		return errors.New(config.UserSrvName, "Email existed", common.ExistedError)
	}

	user := &dao.User{
		Email:    req.Email,
		Username: req.Username,
		Password: util.Bcrypt(req.Password),
	}

	err = dao.CreateUser(user)
	if err != nil {
		return err
	}

	bottleClient := common.BottleSrvClient()
	_, err = bottleClient.InitBottle(ctx, &pbbottle.InitBottleRequest{Uid: user.ID, Capacity: config.DefaultCapacity})
	if err != nil {
		log.WithError(err).Errorf("Cannot init user bottle")
	}

	_ = kv.GetM(fmt.Sprintf(UserInfoKey, user.ID), user)
	return nil
}

func (u *UserHandler) ModifyUser(ctx context.Context, req *pb.ModifyUserRequest, resp *pb.ModifyUserResponse) error {
	uid := req.GetUid()

	user, err := dao.GetUsersById(uid)
	if err != nil {
		return err
	} else if user == nil {
		return errors.New(config.UserSrvName, "User not found", common.NotFoundError)
	}

	fields := db.Fields{}
	if req.Username != "" {
		fields["username"] = req.GetUsername()
	}
	if req.Password != "" {
		fields["password"] = util.Bcrypt(req.GetPassword())
	}

	if err = dao.UpdateUser(user, fields); err != nil {
		return err
	}
	_ = kv.Del(fmt.Sprintf(UserInfoKey, uid))
	return nil
}

func (u *UserHandler) UserLogin(ctx context.Context, req *pb.UserLoginRequest, resp *pb.UserLoginResponse) error {
	email := req.GetEmail()
	user, err := dao.GetUsersByEmail(email)
	if err != nil {
		return err
	} else if user == nil {
		return errors.New(config.UserSrvName, "User not found", common.NotFoundError)
	}

	if pass := util.BcryptCompare(user.Password, req.Password); !pass {
		return errors.New(config.UserSrvName, "Password error", common.PasswordError)
	}

	expire := time.Now().Unix() + config.JwtTokenExpire

	ip, _ := metadata.Get(ctx, "ip")
	token, err := util.GenerateJwtToken(user.ID, expire, ip)
	if err != nil {
		return errors.New(config.UserSrvName, err.Error(), common.JwtError)
	}

	resp.Token, resp.Expire, resp.Uid = token, expire, user.ID
	return nil
}

func (u *UserHandler) GetUserAvatar(ctx context.Context, req *pb.GetUserAvatarRequest, resp *pb.GetUserAvatarResponse) error {
	uid := req.GetUid()
	user := &dao.User{}
	key := fmt.Sprintf(UserInfoKey, uid)
	if err := kv.GetM(key, user); err != nil {
		user, err = dao.GetUsersById(uid)
		if err != nil {
			return err
		} else if user == nil {
			return errors.New(config.UserSrvName, "User not found", common.NotFoundError)
		}
		_ = kv.SetM(key, user, 24*time.Hour)
	}

	avatarFid := user.AvatarFid
	if avatarFid == "" {
		return errors.New(config.UserSrvName, "User has no avatar", common.EmptyAvatarError)
	}

	avatarResp, err := weed.FetchFile(avatarFid)
	if err != nil {
		return err
	}

	body := avatarResp.Body
	defer body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, body); err != nil {
		return errors.New(config.UserSrvName, "Avatar lost due to: "+err.Error(), common.FileFetchError)
	}

	resp.Avatar = buf.Bytes()
	resp.ContentType = avatarResp.Header.Get("Content-Type")
	return nil
}

func (u *UserHandler) UploadUserAvatar(ctx context.Context, req *pb.UploadUserAvatarRequest, resp *pb.UploadUserAvatarResponse) error {
	user, err := dao.GetUsersById(req.GetUid())
	if err != nil {
		return err
	} else if user == nil {
		return errors.New(config.UserSrvName, "User not found", common.NotFoundError)
	}

	var volume *weed.VolumeLookupInfo
	var volumeUrl string
	fid := user.AvatarFid
	// if already exist, then try to rewrite
	if user.AvatarFid != "" {
		f, err := weed.ParseFid(user.AvatarFid)
		if err == nil {
			volume, err = weed.LookupVolume(f.VolumeId)
		}
		if err != nil {
			fid = "" // reset fid
			log.WithFields(logrus.Fields{
				"userId": user.ID,
			}).WithError(errors.New(config.UserSrvName, "Dirty avatar fid", common.InternalError))
		}
	}

	// if failed, get new file key
	if fid != "" {
		volumeUrl = volume.Locations[0].Url
	} else {
		key, err := weed.AssignFileKey()
		if err != nil {
			return err
		}
		fid = key.Fid
		volumeUrl = key.Url
	}

	// write db first
	if fid != user.AvatarFid {
		err = dao.UpdateUser(user, db.Fields{
			"avatar_fid": fid,
		})
		if err != nil {
			return err
		}
	}

	_, err = weed.UploadSingleFile(bytes.NewReader(req.Avatar), fmt.Sprint(user.ID, "-", time.Now().Unix()),
		fid, volumeUrl, false)
	_ = kv.Del(fmt.Sprintf(UserInfoKey, req.GetUid()))
	return err
}
