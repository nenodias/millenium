package repositories

import (
	"github.com/rs/zerolog"
	xlog "xorm.io/xorm/log"
)

type CurrentLogger struct {
	logger  *zerolog.Logger
	showSQL bool
}

func (c *CurrentLogger) Debug(v ...interface{}) {
	c.logger.Debug().Msgf("%v", v...)
}
func (c *CurrentLogger) Debugf(format string, v ...interface{}) {
	c.logger.Debug().Msgf(format, v...)
}
func (c *CurrentLogger) Error(v ...interface{}) {
	c.logger.Error().Msgf("%v", v...)
}
func (c *CurrentLogger) Errorf(format string, v ...interface{}) {
	c.logger.Error().Msgf(format, v...)
}
func (c *CurrentLogger) Info(v ...interface{}) {
	c.logger.Info().Msgf("%v", v...)
}
func (c *CurrentLogger) Infof(format string, v ...interface{}) {
	c.logger.Info().Msgf(format, v...)
}
func (c *CurrentLogger) Warn(v ...interface{}) {
	c.logger.Warn().Msgf("%v", v...)
}
func (c *CurrentLogger) Warnf(format string, v ...interface{}) {
	c.logger.Warn().Msgf(format, v...)
}

func (c *CurrentLogger) Level() xlog.LogLevel {
	level := c.logger.GetLevel()
	switch level {
	case zerolog.ErrorLevel:
		return xlog.LOG_ERR
	case zerolog.DebugLevel:
		return xlog.LOG_DEBUG
	case zerolog.FatalLevel:
		return xlog.LOG_ERR
	case zerolog.InfoLevel:
		return xlog.LOG_INFO
	case zerolog.NoLevel:
		return xlog.LOG_OFF
	case zerolog.WarnLevel:
		return xlog.LOG_WARNING
	default:
		return xlog.LOG_UNKNOWN
	}
}
func (c *CurrentLogger) SetLevel(l xlog.LogLevel) {
	c.logger.Debug().Msgf("SetLevel called %v", l)
}

func (c *CurrentLogger) ShowSQL(show ...bool) {
	c.showSQL = show[0]
}
func (c *CurrentLogger) IsShowSQL() bool {
	return c.showSQL
}
