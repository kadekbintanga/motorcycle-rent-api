package helper

import (
	"motorcycle-rent-api/app/global"

	"github.com/sirupsen/logrus"
)

func LogDebug(apiCallID string, message any) {
	global.Logger.WithFields(logrus.Fields{"api_call_id": apiCallID}).Debug(message)
}

func LogInfo(apiCallID string, message any) {
	global.Logger.WithFields(logrus.Fields{"api_call_id": apiCallID}).Info(message)
}

func LogWarning(apiCallID string, message any) {
	global.Logger.WithFields(logrus.Fields{"api_call_id": apiCallID}).Warn(message)
}

func LogError(apiCallID string, message any) {
	global.Logger.WithFields(logrus.Fields{"api_call_id": apiCallID}).Error(message)
}
