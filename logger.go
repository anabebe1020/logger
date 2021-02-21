package logger

import (
	"fmt"
	"os"
	"time"
)

// LogConf ...
type LogConf struct {
	LogLevel string
	LogPath  string
	PrgName  string
}

// output ...
func output(logConf LogConf, logLevel, logStr string) {
	// get time.
	t := time.Now()
	const layout = "2006/01/02 15:04:05"
	// create log string.
	log := "[" + string(t.Format(layout)) + "] [" + logLevel + "] " + logStr
	// file open.
	file, err := os.OpenFile(logConf.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	// file close.
	defer file.Close()
	// file write.
	fmt.Fprintln(file, log)
	//fmt.Println(log)
}

// ErrorLog ...
func ErrorLog(logConf LogConf, logStr string) {
	output(logConf, "ERR", logStr)
}

// WarningLog ...
func WarningLog(logConf LogConf, logStr string) {
	output(logConf, "WRN", logStr)
}

// InfoLog ...
func InfoLog(logConf LogConf, logStr string) {
	output(logConf, "INF", logStr)
}

// DebugLog ...
func DebugLog(logConf LogConf, logStr string) {
	if logConf.LogLevel == "DBG" {
		output(logConf, "DBG", logStr)
	}
}
