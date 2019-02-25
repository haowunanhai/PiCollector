package timerecorder

import (
	"testing"
	"time"
)

func TestTimeRecorder(t *testing.T) {
	r := New("TestTimeRecorder")
	time.Sleep(1e7)
	r.Mark("sleep1")
	time.Sleep(3e7)
	r.Mark("sleep2")
	r.End()
	t.Log(r)
}
