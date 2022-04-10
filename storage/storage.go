package storage

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Binaretech/classroom-main/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

const (
	USERS_BUCKET = "user"
)

var instance *s3.S3

func OpenStorage() {
	sess := s3.New(session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			viper.GetString("s3_access_key"),
			viper.GetString("s3_secret_key"),
			"",
		),
		Endpoint:         aws.String(viper.GetString("s3_endpoint")),
		Region:           aws.String(viper.GetString("s3_region")),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	})))

	out, _ := sess.ListBuckets(&s3.ListBucketsInput{})

	diff := utils.ArrayMissing(utils.ArrayMap(out.Buckets, func(bucket *s3.Bucket) string {
		return *bucket.Name
	}), buckets())

	for _, bucket := range diff {
		sess.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucket),
		})

		publicPolicy, _ := json.Marshal(map[string]any{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				{
					"Sid":       "AddPerm",
					"Effect":    "Allow",
					"Principal": "*",
					"Action": []string{
						"s3:GetObject",
					},
					"Resource": []string{
						fmt.Sprintf("arn:aws:s3:::%s/*", bucket),
					},
				},
			},
		})

		sess.PutBucketPolicy(&s3.PutBucketPolicyInput{
			Bucket: aws.String(bucket),
			Policy: aws.String(string(publicPolicy)),
		})
	}

	instance = sess
}

func buckets() []string {
	return []string{
		USERS_BUCKET,
	}
}

func Put(bucket, key string, data []byte, contentType string) error {
	_, err := instance.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: &contentType,
	})

	return err
}

func Get(bucket, key string) ([]byte, error) {
	resp, err := instance.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.Bytes(), nil
}
