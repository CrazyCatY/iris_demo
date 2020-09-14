package log

import (
	"fmt"
	"runtime"

	"github.com/kataras/golog"
)

//init package
func init() {
	golog.SetTimeFormat("2006/01/02 15:04")

	// Level defaults to "info",
	// but you can change it:
	//golog.SetLevel("debug")

	golog.Handle(func(l *golog.Log) bool {
		prefix := golog.GetTextForLevel(l.Level, true)
		pc, fn, line, _ := runtime.Caller(7)
		message := fmt.Sprintf("%s line %d (%s) (%s) %s: %s",
			prefix, line, runtime.FuncForPC(pc).Name(), fn, l.FormatTime(), l.Message)

		if l.NewLine {
			message += "\n"
		}
		
		fmt.Print(message)
		return true
	})
}

//Println wrap
func Println(v ...interface{}) {
	golog.Println(v...)
}

//Fatal wrap
func Printf(format string, v ...interface{}) {
	golog.Println(fmt.Sprintf(format, v...))
}

//Fatal wrap
func Info(v ...interface{}) {
	golog.Info(v...)
}

//Fatal wrap
func Warn(v ...interface{}) {
	golog.Warn(v...)
}

//Fatal wrap
func Error(v ...interface{}) {
	golog.Error(v...)
}

//Fatal wrap
func Debug(v ...interface{}) {
	golog.Debug(v...)
}

//Fatal wrap
func Fatal(v ...interface{}) {
	golog.Fatal(v...)
}
