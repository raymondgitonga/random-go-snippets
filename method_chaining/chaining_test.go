package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MethodChaining(t *testing.T) {
	card := Card{}

	t.Run("Test charge", func(t *testing.T) {
		cardResp := NewCardResp(card, nil)
		cardResp.Charge(10.0)

		assert.Equal(t, cardResp.card.amount, 10.0)
	})

	t.Run("Test address", func(t *testing.T) {
		cardResp := NewCardResp(card, nil)
		cardResp.WithAddress("Kenya")

		assert.Equal(t, cardResp.card.address, "Kenya")
	})

	t.Run("Test currency", func(t *testing.T) {
		cardResp := NewCardResp(card, nil)
		cardResp.WithCurrency("KES")

		assert.Equal(t, cardResp.card.currency, "KES")
	})

	t.Run("Test whole chain", func(t *testing.T) {
		cardResp := NewCardResp(card, nil)
		resp, err := cardResp.WithCurrency("KES").WithAddress("Kenya").Charge(10.0).Execute()

		assert.Equal(t, resp, "Dear Customer, \nKES 10 is Debited from your account ")
		assert.NoError(t, err)
	})

	t.Run("Test error handling", func(t *testing.T) {
		cardResp := &CardResp{}
		x, err := cardResp.WithCurrency("KESy").WithAddress("Kenya").Charge(10.0).Execute()

		assert.Equal(t, x, "")
		assert.Error(t, err)
	})

}
