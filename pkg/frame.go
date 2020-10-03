package pkg

type Frame struct {
	First, Second int
}

func NewFrame(a, b int) Frame {
	return Frame{
		First:  a,
		Second: b,
	}
}

func (f *Frame) IsStrike() bool {
	return f.First == 10
}

func (f *Frame) IsSpare() bool {
	return f.First+f.Second == 10
}
