package go_progress

type Bar struct {
	Max int
	Progress int
}

const (
	defaultMAX = 100
)

func newBar() *Bar {
	return &Bar{Max:defaultMAX, Progress: 0}
}

func (b *Bar) isExceed() bool {
	if b.Progress >= b.Max {
		return true
	}
	return false
}

func (b *Bar) increment(inc int) {
	b.Progress += inc
	if b.isExceed() {
		b.Progress = b.Max
	}
}

func (b *Bar) percent() float64 {
	return float64(b.Progress) / float64(b.Max) * 100
}