package example6

type Wallet struct {
	Balance float64
}
type Deposit interface {
	Deposit()
	Balance()
}

func (w Wallet) Deposit(amount float64) {

}