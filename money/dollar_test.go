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
	b := &bank{}

	fiveUSD := NewDollar(5)
	s := fiveUSD.Plus(NewDollar(5)).Reduce(b, CurUSD)
	assert.Equal(t, NewDollar(10), s)

	fiveTWD := NewNewTaiwanDollar(5)
	sumOfTWD := fiveTWD.Plus(NewNewTaiwanDollar(5))
	assert.Equal(t, NewNewTaiwanDollar(10), b.Reduce(sumOfTWD, CurTWD))
}

func Test_bank_Reduce(t *testing.T) {
	b := &bank{}

	oneUSD := NewDollar(1)
	assert.Equal(t, NewNewTaiwanDollar(29), b.Reduce(oneUSD, CurTWD))
}
