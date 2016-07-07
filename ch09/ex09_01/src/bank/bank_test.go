package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	done := make(chan struct{})
	try := 10
	success := 0
	start := Balance()

	go func() {
		for i := 0; i < try; i++ {
			Deposit(100)
		}
		done <- struct{}{}
	}()

	go func() {
		for i := 0; i < try; i++ {
			if Withdraw(200) {
				success++
			}
		}
		done <- struct{}{}
	}()

	<-done
	<-done

	if act, want := 100*try+start-(200*success), Balance(); act != want {
		t.Errorf("Balance = %d, want %d", act, want)
	}
}
