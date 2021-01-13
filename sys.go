package rslog

import (
    "fmt"
    "log"
)

type sysLog struct {
    Module string

    level string
}

func newSysLog(module string) *sysLog {
    return &sysLog{
        Module: module,
        level:  "info",
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
    p.print("info", log.Printf, datas...)
}
func (p *sysLog) Infof(format string, datas ...interface{}) {
    p.printf("info", log.Printf, format, datas...)
}

func (p *sysLog) Debug(datas ...interface{}) {
    p.print("debug", log.Printf, datas...)
}
func (p *sysLog) Debugf(format string, datas ...interface{}) {
    p.printf("debug", log.Printf, format, datas...)
}

func (p *sysLog) Warn(datas ...interface{}) {
    p.print("warn", log.Printf, datas...)
}
func (p *sysLog) Warnf(format string, datas ...interface{}) {
    p.printf("warn", log.Printf, format, datas...)
}
func (p *sysLog) Error(datas ...interface{}) {
    p.print("error", log.Printf, datas...)
}
func (p *sysLog) Errorf(format string, datas ...interface{}) {
    p.printf("error", log.Printf, format, datas...)
}

func (p *sysLog) SetLevel(l string) {
    p.level = l
}
