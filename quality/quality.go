package quality

type quality struct {
	value int
}

func (q quality) Value() int {
	return q.value
}

func (q quality) Inc(i int) quality {
	if newVal := q.value + i; newVal >= 0 {
		return quality{newVal}
	}
	return Zero
}

func (q quality) Dec() quality {
	if newVal := q.value - 1; newVal >= 0 {
		return quality{newVal}
	}
	return Zero
}

type Q interface {
	Value() int
	Inc(i int) quality
	Dec() quality
}

var Zero = quality{0}

type ErrNegativeQuality struct{}

func (e ErrNegativeQuality) Error() string {
	return "Quality can't be less than zero."
}

func New(value int) (Q, error) {
	if value < 0 {
		return nil, ErrNegativeQuality{}
	}
	return quality{value}, nil
}
