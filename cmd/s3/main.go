package main

import (
	"fmt"
	"github.com/atreugo/websocket"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/certifi/gocertifi"
	"github.com/fasthttp-contrib/render"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/savsgio/atreugo/v11"
	"log"
	"os"
	"socket-storage/controller"
	logger "socket-storage/log"
	"socket-storage/persistence"
	"socket-storage/platform"
	"socket-storage/repository"
	"socket-storage/service"
	"time"
)

const domain = "kalkula-storage"

var upgrader = websocket.New(websocket.Config{
	AllowedOrigins:  []string{"*"},
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
})

var (
	s3Object  *s3.S3
	awSession *session.Session
	basePath  string
	filePath  string
)

func init() {
	os.Setenv("TZ", "Asia/Jakarta")
	loc, err := time.LoadLocation(os.Getenv("TZ"))
	time.Local = loc

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

	rootCAs, err := gocertifi.CACerts()
	sentryClientOptions := sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: "https://290acf839ffd475a8745a91c0721b0e9@o483966.ingest.sentry.io/5536611",
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		Environment: "",
		Release:     "",
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	}

	if err != nil {
		log.Println("Could not load CA Certificates: %v\n", err)
	} else {
		sentryClientOptions.CaCerts = rootCAs
	}

	err = sentry.Init(sentryClientOptions)
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	logger.NewLogger()

}

func main() {
	// Configure grpc connection
	rpcClient := platform.InitializeGrpc(os.Getenv("RPC_HOST"), os.Getenv("RPC_PORT"), domain)
	rpcConn := rpcClient.Open()
	defer rpcConn.Close()

	mysqlURL := fmt.Sprintf("%s:%s@(%s:3306)/%s?parseTime=true",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"))
	mysqlDriver := platform.InitializeORM(mysqlURL, domain)
	mysqlConn := mysqlDriver.Open()

	template := render.New(&render.Config{
		Directory: "views",
		//Layout: "layouts",
		HTMLContentType: "text/html",
		Extensions:      []string{".html"},
	})

	var (
		persistenceDb     = persistence.NewStorageS3Service(mysqlConn)
		storageRepo       = repository.NewStorageS3Repo(os.Getenv("BUCKET_NAME"), s3Object, awSession, rpcConn, persistenceDb)
		storageService    = service.NewStorageS3Service(storageRepo)
		storageController = controller.NewStorageS3Controller(storageService, template)
	)

	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}
	server := atreugo.New(config)
	server.UseBefore(func(ctx *atreugo.RequestCtx) error {
		ctx.SetUserValue("name", "atreugo")

		return ctx.Next()
	})

	server.GET("/ws", upgrader.Upgrade(storageController.UploadFile))
	//server.GET("/log", upgrader.Upgrade(storageController.LogFile))
	server.GET("/", storageController.UploadFileRender)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
