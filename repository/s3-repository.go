package repository

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"socket-storage/interfaces"
	file_stream "socket-storage/py-rpc/proto"
	"time"
)

type StorageS3Repository struct {
	bucket string
	svc    *s3.S3
	sess   *session.Session
	_rpc *grpc.ClientConn
}

func (s *StorageS3Repository) PutFileStreamProto(request *file_stream.InputFrame) (*file_stream.OutputFrame, error) {
	logger := logrus.WithFields(logrus.Fields{
		"domain":     "s3-socket",
		"action":     "Rpc response",
		"repository": "PutFileStreamProto",
	})

	//ctx := metadata.AppendToOutgoingContext(context.Background(), "users", "tests")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := file_stream.NewStreamInputClient(s._rpc)
	fileOutput, err := service.ConvertDataframe(ctx, request)
	if err != nil {
		logger.WithField("type", "gRPC ConvertDataframe").Errorln(err)
		return nil, err
	}

	return fileOutput, nil
}

func (s *StorageS3Repository) HashFileMD5(fileReader *bytes.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, fileReader); err != nil {
		return "", err
	}

	inByte := hash.Sum(nil)[:16]
	return hex.EncodeToString(inByte), nil
}

func (s *StorageS3Repository) PutObject(fileReader *bytes.Reader, message []byte, fileName, fileExt, userID string) (*file_stream.OutputFrame, error) {
	logger := logrus.WithFields(logrus.Fields{
		"domain":     "s3-socket",
		"action":     "upload file",
		"repository": "PutObject",
	})

	inputFrame := &file_stream.InputFrame{
		Data: message,
		FileName: fileName,
		FileType: fileExt,
		UserId: userID,
	}

	response, err := s.PutFileStreamProto(inputFrame)
	if err != nil {
		logger.WithField("type", "PutFileStreamProto repository").Errorln(err)
		return nil, err
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
		logger.WithField("type", "S3 PutObject").Errorln(err)
		return nil, err
	}

	return response, nil
}

func NewStorageS3Repo(bucket string, svc *s3.S3, session *session.Session, rpc *grpc.ClientConn) interfaces.StorageS3Repository {
	return &StorageS3Repository{bucket: bucket, svc: svc, sess: session, _rpc: rpc}
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
