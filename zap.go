package rslog

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 格式化时间
func ZapTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func ZapNewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",                           // json时时间键
		LevelKey:       "L",                           // json时日志等级键
		NameKey:        "N",                           // json时日志记录器键
		CallerKey:      "C",                           // json时日志文件信息键
		MessageKey:     "M",                           // json时日志消息键
		StacktraceKey:  "S",                           // json时堆栈键
		LineEnding:     zapcore.DefaultLineEnding,     // 友好日志换行符
		EncodeLevel:    zapcore.CapitalLevelEncoder,   // 友好日志等级名大小写（info INFO）
		EncodeTime:     ZapTimeEncoder,                // 友好日志时日期格式化
		EncodeDuration: zapcore.StringDurationEncoder, // 时间序列化
		EncodeCaller:   zapcore.ShortCallerEncoder,    // 日志文件信息（包/文件.go:行号）
	}
}

type zapLog struct {
	Logger *zap.Logger

	level    string
	sugarLog *zap.SugaredLogger
}

func GetZapLog() *zap.Logger {
	if zlog, ok := v_logger.(*zapLog); ok {
		return zlog.Logger
	}
	return nil
}

func newZapLog() *zapLog {
	p := &zapLog{}

	p.Logger = initZapLog(zapcore.InfoLevel)
	p.sugarLog = p.Logger.Sugar()

	return p
}

func initZapLog(l zapcore.Level) *zap.Logger {
	core := zapcore.NewCore(
		// zapcore.NewJSONEncoder(ZapNewEncoderConfig()),
		zapcore.NewConsoleEncoder(ZapNewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(os.Stdout),
		l,
	)
	return zap.New(core, zap.AddCaller())
}

func (p *zapLog) Info(datas ...interface{}) {
	p.sugarLog.Info(datas...)
}

func (p *zapLog) Infof(format string, datas ...interface{}) {
	p.sugarLog.Infof(format, datas...)
}

func (p *zapLog) Debug(datas ...interface{}) {
	p.sugarLog.Debug(datas...)
}

func (p *zapLog) Debugf(format string, datas ...interface{}) {
	p.sugarLog.Debugf(format, datas...)
}

func (p *zapLog) Warn(datas ...interface{}) {
	p.sugarLog.Warn(datas...)
}

func (p *zapLog) Warnf(format string, datas ...interface{}) {
	p.sugarLog.Warnf(format, datas...)
}

func (p *zapLog) Error(datas ...interface{}) {
	p.sugarLog.Error(datas...)
}

func (p *zapLog) Errorf(format string, datas ...interface{}) {
	p.sugarLog.Errorf(format, datas...)
}

func (p *zapLog) SetLevel(l string) {
	p.Logger = initZapLog(GetZapLevel(l))
	p.sugarLog = p.Logger.Sugar()
}

// "debug": zapcore.DebugLevel,
// "info": zapcore.InfoLevel,
// "warn": zapcore.WarnLevel,
// "error": zapcore.ErrorLevel,
// "dpanic": zapcore.DPanicLevel,
// "panic": zapcore.PanicLevel,
// "fatal": zapcore.FatalLevel,
func GetZapLevel(l string) zapcore.Level {
	switch l {
	case "info":
		return zapcore.InfoLevel
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	case "dpanic":
		return zapcore.DPanicLevel
	}
	return zapcore.InfoLevel
}
