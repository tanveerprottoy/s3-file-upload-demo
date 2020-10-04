package app

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	// Uploader declaration
	Uploader *s3manager.Uploader
	// Session declaration
	Session *session.Session
	region  string
)

// InitS3 to use globally
func InitS3() {
	region = os.Getenv("AWS_REGION")
	Session = session.Must(
		session.NewSession(
			&aws.Config{
				Region: aws.String(region),
			},
		),
	)
	// Create an uploader with the session and default options
	Uploader = s3manager.NewUploader(Session)
}

// GetBucket determines whether we have this bucket
func GetBucket(bucket string) error {
	// Create S3 service client
	svc := s3.New(Session)

	// Do we have this Bucket?
	_, err := svc.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}
	return nil
}

// CreateBucket creates a bucket
func CreateBucket(bucket string) error {
	// Create S3 service client
	svc := s3.New(Session)
	// Create the S3 Bucket
	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}
	// Wait until bucket is created before finishing
	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}
	return nil
}

// UploadObject uploads to s3
func UploadObject(
	bucket string,
	fileName string,
	file multipart.File,
) (string, error) {
	defer file.Close()
	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename.
	_, err := Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),

		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(fileName),

		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body: file,
	})
	if err != nil {
		// Print the error and exit.
		log.Println("unable to upload")
		return "", err
	}
	fmt.Printf("Successfully uploaded %q to %q\n", fileName, bucket)
	// https://<region>.amazonaws.com/<bucket-name>/<key>
	url := fmt.Sprintf(
		"https://%s.amazonaws.com/%s/%s",
		region,
		bucket,
		fileName,
	)
	return url, nil
}
