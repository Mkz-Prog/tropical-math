package math

type Rational struct  {
	Value float64
	IsInfinity bool
}

var Inf = Rational{Value: 0, IsInfinity: true}

func (r1 Rational) Add(r2 Rational) Rational {
	if r1.IsInfinity {
		return r2
	} else if r2.IsInfinity {
		return r1
	} else {
		return Rational{Value: min(r1.Value, r2.Value), IsInfinity: false}
	}
}

func (r1 Rational) Multiply(r2 Rational) Rational {
	if r1.IsInfinity {
		return Inf
	} else if r2.IsInfinity {
		return Inf
	} else {
		return Rational{Value: r1.Value + r2.Value, IsInfinity: false}
	}
}


