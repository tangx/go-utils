package logx

import "github.com/sirupsen/logrus"

func InitLogrus() {

	logrus.SetReportCaller(true)
}

func InitLogursJson(level logrus.Level) {

	logrus.SetFormatter(&logrus.JSONFormatter{})
}
