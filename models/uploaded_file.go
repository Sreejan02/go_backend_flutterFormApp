package models

import "time"

type UploadedFile struct {
    ID uint `gorm:"primaryKey"`

    CreatedAt time.Time

    PersonalID uint `gorm:"index"`

    Type string

    FileName string
    FileURL  string

    OriginalName string
    Size         int64
}
