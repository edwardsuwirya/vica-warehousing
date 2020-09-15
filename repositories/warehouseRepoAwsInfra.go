package repositories

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"warehousing/models"
)

type WarehouseRepositoryAwsInfrastructure struct {
	bucketName string
	dataPath   string
	sess       *session.Session
	fileRepo   IWarehouseFileRepo
}

func NewWarehouseRepoAwsInfra(bucketName, dataPath string, aws *session.Session) IWarehouseFileRepo {
	return &WarehouseRepositoryAwsInfrastructure{
		bucketName, dataPath, aws, NewWarehouseRepoInfra(dataPath),
	}
}

func (bri *WarehouseRepositoryAwsInfrastructure) SaveToFile(warehouseCollection *models.Warehouse) {
	bri.fileRepo.SaveToFile(warehouseCollection)
	file, err := os.OpenFile(bri.dataPath, os.O_RDONLY, 0644)
	uploader := s3manager.NewUploader(bri.sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bri.bucketName),
		ACL:    aws.String("public-read"),
		Key:    aws.String("test-golang.txt"),
		Body:   file,
	})
	if err != nil {
		panic(err)
	}
}

func (bri *WarehouseRepositoryAwsInfrastructure) ReadFile() []*models.Warehouse {
	downloader := s3manager.NewDownloader(bri.sess)
	_, err := os.Stat(bri.dataPath)
	var file *os.File
	defer file.Close()
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(bri.dataPath)
			if err != nil {
				panic(err)
			}
		}
	} else {
		file, _ = os.OpenFile(bri.dataPath, os.O_RDWR, 0644)
	}
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bri.bucketName),
			Key:    aws.String("test-golang.txt"),
		})
	if err != nil {
		fmt.Println(err)
	}
	return bri.fileRepo.ReadFile()
}
