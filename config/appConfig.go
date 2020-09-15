package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"warehousing/appLogging"
)

type AppConfig struct {
	logFilePath string
	DataPath    string
	Aws         *session.Session
	BucketName  string
}

func NewConfig() *AppConfig {
	c := &AppConfig{}
	logFile := c.ViperGetEnv("LOG_FILE", "/tmp/app.log")
	logLevel := c.ViperGetEnv("LOG_LEVEL", "debug")
	dataPath := c.ViperGetEnv("FILE_PATH", "/tmp/warehousing.csv")
	bucketName := c.ViperGetEnv("BUCKET_NAME", "")
	c.logFilePath = logFile
	c.logger(logLevel)
	c.DataPath = dataPath
	c.connectAws()
	c.BucketName = bucketName
	return c
}

func (c *AppConfig) logger(level string) {
	logFormat := new(logrus.JSONFormatter)
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	log := logrus.Logger{
		Out:       os.Stdout,
		Formatter: logFormat,
		Level:     logLevel,
	}
	var file, errFile = os.OpenFile(c.logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errFile != nil {
		panic(err)
	}
	log.SetOutput(file)
	appLogging.Logger = appLogging.NewAppLogger(log)
}

func (c *AppConfig) ViperGetEnv(key, defaultValue string) string {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}
func (c *AppConfig) connectAws() {
	accessKeyID := c.ViperGetEnv("AWS_ACCESS_KEY_ID", "")
	secretAccessKey := c.ViperGetEnv("AWS_SECRET_ACCESS_KEY", "")
	myRegion := c.ViperGetEnv("AWS_REGION", "")
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(myRegion),
			Credentials: credentials.NewStaticCredentials(
				accessKeyID,
				secretAccessKey,
				"",
			),
		})
	if err != nil {
		panic(err)
	}
	c.Aws = sess
}
