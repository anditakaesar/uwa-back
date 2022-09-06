package application

import (
	"go.uber.org/zap"
)

type Context struct {
	Log *zap.Logger
}
