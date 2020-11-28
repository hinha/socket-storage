package main

import (
	"github.com/atreugo/websocket"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"github.com/savsgio/atreugo/v11"
	"log"
	"os"
	"socket-storage/controller"
	"socket-storage/platform"
	"socket-storage/repository"
	"socket-storage/service"
)

const domain = "kalkula-storage"

var upgrader = websocket.New(websocket.Config{
	AllowedOrigins: []string{"*"},
})

var (
	s3Object  *s3.S3
	awSession *session.Session
	basePath  string
	filePath  string
)

func init() {
	basePath, _ = os.Getwd()
	if len(os.Args) == 0 {
		panic("environment variable required")
	}

	switch os.Args[1] {
	case "development":
		filePath = basePath + "/cmd/s3/.env.production"
	default:
		filePath = basePath + "/cmd/s3/.env"
	}
	if err := godotenv.Load(filePath); err != nil {
		log.Fatal("Error loading .env file")
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("BUCKET_REGION")),
	})
	if err != nil {
		panic(err)
	}

	s3Object = s3.New(sess)
	awSession = sess

}

func main() {
	// Configure grpc connection
	rpcClient := platform.InitializeGrpc(os.Getenv("RPC_HOST"), os.Getenv("RPC_PORT"), domain)
	rpcConn := rpcClient.Open()
	defer rpcConn.Close()

	var (
		storageRepo       = repository.NewStorageS3Repo(os.Getenv("BUCKET_NAME"), s3Object, awSession, rpcConn)
		storageService    = service.NewStorageS3Service(storageRepo)
		storageController = controller.NewStorageS3Controller(storageService)
	)

	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}
	server := atreugo.New(config)

	server.UseBefore(func(ctx *atreugo.RequestCtx) error {
		ctx.SetUserValue("name", "atreugo")

		return ctx.Next()
	})

	server.GET("/ws", upgrader.Upgrade(storageController.Test))

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
