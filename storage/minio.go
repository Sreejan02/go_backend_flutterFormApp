package storage

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var BucketName string

func InitMinio() {

	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	BucketName = os.Getenv("MINIO_BUCKET")

	useSSL := false

	client, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatal("❌ MinIO connect failed:", err)
	}

	MinioClient = client

	ctx := context.Background()

	exists, err := client.BucketExists(ctx, BucketName)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		err = client.MakeBucket(ctx, BucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatal("❌ Cannot create bucket:", err)
		}

		log.Println("✅ Created bucket:", BucketName)
	}

	log.Println("✅ MinIO Connected Successfully")
}
