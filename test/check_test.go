package test

import (
	"github.com/MaJloe3Jlo/mapisacard_test/lib"
	"testing"
)

func TestCheckLuna(t *testing.T) {
	pan := "5469345678901234"
	val := lib.CheckLuna(pan)
	if val != true {
		t.Error("PAN number is wrong")
	}
}

func TestCheckDate(t *testing.T) {
	month := 6
	year := 2018
	val := lib.CheckDate(month, year)
	if val != true {
		t.Error("Date is wrong")
	}
}

func TestCheckHolder(t *testing.T) {
	holder := "DMITRIY KLESTOV"
	val := lib.CheckHolder(holder)
	if val != true {
		t.Error("Cardholder is wrong")
	}
}

func TestCheckOrderId(t *testing.T) {
	orderId := "124TestPayment"
	val := lib.CheckOrderID(orderId)
	if val != true {
		t.Error("OrderID is wrong")
	}
}

func TestCheckAmount(t *testing.T) {
	amount := 99
	val := lib.CheckAmount(amount)
	if val != true {
		t.Error("Amount is wrong")
	}
}
