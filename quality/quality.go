package quality

type quality int

func (q quality) Value() int {
	return int(q)
}

func (q quality) Inc(i int) quality {
	if newVal := q.Value() + i; newVal >= 0 {
		return quality(newVal)
	}
	return Zero
}

func (q quality) Dec() quality {
	if newVal := q.Value() - 1; newVal >= 0 {
		return quality(newVal)
	}
	return Zero
}

type Q interface {
	Value() int
	Inc(i int) quality
	Dec() quality
}

const (
	Zero = quality(iota)
	One
	Two
)

type ErrNegativeQuality struct{}

func (e ErrNegativeQuality) Error() string {
	return "Quality can't be less than zero."
}

func New(value int) (Q, error) {
	if value < 0 {
		return nil, ErrNegativeQuality{}
	}
	return quality(value), nil
}
