package dataaccess

type s3 struct {
	// things needed to make S3DataReader functional
	BucketName string
	Region string
	Session string
}

func NewS3() (ReaderWriter, error) {
	return config{
		ReaderWriter: s3{
			BucketName: "test",
			Region: "us-east-1",
			Session: "session",
		},
	}, nil
}

func NewS3Reader() (Reader, error) {
	return s3{
			BucketName: "test",
			Region: "us-east-1",
			Session: "Test",
		}, nil
}

func NewS3Writer() (Writer, error) {
	return s3{
		BucketName: "test",
		Region: "us-east-1",
		Session: "Test",
	}, nil
}


func (r s3) Read() (*Employee, error) {
	return nil, nil
}

func (r s3) ReadAll() ([]Employee, error) {
	return nil, nil
}

func (r s3) Write() error {
	return nil
}