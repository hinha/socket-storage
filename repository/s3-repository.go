package repository

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/imports"
	"io"
	"net/http"
	"socket-storage/interfaces"
	"strings"
)

type StorageS3Repository struct {
	bucket string
	svc    *s3.S3
	sess   *session.Session
}

func (s *StorageS3Repository) HashFileMD5(fileReader *bytes.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, fileReader); err != nil {
		return "", err
	}

	inByte := hash.Sum(nil)[:16]
	return hex.EncodeToString(inByte), nil
}

func (s *StorageS3Repository) PutObject(fileReader *bytes.Reader, message []byte, fileName, fileExt string) error {
	ctx := context.Background()
	var (
		rows int
		cols int
	)
	//fmt.Println(message)
	hashing, err := s.HashFileMD5(fileReader)
	if err != nil {
		return err
	}

	fmt.Println(hashing)

	if fileExt == "csv" {
		// load frame must bytes
		df, err := imports.LoadFromCSV(ctx, strings.NewReader(string(message)))
		if err != nil {
			return err
		}
		rows = df.NRows()

		// Don't apply read lock because we are write locking from outside.
		iterator := df.ValuesIterator(dataframe.ValuesOptions{Step: 1, DontReadLock: true})
		df.Lock()
		for {
			row, vals, _ := iterator()
			if row == nil {
				break
			}
			cols = len(vals) / 2
		}
		df.Unlock()

		fmt.Println(rows, cols)
	} else {
		//_, err := imports.LoadFromCSV(ctx, strings.NewReader(string(message)))
		//fmt.Println(err)
		//if err != nil {
		//	return err
		//}
	}

	input := &s3.PutObjectInput{
		Bucket:      &s.bucket,
		Key:         &fileName,
		ACL:         aws.String("private"),
		Body:        fileReader,
		ContentType: aws.String(http.DetectContentType(message)),
	}

	_, err = s.svc.PutObject(input)
	if err != nil {
		return err
	}

	return nil
}

func NewStorageS3Repo(bucket string, svc *s3.S3, session *session.Session) interfaces.StorageS3Repository {
	return &StorageS3Repository{bucket: bucket, svc: svc, sess: session}
}

//func backup() {
//	//fmt.Println(df)
//	filterFn := dataframe.FilterDataFrameFn( func(vals map[interface{}]interface{}, row int, nRows int) (dataframe.FilterAction, error) {
//		for _, v := range vals {
//			if v == "NA" {
//				return dataframe.DROP, nil
//			}
//		}
//		return dataframe.KEEP, nil
//	})
//	dtFiltered, err := dataframe.Filter(ctx, df, filterFn)
//	if err != nil {
//		return err
//	}
//	getVals := dataframe.FilterDataFrameFn(func(vals map[interface{}]interface{}, row, nRows int) (dataframe.FilterAction, error) {
//		cols = len(vals) / 2
//		rows = nRows
//		return dataframe.KEEP, nil
//	})
//
//	dataframe.Filter(ctx, dtFiltered, getVals)
//}
