package log

import (
	"os"
)

var (
	pid = os.Getpid()

	// 设置缺省的logger
	defaultLogger = &logger{
		level:      infoLevel,
		file:       os.Stdout,
		calldepth:  2,
		bufferChan: make(chan []byte, 1024),
	}
)

// Debug 输出
func Debug(v ...interface{}) {
	defaultLogger.makeBuffer(debugLevel, v...)
}

// Trace 输出
func Trace(v ...interface{}) {
	defaultLogger.makeBuffer(traceLevel, v...)
}

// Info 输出
func Info(v ...interface{}) {
	defaultLogger.makeBuffer(infoLevel, v...)
}

// Warning 输出
func Warning(v ...interface{}) {
	defaultLogger.makeBuffer(warningLevel, v...)
}

// Error 输出
func Error(v ...interface{}) {
	defaultLogger.makeBuffer(errorLevel, v...)
}

// Fatal 输出
func Fatal(v ...interface{}) {
	defaultLogger.makeBuffer(fatalLevel, v...)
}

// SetFile设置日志文件
func SetFile(fileName string) {
	if fileName != "" {
		// 设置日志文件出错，转到标准输出
		if err := defaultLogger.setFile(fileName); err != nil {
			Error("set log file error", "filename", fileName, "error", err)
		}
	}
}

// SetLevel 设置日志级别
func SetLevel(level string) {
	if level != "" {
		defaultLogger.setLevel(level)
	}
}

// Rotate 执行日志滚动,相当于重新设置文件
func Rotate() {
	// 只有指定了日志文件时才能Rotate
	if defaultLogger.fileName != "" {
		// 设置日志文件出错,保持不变
		if err := defaultLogger.setFile(defaultLogger.fileName); err != nil {
			Error("rotate log file error", "filename", defaultLogger.fileName, "error", err)
			return
		}
		Info("rotate log file", "filename", defaultLogger.fileName)
	}
}

// Shutdown 关闭日志，因为日志为异步写入,不主动关闭可能会导致日志丢失
func Shutdown() {
	close(defaultLogger.bufferChan)
	defaultLogger.wg.Wait()
}

// 启动异步写入
func init() {
	defaultLogger.wg.Add(1)
	go defaultLogger.asyncWrite()
}
