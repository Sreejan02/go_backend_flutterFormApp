package controllers

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"backend/database"
	"backend/models"
	"backend/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)


// ===============================
// UPLOAD FILE + SAVE IN DB
// POST /files/upload
// ===============================
func UploadFile(c *gin.Context) {

	// Get file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "No file uploaded",
		})
		return
	}

	// Get metadata
	personalIDStr := c.PostForm("personal_id")
	fileType := c.PostForm("type")

	if personalIDStr == "" || fileType == "" {
		c.JSON(400, gin.H{
			"error": "personal_id and type required",
		})
		return
	}

	personalID, err := strconv.ParseUint(personalIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid personal_id",
		})
		return
	}

	// Generate unique file name
	ext := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + ext

	// Open file
	src, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	ctx := context.Background()

	// Upload to MinIO
	_, err = storage.MinioClient.PutObject(
		ctx,
		storage.BucketName,
		fileName,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Build URL
	url := fmt.Sprintf(
		"http://localhost:9000/%s/%s",
		storage.BucketName,
		fileName,
	)

	// Save in DB
	record := models.UploadedFile{
		PersonalID:   uint(personalID),
		Type:         fileType,
		FileName:     fileName,
		FileURL:      url,
		OriginalName: file.Filename,
		Size:         file.Size,
		CreatedAt:    time.Now(),
	}

	if err := database.DB.Create(&record).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "DB save failed: " + err.Error(),
		})
		return
	}

	// Response
	c.JSON(200, gin.H{
		"id":          record.ID,
		"personal_id": record.PersonalID,
		"type":        record.Type,
		"url":         record.FileURL,
		"uploaded":    record.CreatedAt,
	})
}


// ===============================
// GET FILE (STREAM FROM MINIO)
// GET /files/:name
// ===============================
func GetFile(c *gin.Context) {

	fileName := c.Param("id")

	ctx := context.Background()

	obj, err := storage.MinioClient.GetObject(
		ctx,
		storage.BucketName,
		fileName,
		minio.GetObjectOptions{},
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	stat, err := obj.Stat()
	if err != nil {
		c.JSON(404, gin.H{
			"error": "File not found",
		})
		return
	}

	c.Header("Content-Disposition", "inline; filename="+stat.Key)
	c.Header("Content-Type", stat.ContentType)

	c.DataFromReader(
		http.StatusOK,
		stat.Size,
		stat.ContentType,
		obj,
		nil,
	)
}


// ===============================
// LIST USER FILES
// GET /files/user/:id
// ===============================
func GetUserFiles(c *gin.Context) {

	id := c.Param("id")

	var files []models.UploadedFile

	if err := database.DB.
		Where("personal_id = ?", id).
		Find(&files).Error; err != nil {

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, files)
}


// ===============================
// DELETE FILE (MINIO + DB)
// DELETE /files/:id
// ===============================
func DeleteFile(c *gin.Context) {

	id := c.Param("id")

	var file models.UploadedFile

	if err := database.DB.First(&file, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	ctx := context.Background()

	// Delete from MinIO
	_ = storage.MinioClient.RemoveObject(
		ctx,
		storage.BucketName,
		file.FileName,
		minio.RemoveObjectOptions{},
	)

	// Delete from DB
	database.DB.Delete(&file)

	c.JSON(200, gin.H{
		"message": "File deleted",
	})
}


// ===============================
// UPDATE FILE (REPLACE)
// PUT /files/:id
// ===============================
func UpdateFile(c *gin.Context) {

	id := c.Param("id")

	var old models.UploadedFile

	if err := database.DB.First(&old, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file uploaded"})
		return
	}

	ctx := context.Background()

	// Delete old file
	_ = storage.MinioClient.RemoveObject(
		ctx,
		storage.BucketName,
		old.FileName,
		minio.RemoveObjectOptions{},
	)

	// Upload new
	src, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	newName := uuid.New().String() + ext

	_, err = storage.MinioClient.PutObject(
		ctx,
		storage.BucketName,
		newName,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	newURL := fmt.Sprintf(
		"http://localhost:9000/%s/%s",
		storage.BucketName,
		newName,
	)

	// Update DB
	old.FileName = newName
	old.FileURL = newURL
	old.OriginalName = file.Filename
	old.Size = file.Size

	database.DB.Save(&old)

	c.JSON(200, gin.H{
		"message": "Updated",
		"url":     newURL,
	})
}
