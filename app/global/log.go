package global

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	Logger  *logrus.Logger
	LogFile *os.File
)
