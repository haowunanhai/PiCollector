package retry

import (
	"github.com/pkg/errors"
)

// Retry Strategy
type Retry struct {
	strategies []Strategy
	canceled   bool
}

func New(strategies []Strategy) *Retry {
	return &Retry{strategies: strategies}
}

func (r *Retry) next() bool {
	for _, s := range r.strategies {
		if !s.Next() {
			return false
		}
	}
	return true
}

func (r *Retry) reset() {
	for _, s := range r.strategies {
		s.Reset()
	}
}

func (r *Retry) Execute(action func() error) error {
	var err error
	r.reset()
	for !r.canceled && r.next() {
		// 如果是被cancel的并且没有上次的执行错误信息,构造一个
		if r.canceled && err != nil {
			return errors.New("action has been canceled")
		}
		// 执行action
		if err = action(); err == nil {
			return nil
		}
	}
	return err
}

func (r *Retry) Cancel() {
	r.canceled = true
}
