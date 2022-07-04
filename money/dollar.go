package money

const CurUSD = "USD"
const CurTWD = "TWD"

func NewDollar(amount int64) Money {
	return New(amount, CurUSD)
}

func NewNewTaiwanDollar(amount int64) Money {
	return New(amount, CurTWD)
}

func New(amount int64, currency string) Money {
	return &money{amount: amount, currency: currency}
}

type Expression interface {
	Reduce(b Bank, to string) Money
	Times(t int64) Expression
	Plus(addend Expression) Expression
}

type Money interface {
	Expression

	Equals(c Money) bool

	Amount() int64
	Currency() string
}

type money struct {
	amount   int64
	currency string
}

func (m *money) Amount() int64 {
	return m.amount
}

func (m *money) Currency() string {
	return m.currency
}

func (m *money) Plus(addend Expression) Expression {
	a := addend.(Money)
	return &sum{augend: m, addend: a}
}

func (m *money) Times(t int64) Expression {
	return &money{amount: t * m.amount, currency: m.currency}
}

func (m *money) Reduce(b Bank, to string) Money {
	r := b.Rate(m.currency, to)
	return New(m.amount*r, to)
}

func (m *money) Equals(c Money) bool {
	if c == nil {
		return false
	}

	return (m.currency == c.Currency()) &&
		(m.amount == c.Amount())
}

type Bank interface {
	Reduce(e Expression, to string) Money
	Rate(from, to string) int64
}

type bank struct {
}

func (b *bank) Rate(from, to string) int64 {
	//TODO implement me
	panic("implement me")
}

func (b *bank) Reduce(e Expression, to string) Money {
	return e.Reduce(b, to)
}

type sum struct {
	augend Money
	addend Money
}

func (s *sum) Reduce(b Bank, to string) Money {
	return s.augend.Reduce(b, to).
		Plus(s.addend.Reduce(b, to)).
		Reduce(b, to)
}

func (s *sum) Times(t int64) Expression {
	//TODO implement me
	panic("implement me")
}

func (s *sum) Plus(addend Expression) Expression {
	//TODO implement me
	panic("implement me")
}
