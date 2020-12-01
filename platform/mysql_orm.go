package platform

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type connectionStringOrm struct {
	master string
	domain string
}

var stringType = "type"
var stringConnection = "connection"

type ConnectionsOrm interface {
	Open() *DatabaseORM
}

// Database where the opened database connection is being used
type DatabaseORM struct {
	Master *gorm.DB
}

func InitializeORM(master string, domain string) ConnectionsOrm {
	return &connectionStringOrm{master: master, domain: domain}
}

func (cs *connectionStringOrm) Open() *DatabaseORM {
	// log fields for logrus, no need to write this multiple times
	logFields := logrus.Fields{
		"platform": "mysql",
		"domain":   cs.domain,
	}
	logMasterFields := logrus.Fields{
		stringType:       "master",
		stringConnection: cs.master,
	}

	logrus.WithFields(logFields).Info("Connecting to Mysql DB")
	logrus.WithFields(logFields).Info("Opening Connection to Master")

	dbMaster, err := gorm.Open("mysql", cs.master)
	if err != nil {
		logrus.WithFields(logMasterFields).Fatal(err)
		panic(err)
	}

	return &DatabaseORM{
		Master: dbMaster,
	}

}
