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
	Reduce(bank Bank, to string) Money
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

func (m *money) Equals(c Money) bool {
	if c == nil {
		return false
	}

	return (m.currency == c.Currency()) &&
		(m.amount == c.Amount())
}

func (m *money) Plus(addend Expression) Expression {
	return NewSum(m, addend)
}

func (m *money) Times(t int64) Expression {
	return &money{amount: t * m.amount, currency: m.currency}
}

func (m *money) Reduce(bank Bank, to string) Money {
	rate := bank.Rate(m.currency, to)

	return New(m.amount/rate, to)
}
