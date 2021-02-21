# How to use

## 1. Import library

```go
import "logger"
```

## 2. Set config

```go
var logConf logger.LogConf

func hoge () {
  logConf.PrgName  = "Program Name"
  logConf.LogLevel = "DBG" or "INF"
  logConf.LogPath  = "Output Dir Path"
}
```

## 3. Exec function

```go
func huga () {
  logger.DebugLog(logConf, "this log is debug.")
  logger.ErrorLog(logConf, "this log is error.")
}
```
↓
```logger.log
[YYYY/MM/DD hh:mm:ss] [DBG] this log is debug.
[YYYY/MM/DD hh:mm:ss] [ERR] this log is error.
```
