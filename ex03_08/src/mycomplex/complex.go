package mycomplex

import (
	"math"
	"math/big"
)

/////////////////////////////////////////////////////////////////////////////

type BigFloatComplex struct {
	Re	*big.Float
	Im	*big.Float
}

func NewBigFloatComplex(r, i float64) *BigFloatComplex {
	z := new(BigFloatComplex)
	z.Re = big.NewFloat(r)
	z.Im = big.NewFloat(i)

	return z
}

func (c1 *BigFloatComplex) Add(c2 *BigFloatComplex) *BigFloatComplex {
	z := NewBigFloatComplex(0, 0)
	z.Re.Add(c1.Re, c2.Re)
	z.Im.Add(c1.Im, c2.Im)
	return z
}

func (c1 *BigFloatComplex) Mul(c2 *BigFloatComplex) *BigFloatComplex {
	z := NewBigFloatComplex(0, 0)
	z.Re.Sub(newBigFloat().Mul(c1.Re, c2.Re), newBigFloat().Mul(c1.Im, c2.Im))
	z.Im.Add(newBigFloat().Mul(c1.Re, c2.Im), newBigFloat().Mul(c1.Im, c2.Re))
	return z
}

func (c *BigFloatComplex) Abs() float64 {
	r, _ := c.Re.Float64()
	i, _ := c.Im.Float64()
	return math.Sqrt(r * r + i * i)
}

/////////////////////////////////////////////////////////////////////////////

type BigRatComplex struct {
	Re	*big.Rat
	Im	*big.Rat
}

func NewBigRatComplex(r, i *big.Rat) *BigRatComplex {
	z := new(BigRatComplex)
	z.Re = r
	z.Im = i
	return z
}

func (c1 *BigRatComplex) Add(c2 *BigRatComplex) *BigRatComplex {
	z := NewBigRatComplex(NewBigRat(), NewBigRat())
	z.Re.Add(c1.Re, c2.Re)
	z.Im.Add(c1.Im, c2.Im)
	return z
}

func (c1 *BigRatComplex) Mul(c2 *BigRatComplex) *BigRatComplex {
	z := NewBigRatComplex(NewBigRat(), NewBigRat())
	z.Re.Sub(NewBigRat().Mul(c1.Re, c2.Re), NewBigRat().Mul(c1.Im, c2.Im))
	z.Im.Add(NewBigRat().Mul(c1.Re, c2.Im), NewBigRat().Mul(c1.Im, c2.Re))
	return z
}

func (c *BigRatComplex) Abs() float64 {
	r, _ := c.Re.Float64()
	i, _ := c.Im.Float64()
	return math.Sqrt(r * r + i * i)
}

/////////////////////////////////////////////////////////////////////////////

func newBigFloat() *big.Float {
	return big.NewFloat(0)
}

func NewBigRat() *big.Rat {
	return big.NewRat(0, 1)
}
