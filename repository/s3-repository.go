package repository

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"net/http"
	"socket-storage/interfaces"
	logger "socket-storage/log"
	"socket-storage/models"
	file_stream "socket-storage/py-rpc/proto"
	"socket-storage/utils"
	"socket-storage/vo"
	"strconv"
)

type StorageS3Repository struct {
	bucket      string
	svc         *s3.S3
	sess        *session.Session
	_rpc        *grpc.ClientConn
	persistence interfaces.StorageS3Persistence
}

func (s *StorageS3Repository) PutFileStreamProto(request *file_stream.InputFrame) (*file_stream.OutputFrame, error) {
	log := logger.Log.WithFields(logrus.Fields{
		"domain":     "s3-socket",
		"action":     "Rpc response",
		"repository": "PutFileStreamProto",
	})

	ctx := metadata.AppendToOutgoingContext(context.Background(), "users", "tests")
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	service := file_stream.NewStreamInputClient(s._rpc)
	fileOutput, err := service.ConvertDataframe(ctx, request)
	if err != nil {
		sentry.CaptureException(err)
		log.WithField("type", "gRPC ConvertDataframe").Errorln(err)
		return nil, err
	}

	return fileOutput, nil
}

func (s *StorageS3Repository) ResponseFileStreamProto(message []byte, fileName, fileExt string, userID int) (string, *file_stream.OutputFrame, error) {
	log := logger.Log.WithFields(logrus.Fields{
		"domain":     "s3-socket",
		"action":     "response file",
		"repository": "ResponseFileStreamProto",
	})

	var fileBytes vo.UploadByteFile // validate bytes data this only bytes!
	if err := json.Unmarshal(message, &fileBytes); err != nil {
		sentry.CaptureException(err)
		log.WithField("type", "json UploadByteFile").Errorln(err)
		return "", nil, err
	}

	inputFrame := &file_stream.InputFrame{
		UserId:   strconv.Itoa(userID),
		Data:     fileBytes.Result,
		FileName: fileName,
		FileType: fileExt,
	}

	response, err := s.PutFileStreamProto(inputFrame)
	if err != nil {
		sentry.CaptureException(err)
		log.WithField("type", "ResponseFileStreamProto repository").Errorln(err)
		return "", nil, err
	}

	bodySize := bytes.NewReader(fileBytes.Result).Size()

	return utils.ByteCountSI(bodySize), response, nil
}

func (s *StorageS3Repository) HashFileMD5(fileReader *bytes.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, fileReader); err != nil {
		return "", err
	}

	inByte := hash.Sum(nil)[:16]
	return hex.EncodeToString(inByte), nil
}

func (s *StorageS3Repository) PutObject(fileReader *bytes.Reader, message []byte, fileName, FileEncrypt string) error {
	log := logger.Log.WithFields(logrus.Fields{
		"domain":     "s3-socket",
		"action":     "upload file",
		"repository": "PutObject",
	})

	input := &s3.PutObjectInput{
		Bucket:      &s.bucket,
		Key:         &fileName,
		ACL:         aws.String("private"),
		Body:        fileReader,
		ContentType: aws.String(http.DetectContentType(message)),
	}

	_, err := s.svc.PutObject(input)
	if err != nil {
		sentry.CaptureException(err)
		log.WithField("type", "S3 PutObject").Errorln(err)
		return err
	}

	log.WithField("type", "Success Upload Object").Info(FileEncrypt)
	return nil
}

func (s *StorageS3Repository) FilterDuplicatesFile(FileEncrypt string, UserID int) error {
	log := logger.Log.WithFields(logrus.Fields{
		"domain":     "s3-socket",
		"action":     "Filter",
		"repository": "FilterDuplicatesFile",
	})
	if err := s.persistence.DuplicatesFile(FileEncrypt, UserID); err != nil {
		log.WithField("type", "DuplicatesFile persistence").Errorln(err)
		return err
	}
	return nil
}

func (s *StorageS3Repository) InsertFileData(modelData *models.DataModel) (int, error) {
	log := logger.Log.WithFields(logrus.Fields{
		"domain":     "s3-socket",
		"action":     "Insert",
		"repository": "InsertFileData",
	})

	ID, err := s.persistence.InsertFileData(modelData)
	if err != nil {
		log.WithField("type", "InsertFileData persistence").Errorln(err)
		return 0, err
	}

	return ID, nil
}

func NewStorageS3Repo(bucket string, svc *s3.S3, session *session.Session, rpc *grpc.ClientConn, s3Persistence interfaces.StorageS3Persistence) interfaces.StorageS3Repository {
	return &StorageS3Repository{bucket: bucket, svc: svc, sess: session, _rpc: rpc, persistence: s3Persistence}
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
