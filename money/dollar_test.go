package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Money_Equality(t *testing.T) {
	five := NewDollar(5)
	assert.True(t, five.Equals(NewDollar(5)))
	assert.False(t, five.Equals(NewDollar(6)))
	assert.False(t, five.Equals(nil))

	five = NewNewTaiwanDollar(5)
	assert.True(t, five.Equals(NewNewTaiwanDollar(5)))
	assert.False(t, five.Equals(NewNewTaiwanDollar(6)))
}

func Test_Money_Equality_Different_Currency(t *testing.T) {
	fiveUSD := NewDollar(5)
	fiveTWD := NewNewTaiwanDollar(5)
	assert.False(t, fiveUSD.Equals(fiveTWD))
}

func Test_Money_Multiplication(t *testing.T) {
	five := NewDollar(5)

	assert.Equal(t, NewDollar(10), five.Times(2))
	assert.Equal(t, NewDollar(15), five.Times(3))

	five = NewNewTaiwanDollar(5)

	assert.Equal(t, NewNewTaiwanDollar(10), five.Times(2))
	assert.Equal(t, NewNewTaiwanDollar(15), five.Times(3))
}

func Test_Money_Currency(t *testing.T) {
	assert.Equal(t, CurTWD, NewNewTaiwanDollar(1).Currency())
	assert.Equal(t, CurUSD, NewDollar(1).Currency())
}

func Test_Money_Addition(t *testing.T) {
	b := NewBank()

	fiveUSD := NewDollar(5)
	sum := fiveUSD.Plus(NewDollar(5))
	reduced := b.Reduce(sum, CurUSD)
	assert.Equal(t, NewDollar(10), reduced)

	fiveTWD := NewNewTaiwanDollar(5)
	sumOfTWD := fiveTWD.Plus(NewNewTaiwanDollar(5))
	assert.Equal(t, NewNewTaiwanDollar(10), b.Reduce(sumOfTWD, CurTWD))
}

func Test_bank_Reduce(t *testing.T) {
	b := NewBank()

	oneUSD := NewDollar(1)
	assert.Equal(t, NewDollar(1), b.Reduce(oneUSD, CurUSD))
}

func Test_bank_Reduce_Different_Currency(t *testing.T) {
	b := NewBank()
	b.SetRate(CurTWD, CurUSD, 29)

	twd29 := NewNewTaiwanDollar(29)

	assert.Equal(t, NewDollar(1), b.Reduce(twd29, CurUSD))
}

func Test_mix_addition(t *testing.T) {
	b := NewBank()
	b.SetRate(CurTWD, CurUSD, 29)

	twd29 := NewNewTaiwanDollar(29)

	sum := NewDollar(1).Plus(twd29)
	assert.Equal(t, NewDollar(2), b.Reduce(sum, CurUSD))
}
