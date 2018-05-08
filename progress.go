package go_progress

import (
	"fmt"
	"strings"
)

type Progress interface {
	IsExceed() bool
	Increment(inc int)
	String() string
}

type progress struct {
	Max int
	Progress int
}

func NewProgress(max int) Progress {
	return &progress{Max: max, Progress: 0}
}

func (p *progress) IsExceed() bool {
	if p.Progress >= p.Max {
		return true
	}
	return false
}

func (p *progress) Increment(inc int) {
	p.Progress += inc
	if p.IsExceed() {
		p.Progress = p.Max
	}
}

func (p *progress) String() string {
	per := float64(p.Progress) / float64(p.Max) * 100
	if p.IsExceed() {
		return ""
	} else {
		return fmt.Sprintf("%5.1f%% [%s>%s] %d \r", per, strings.Repeat("=", int(per)), strings.Repeat(" ", 100-int(per)), p.Progress)
	}
}
