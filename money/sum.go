package money

type Sum interface {
	Expression
}

type sum struct {
	augend Money
	addend Money
}

func (s *sum) Reduce(bank Bank, to string) Money {
	amount := s.augend.Amount() + s.addend.Amount()

	return New(amount, to)
}
