package rslog

import (
    "fmt"
    "log"
    "os"
)

type sysLog struct {
    Module string

    level string

    logger *log.Logger
}

func newSysLog(module string) *sysLog {
    return &sysLog{
        Module: module,
        level:  "info",
        logger: log.New(os.Stdout, "", log.LstdFlags),
    }
}

func (p *sysLog) print(level string, fprintf func(string, ...interface{}), datas ...interface{}) {
    if getLogLevel(level) > getLogLevel(p.level) {
        return
    }
    format := fmt.Sprintf("[%v] [%v] %v", level, p.Module, stackTrace(3)) + " %v"
    fprintf(format, datas...)
}

func (p *sysLog) printf(level string, fprintf func(string, ...interface{}), format string, datas ...interface{}) {
    if getLogLevel(level) > getLogLevel(p.level) {
        return
    }
    format = fmt.Sprintf("[%v] [%v] %v ", level, p.Module, stackTrace(3)) + format
    fprintf(format, datas...)
}

func (p *sysLog) Info(datas ...interface{}) {
    p.print("info", p.logger.Printf, datas...)
}
func (p *sysLog) Infof(format string, datas ...interface{}) {
    p.printf("info", p.logger.Printf, format, datas...)
}

func (p *sysLog) Debug(datas ...interface{}) {
    p.print("debug", p.logger.Printf, datas...)
}
func (p *sysLog) Debugf(format string, datas ...interface{}) {
    p.printf("debug", p.logger.Printf, format, datas...)
}

func (p *sysLog) Warn(datas ...interface{}) {
    p.print("warn", p.logger.Printf, datas...)
}
func (p *sysLog) Warnf(format string, datas ...interface{}) {
    p.printf("warn", p.logger.Printf, format, datas...)
}
func (p *sysLog) Error(datas ...interface{}) {
    p.print("error", p.logger.Printf, datas...)
}
func (p *sysLog) Errorf(format string, datas ...interface{}) {
    p.printf("error", p.logger.Printf, format, datas...)
}

func (p *sysLog) SetLevel(l string) {
    p.level = l
}
func (p *sysLog) ResetLog(l interface{}) {
    if logger, ok := l.(*log.Logger); ok {
        p.logger = logger
    }
}
