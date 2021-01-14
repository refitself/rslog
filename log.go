package rslog

const (
	C_Log_Sys = "sys"
	C_Log_Zap = "zap"
)

func Info(datas ...interface{}) {
	v_logger.Info(datas...)
}

func Infof(format string, datas ...interface{}) {
	v_logger.Infof(format, datas...)
}

func Debug(datas ...interface{}) {
	v_logger.Debug(datas...)
}

func Debugf(format string, datas ...interface{}) {
	v_logger.Debugf(format, datas...)
}

func Warn(datas ...interface{}) {
	v_logger.Warn(datas...)
}

func Warnf(format string, datas ...interface{}) {
	v_logger.Warnf(format, datas...)
}

func Error(datas ...interface{}) {
	v_logger.Error(datas...)
}

func Errorf(format string, datas ...interface{}) {
	v_logger.Errorf(format, datas...)
}

func SetLevel(l string) {
	if getLogLevel(l) != 0 {
		v_level = l
	}
	v_logger.SetLevel(l)
}

func ResetLog(l interface{}) {
	v_logger.ResetLog(l)
}

func init() {
	UseLog(C_Log_Sys)
}

func UseLog(name string) {
	if name == "" {
		name = C_Log_Zap
	}

	switch name {
	case C_Log_Sys:
		v_logger = newSysLog(C_Log_Sys)
	case C_Log_Zap:
		v_logger = newZapLog()
	}
}
