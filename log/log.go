package log

import (
	"fmt"
	"github.com/mgutz/ansi"
	syslog "log"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	MaxStackDepth = 25
)

var Levels = map[string]int{
	"DEBUG":   0,
	"INFO":    1,
	"NOTICE":  2,
	"WARNING": 3,
	"ERROR":   4,
	"FATAL":   5,
	"PANIC":   5,
	"LOGSYS":  99,
}

type Options struct {
	Name           string
	Level          int
	OutputFilepath string
}

var Opts = Options{
	"log",
	0,
	"",
}

func Setup(opts *Options) {
	syslog.SetFlags(syslog.LstdFlags)

	if opts != nil {
		if opts.Name != "" {
			Opts.Name = opts.Name
		}
		if opts.Level != 0 {
			Opts.Level = opts.Level
		}
		if opts.OutputFilepath != "" {
			Opts.OutputFilepath = opts.OutputFilepath
		}
	}
	write("logsys", "Setup", "Log ready.")
}

func IfError(errs ...interface{}) {
	var err error
	var tag string
	for _, v := range errs {
		switch v.(type) {
		case string:
			tag = v.(string)
			break
		case error:
			err = v.(error)
			break
		default:
			break
		}
	}
	if tag != "" && err != nil {
		write("error", tag, err)
	} else if err != nil {
		write("error", err)
	}
}

// Fatal error printing
func Fatal(v ...interface{}) {
	write("fatal", v...)
}

func Error(err error) {
	write("error", err)
}

func Warning(v ...interface{}) {
	write("warning", v...)
}

func Notice(v ...interface{}) {
	write("notice", v...)
}

func Info(v ...interface{}) {
	write("info", v...)
}

func Debug(v ...interface{}) {
	write("debug", v...)
}

func DebugWithStack(v ...interface{}) {
	write("stacktrace")
	write("debug", v...)
}

func AndPanic(err error) {
	if err != nil {
		write("panic", "", err)
		panic(err)
	}
}

func AndPanicWithMessage(err error, msg string) {
	if err != nil {
		write("panic", "", msg, err)
		panic(err)
	}
}

func write(name string, v ...interface{}) {
	now := time.Now()
	year, month, day := now.Date()
	hour, min, sec := now.Clock()

	// Inspired by/Borrowed from http://golang.org/src/log/log.go ~Line 80
	ts := new([]byte)

	itoa(ts, year, 4)
	*ts = append(*ts, '/')
	itoa(ts, int(month), 2)
	*ts = append(*ts, '/')
	itoa(ts, day, 2)
	*ts = append(*ts, ' ')
	itoa(ts, hour, 2)
	*ts = append(*ts, ':')
	itoa(ts, min, 2)
	*ts = append(*ts, ':')
	itoa(ts, sec, 2)
	*ts = append(*ts, '.')
	itoa(ts, now.Nanosecond()/1e3, 6)
	*ts = append(*ts, ' ')

	level := Levels[name]
	if level >= Opts.Level {

		// Grab the reporting file and line number
		_, file, line, _ := runtime.Caller(2)

		color := "white"
		switch name {
		case "error":
			color = "red+hu"
			break
		case "fatal":
			color = "red+b"
			break
		case "panic":
			color = "red+b:white+h"
			break
		case "warning":
			color = "yellow+b"
			break
		case "notice":
			color = "magenta+b"
			break
		case "info":
			color = "cyan"
			break
		case "debug":
			color = "cyan+hbi"
			break
		case "stacktrace":
			color = "white+b:magenta+h"
			break
		default:
			color = "gray+u:white+h"
			break
		}

		stacktrace := false
		if name == "stacktrace" {
			stacktrace = true
		}

		name = ansi.Color(" "+strings.ToUpper(name)+" ", color)

		// fileDisplay := filepath.Base(file)
		fileDisplay := filepath.Clean(file)

		// if Opts.OutputFilepath {
		// 	fileDisplay = fileDisplay
		// }
		if stacktrace {

			stack := make([]uintptr, MaxStackDepth)
			length := runtime.Callers(2, stack[:])
			stack = stack[:length]
			fmt.Printf("%s %s [%s] ( %s:%d )", ts, Opts.Name, name, fileDisplay, line)
			spaces := " "
			for i, _ := range stack {
				_, file, line, _ = runtime.Caller(2 + (length - i))
				fmt.Printf("%s%s:%d\n\n", spaces, file, line)
				spaces = spaces + " "
			}
			return
		}

		fmt.Printf("%s %s [%s] ( %s:%d )  %s", ts, Opts.Name, name, fileDisplay, line, fmt.Sprintln(v...))
	}
}

// Borrowed from http://golang.org/src/log/log.go ~Line 62
func itoa(buf *[]byte, i int, wid int) {
	var u uint = uint(i)
	if u == 0 && wid <= 1 {
		*buf = append(*buf, '0')
		return
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}
	*buf = append(*buf, b[bp:]...)
}
