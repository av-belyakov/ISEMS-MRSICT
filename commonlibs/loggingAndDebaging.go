package commonlibs

import (
	"fmt"
	"runtime"
)

//GetFuncName возвращает название функции используя стек вызовов Go
func GetFuncName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	msg := fmt.Sprintf("function: %s", frame.Function)
	return msg
}

//GetInfoAboutFunc возвращает информащию о месте вызова и названии функции используя стек вызовов Go
func GetInfoAboutFunc() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	msg := fmt.Sprintf("function: %s, file: %s, line: %d", frame.Function, frame.File, frame.Line)
	return msg
}
