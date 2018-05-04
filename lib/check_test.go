package lib

import (
	"testing"
)

func TestValidate(t *testing.T) {
	//Проверка с корректниыми данными
	blockTest := BlockRequest{}
	blockTest.MerchantContactID = 1
	blockTest.Card.PAN = "5469345678901234"
	blockTest.Card.EMonth = 6
	blockTest.Card.EYear = 2020
	blockTest.Card.Holder = "VASYA PUPKIN"
	blockTest.Card.CVV = 443
	blockTest.OrderID = "IPhone X"
	blockTest.Amount = 999
	val := Validate(blockTest)
	if val.DealID == -1 {
		t.Error("Validate doesn't work")
	}
	//Проверка с некорректными данными
	blockTest.Card.EYear = 2012
	val = Validate(blockTest)
	if val.DealID != -1 {
		t.Error("Validate doesn't work")
	}
}

func TestMerchantID(t *testing.T) {
	//Проверка с корректниыми данными
	mid := 23434
	val := CheckMerchantID(mid)
	if val != true {
		t.Error("Merchant ID is wrong")
	}
	//Проверка с некорректными данными
	mid = -1
	val = CheckMerchantID(mid)
	if val != false {
		t.Error("Merchant ID is wrong")
	}
}

func TestCheckLuhn(t *testing.T) {
	//Проверка с корректниыми данными
	PAN := "5469345678901234"
	val := CheckLuhn(PAN)
	if val != true {
		t.Error("PAN number is wrong")
	}
	//Проверка с некорректными данными
	PAN = "123"
	val = CheckLuhn(PAN)
	if val != false {
		t.Error("PAN number is wrong")
	}
}

func TestCheckDate(t *testing.T) {
	//Проверка с корректниыми данными
	month := 6
	year := 2018
	val := CheckDate(month, year)
	if val != true {
		t.Error("Date is wrong")
	}
	//Проверка с некорректными данными
	month = 6
	year = 2012
	val = CheckDate(month, year)
	if val != false {
		t.Error("Date is wrong")
	}
}

func TestCheckHolder(t *testing.T) {
	//Проверка с корректниыми данными
	holder := "DMITRIY KLESTOV"
	val := CheckHolder(holder)
	if val != true {
		t.Error("Cardholder is wrong")
	}
	//Проверка с некорректными данными
	holder = "KLES435 TOV"
	val = CheckHolder(holder)
	if val != false {
		t.Error("Cardholder is wrong")
	}
}

func TestCheckCVV(t *testing.T) {
	//Проверка с корректниыми данными
	CVV := 345
	val := CheckCVV(CVV)
	if val != true {
		t.Error("CVV code is wrong")
	}
	//Проверка с некорректными данными
	CVV = 23
	val = CheckCVV(CVV)
	if val != false {
		t.Error("CVV code is wrong")
	}
}

func TestCheckOrderId(t *testing.T) {
	//Проверка с корректниыми данными
	orderID := "124TestPayment"
	val := CheckOrderID(orderID)
	if val != true {
		t.Error("OrderID is wrong")
	}
	//Проверка с некорректными данными
	orderID = ""
	val = CheckOrderID(orderID)
	if val != false {
		t.Error("OrderID is wrong")
	}
}

func TestCheckAmount(t *testing.T) {
	//Проверка с корректниыми данными
	amount := 99
	val := CheckAmount(amount)
	if val != true {
		t.Error("Amount is wrong")
	}
	//Проверка с некорректными данными
	amount = 0
	val = CheckAmount(amount)
	if val != false {
		t.Error("Amount is wrong")
	}
}
