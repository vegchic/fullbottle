package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/vegchic/fullbottle/common"
	"github.com/vegchic/fullbottle/common/db"
	"github.com/vegchic/fullbottle/common/log"
)

type FileMeta struct {
	db.BasicModel
	FId  string `gorm:"type:varchar(64);not null"`
	Size int64  `gorm:"type:bigint;not null"`
	Hash string `gorm:"type:varchar(128);not null"`
}

func (FileMeta) TableName() string {
	return "file_meta"
}

type FileInfo struct {
	db.BasicModel
	Name     string   `gorm:"type:varchar(128);not null"`
	FileId   int64    `grom:"not null"`
	OwnerId  int64    `gorm:"not null"`
	FolderId int64    `gorm:"not null"`
	Metadata FileMeta `gorm:"foreignkey:FileId;save_associations:false;preload:false"`
}

func (FileInfo) TableName() string {
	return "file_info"
}

func GetFileById(ownerId, id int64) (*FileInfo, error) {
	var file FileInfo
	if err := db.DB().Where("id = ? AND owner_id = ? AND status = ?", id, ownerId, db.Valid).
		Preload("Metadata").First(&file).Error; err != nil {

		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		log.WithError(err).Errorf("DB error")
		return nil, common.NewDBError(err)
	}
	return &file, nil
}

func GetFilesByFolderId(ownerId, folderId int64) ([]*FileInfo, error) {
	// TODO check result
	var files []*FileInfo
	if err := db.DB().Where("folder_id = ? AND owner_id = ? AND status = ?", folderId, ownerId, db.Valid).
		Preload("Metadata").Find(&files).Error; err != nil {

		log.WithError(err).Errorf("DB error")
		return nil, common.NewDBError(err)
	}
	return files, nil
}

func GetFilesByFolderIds(ownerId int64, parentIds []int64) ([]*FileInfo, error) {
	var files []*FileInfo
	if err := db.DB().Where("folder_id in (?) AND owner_id = ? AND status = ?", parentIds, ownerId, db.Valid).
		Preload("Metadata").Find(&files).Error; err != nil {

		log.WithError(err).Errorf("DB error")
		return nil, common.NewDBError(err)
	}
	return files, nil
}

func UpdateFiles(file *FileInfo) error {
	if err := db.DB().Table("file_info").Updates(file).Error; err != nil {
		log.WithError(err).Errorf("DB error")
		return common.NewDBError(err)
	}
	return nil
}