package internal

import (
	"gorm.io/gorm/logger"
	"server/config"
)

type Writer struct {
	config config.GeneralDB
	writer logger.Writer
}
