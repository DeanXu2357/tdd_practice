package money

type Sum interface {
	Expression
}

type sum struct {
	augend Money
	addend Money
}

func NewSum(augend, addend Money) Sum {
	return &sum{augend, addend}
}

func (s *sum) Reduce(bank Bank, to string) Money {
	amount := s.augend.Reduce(bank, to).Amount() + s.addend.Reduce(bank, to).Amount()

	return New(amount, to)
}
