package dataaccess

import (
	"context"
	"log"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3service struct {
	BucketName string // name of the S3 Bucket
	Region string // aws region e.g us-east-1
	S3Client s3.Client // S3 Client 
}

// Creates a s3Service capable of reading and writing
// to AWS S3
// Inputs:
//     bucketName - name of the s3 bucket to be used for operations
//     region - the aws region where the bucket is located
// Output:
//     If success, a ReaderWriter containing an s3.Client
//     is returned, otherwise, nil and an error
func NewS3(bucketName string, region string) (ReaderWriter, error) {
	s3svc := createS3(region)

	return s3service{
			BucketName: bucketName,
			Region: region,
			S3Client: *s3svc,
		}, nil
}

// Creates a s3Service capable of reading from AWS S3
// Inputs:
//     bucketName - name of the s3 bucket to be used for operations
//     region - the aws region where the bucket is located
// Output:
//     If success, a Reader containing an s3.Client
//     is returned, otherwise, nil and an error
func NewS3Reader(bucketName string, region string) (Reader, error) {
	s3svc := createS3(region)

	return s3service{
			BucketName: bucketName,
			Region: region,
			S3Client: *s3svc,
		}, nil
}

// Creates a s3Service capable of writing to AWS S3
// Inputs:
//     bucketName - name of the s3 bucket to be used for operations
//     region - the aws region where the bucket is located
// Output:
//     If success, a Writer containing an s3.Client
//     is returned, otherwise, nil and an error
func NewS3Writer(bucketName string, region string) (Writer, error) {
	s3svc := createS3(region)

	return s3service{
		BucketName: bucketName,
		Region: region,
		S3Client: *s3svc,
	}, nil
}

func (r s3service) Read() (*Employee, error) {
	return nil, nil
}

func (r s3service) ReadAll() ([]Employee, error) {
	// Just list all buckets for now
	input := &s3.ListBucketsInput{}
	output, err := r.S3Client.ListBuckets(context.TODO(), input)

	if err != nil {
		log.Println("Could not retrieve buckets", err)
		return nil, err
	}

	var employees []Employee
	for _, bucket := range output.Buckets {
		if *bucket.Name != "" {
			employees = append(employees, Employee{
				Name: *bucket.Name,
			})
		}
	}
	return employees, nil
}

func (r s3service) Write() error {
	return nil
}


func createS3(region string) *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	return s3.NewFromConfig(cfg)
}