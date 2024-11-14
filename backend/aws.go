package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/cloudflare/cloudflare-go"
)

var svc *s3.S3
var sess *session.Session
var cfAPI *cloudflare.API

func initAWS() {

	var err error

	sess, err = session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("CLOUDFLARE_ACCESS_KEY"),
			os.Getenv("CLOUDFLARE_SECRET_KEY"),
			"", // token can be left empty for non-temporary credentials
		),
		Endpoint: aws.String(os.Getenv("CLOUDFLARE_S3_URL")),
		Region:   aws.String("auto"),
	})

	if err != nil {
		log.Fatal(err)
	}

	svc = s3.New(sess)

	cfAPI, err = cloudflare.New(os.Getenv("CLOUDFLARE_GLOBAL_API_KEY"), os.Getenv("CLOUDFLARE_GLOBAL_API_EMAIL"))
	// alternatively, you can use a scoped API token
	// cfAPI, err = cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

}

// UploadTextFile writes text extraction to a text file in the specified bucket and key
func UploadTextFile(text string, bucket string, key string) error {
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(text),
	}

	_, err := svc.PutObject(input)
	if err != nil {
		return fmt.Errorf("error uploading extraction file: %s\n%v", key, err)
	}

	return nil
}

func GetObjectFromS3(bucket string, key string) (*s3.GetObjectOutput, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	return svc.GetObject(input)
}

func createCloudflareR2Bucket(bucketName string) error {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return fmt.Errorf("CLOUDFLARE_ACCOUNT_ID environment variable is not set")
	}

	_, err := cfAPI.CreateR2Bucket(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.CreateR2BucketParameters{
		Name: bucketName,
	})

	if err != nil {
		return fmt.Errorf("error creating Cloudflare R2 bucket: %v", err)
	}

	return nil
}
