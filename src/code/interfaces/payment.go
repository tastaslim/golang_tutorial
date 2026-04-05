package interfaces

import "fmt"

/*
Create an interface Payment with method Pay(amount int). Implement CreditCard, UPI, NetBanking.
*/
type Payment interface {
	Pay(amount int)
}

type CreditCard struct {
	Name string
}

type UPI struct {
	Name string
}

type NetBanking struct {
	Name string
}

func (c CreditCard) Pay(amount int) {
	fmt.Printf("The Payment is done via %s and amount is %d", c.Name, amount)
}

func (u UPI) Pay(amount int) {
	fmt.Printf("The Payment is done via %s and amount is %d", u.Name, amount)
}

func (n NetBanking) Pay(amount int) {
	fmt.Printf("The Payment is done via %s and amount is %d", n.Name, amount)
}

func MakePayment(p Payment, amount int) {
	p.Pay(amount)
}
