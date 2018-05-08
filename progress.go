package go_progress

import (
	"fmt"
	"strings"
	"errors"
)

type Progress interface {
	String() string
}

type progress struct {
	Width int
	Bar *Bar
}

const (
	defaultWidth = 80
)

func New() Progress {
	return &progress{Width: defaultWidth, Bar : newBar()}
}

func (p *progress) WithWidth(w int) (Progress, error) {
	if w < 10 {
		return nil, errors.New("width must be more than 10")
	}
	p.Width = w
	return p, nil
}

func (p *progress) String() string {
	per := p.Bar.percent()
	if p.Bar.isExceed() {
		return ""
	} else {
		return fmt.Sprintf("%5.1f%% [%s>%s] %d \r", per, strings.Repeat("=", int(per)), strings.Repeat(" ", 100-int(per)), p.Bar.Progress)
	}
}
