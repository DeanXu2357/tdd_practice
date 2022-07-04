package money

type Bank interface {
	Reduce(e Expression, to string) Money
	Rate(from, to string) int64
	SetRate(from, to string, r int64)
}

type bank struct {
	rates map[pair]int64
}

func NewBank() Bank {
	return &bank{rates: map[pair]int64{}}
}

func (b *bank) SetRate(from, to string, r int64) {
	b.rates[pair{from: from, to: to}] = r
}

func (b *bank) Rate(from, to string) int64 {
	if from == to {
		return 1
	}

	return b.rates[pair{from, to}]
}

func (b *bank) Reduce(e Expression, to string) Money {

	return e.Reduce(b, to)
}
