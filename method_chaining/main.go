package main

import (
	"errors"
	"fmt"
)

type Card struct {
	amount   float64
	currency string
	address  string
}

type CardResp struct {
	card Card
	err  error
}

func NewCardResp(card Card, err error) *CardResp {
	return &CardResp{
		card: card,
		err:  err,
	}
}

func (c *CardResp) Charge(amount float64) *CardResp {
	if c.err != nil {
		return c
	}
	c.card.amount = amount
	return &CardResp{c.card, nil}
}

func (c *CardResp) WithCurrency(currency string) *CardResp {
	if c.err != nil {
		return c
	}
	if currency != "KES" {
		return &CardResp{
			card: Card{},
			err:  errors.New("wrong currency type"),
		}
	}
	c.card.currency = currency
	return &CardResp{c.card, nil}
}

func (c *CardResp) WithAddress(address string) *CardResp {
	if c.err != nil {
		return c
	}
	c.card.address = address
	return &CardResp{c.card, nil}
}

func (c *CardResp) Execute() (string, error) {
	if c.err != nil {
		return "", c.err
	}
	return fmt.Sprintf("Dear Customer, \n%s %v is Debited from your account ", c.card.currency, c.card.amount), nil
}

func main() {
	card := Card{}

	cardResp := NewCardResp(card, nil)

	x, err := cardResp.Charge(10.0).
		WithCurrency("KES").
		WithAddress("NAIROBI").
		Execute()

	if err != nil {
		fmt.Printf("error processing card: %s", err)
	}
	fmt.Println(x)
}
