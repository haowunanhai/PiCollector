package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type logger struct {
	file     *os.File
	fileName string

	level     logLevel
	calldepth int

	bufferChan chan []byte
	wg         sync.WaitGroup
}

func (l *logger) setLevel(name string) {
	switch strings.ToUpper(name) {
	case "DEBUG":
		l.level = debugLevel
	case "TRACE":
		l.level = traceLevel
	case "INFO":
		l.level = infoLevel
	case "WARNING":
		l.level = warningLevel
	case "ERROR":
		l.level = errorLevel
	case "FATAL":
		l.level = fatalLevel
	default:
		l.level = infoLevel
	}
}

func (l *logger) setFile(fileName string) error {
	// fileName为空,不做处理
	if fileName == "" {
		return nil
	}
	// 打开新文件
	newfile, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	// 设置新日志文件名
	l.fileName = fileName

	// 交换文件
	old := atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&l.file)), unsafe.Pointer(newfile))

	// 获取旧文件,必要的话关闭
	oldfile := (*os.File)(old)
	if oldfile != os.Stdout {
		oldfile.Close()
	}
	return nil
}

func (l *logger) makeBuffer(level logLevel, v ...interface{}) {
	if level < l.level || len(v) <= 0 {
		return
	}
	_, file, line, ok := runtime.Caller(l.calldepth)
	if !ok {
		file = "???"
		line = 0
	}

	var buf []byte
	timeFormat := time.Now().Format("2006-01-02 15:04:05.000")
	buf = append(buf, fmt.Sprintf("%s`%s`%d`%s:%d`%v", level, timeFormat, pid, path.Base(file), line, v[0])...)
	for i := 1; i < len(v)-1; i += 2 {
		buf = append(buf, fmt.Sprintf("`%v=%v", v[i], v[i+1])...)
	}
	if len(v)%2 == 0 {
		buf = append(buf, fmt.Sprintf("`%v", v[len(v)-1])...)
	}

	// 为console上色
	if l.file == os.Stdout {
		buf = []byte(colors[level](string(buf)))
	}
	buf = append(buf, '\n')

	l.bufferChan <- buf
}

// 执行异步输出
func (l *logger) asyncWrite() {
	for buf := range l.bufferChan {
		l.file.Write(buf)
	}
	l.wg.Done()
}
