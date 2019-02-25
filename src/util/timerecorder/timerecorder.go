package timerecorder

import (
	"fmt"
	"strings"
	"time"
)

type recorder struct {
	mark     string
	markTime int64
}

type TimeRecorder struct {
	flag      string
	startTime int64
	endTime   int64
	recorders []recorder
}

func New(flag string) *TimeRecorder {
	return &TimeRecorder{
		flag:      flag,
		startTime: time.Now().UnixNano(),
	}
}

func NewWithTime(flag string, startTime int64) *TimeRecorder {
	return &TimeRecorder{
		flag:      flag,
		startTime: startTime,
	}
}

func (r *TimeRecorder) Mark(mark string) {
	r.recorders = append(r.recorders, recorder{mark: mark, markTime: time.Now().UnixNano()})
}

func (r *TimeRecorder) End() {
	r.endTime = time.Now().UnixNano()
}

func (r *TimeRecorder) ElapsedTime() int64 {
	elapsedTime := r.endTime - r.startTime
	if elapsedTime < 0 {
		return 0
	}
	return elapsedTime / 1e6
}

func (r *TimeRecorder) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("[%s]", r.flag))
	for _, rec := range r.recorders {
		b.WriteString(fmt.Sprintf("[+%dms,%s]", (rec.markTime-r.startTime)/1e6, rec.mark))
	}
	b.WriteString(fmt.Sprintf("[+%dms,end]", (r.endTime-r.startTime)/1e6))
	return b.String()
}
