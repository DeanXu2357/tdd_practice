package money

type Sum interface {
	Expression
}

type sum struct {
	augend Expression
	addend Expression
}

func (s *sum) Times(t int64) Expression {
	//TODO implement me
	panic("implement me")
}

func (s *sum) Plus(addend Expression) Expression {
	//TODO implement me
	panic("implement me")
}

func NewSum(augend, addend Expression) Sum {
	return &sum{augend, addend}
}

func (s *sum) Reduce(bank Bank, to string) Money {
	amount := s.augend.Reduce(bank, to).Amount() + s.addend.Reduce(bank, to).Amount()

	return New(amount, to)
}
