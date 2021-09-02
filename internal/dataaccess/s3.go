package dataaccess

import (
	"fmt"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3service struct {
	// things needed to make S3DataReader functional
	BucketName string
	Region string
	Session session.Session
	S3Service s3.S3
}

func NewS3(bucketName string, region string) (ReaderWriter, error) {
	session, err := createSession(region)

	if err != nil {
		fmt.Println("Could not create aws session", err)
		return nil, err
	}

	creds, err := session.Config.Credentials.Get()

	if err != nil {
		log.Println("Could not retrieve credentials", err)
	}

	log.Println("Using credentials for", creds.AccessKeyID)

	s3svc := createS3(session)

	return config{
		ReaderWriter: s3service{
			BucketName: bucketName,
			Region: region,
			Session: *session,
			S3Service: *s3svc,
		},
	}, nil
}

func NewS3Reader(bucketName string, region string) (Reader, error) {
	session, err := createSession(region)

	if err != nil {
		fmt.Println("Could not create aws session", err)
		return nil, err
	}

	s3svc := createS3(session)

	return s3service{
			BucketName: bucketName,
			Region: region,
			Session: *session,
			S3Service: *s3svc,
		}, nil
}

func NewS3Writer(bucketName string, region string) (Writer, error) {
	session, err := createSession(region)
	
	if err != nil {
		fmt.Println("Could not create aws session", err)
		return nil, err
	}

	s3svc := createS3(session)

	return s3service{
		BucketName: bucketName,
		Region: region,
		Session: *session,
		S3Service: *s3svc,
	}, nil
}

func (r s3service) Read() (*Employee, error) {
	return nil, nil
}

func (r s3service) ReadAll() ([]Employee, error) {
	// Just list all buckets for now
	output, err := r.S3Service.ListBuckets(nil)

	if err != nil {
		log.Println("Could not retrieve buckets", err)
		return nil, err
	}

	log.Println("Buckets: ", output.Buckets)

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

func createSession(region string) (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
}

func createS3(session *session.Session) *s3.S3 {
	return s3.New(session)
}