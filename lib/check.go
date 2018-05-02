package lib

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

//Метод проверки всех полей метода Block
func Validate(b Block_req) *Block_resp {
	var resp Block_resp

	if CheckMerchantId(b.MerchantContactId) == false {
		resp.Error = append(resp.Error, "Error terminal ID; ")
	}
	if CheckLuna(b.Card.Pan) == false {
		resp.Error = append(resp.Error, "Error PAN number; ")
	}
	if CheckDate(b.Card.EMonth, b.Card.EYear) == false {
		resp.Error = append(resp.Error, "Error date of card; ")
	}
	if CheckCvv(b.Card.Cvv) == false {
		resp.Error = append(resp.Error, "Error CVV; ")
	}
	if CheckHolder(b.Card.Holder) == false {
		resp.Error = append(resp.Error, "Error holder name; ")
	}
	if CheckOrderID(b.OrderId) == false {
		resp.Error = append(resp.Error, "Error order ID; ")
	}
	if CheckAmount(b.Amount) == false {
		resp.Error = append(resp.Error, "Error amount; ")
	}

	if len(resp.Error) == 0 {
		resp.DealId = rand.Int()
		resp.Amount = b.Amount
	} else {
		resp.DealId = -1
	}
	return &resp
}

//CheckMerchantId - метод проверки номера терминала
func CheckMerchantId(mid int) bool {
	if mid == 0 {
		return false
	} else {
		return true
	}
}

//CheckLuna - метод проверки Луна номера карты
func CheckLuna(PAN string) bool {
	if PAN == "" {
		return false
	}
	var (
		sum     = 0
		nDigits = len(PAN)
		parity  = nDigits % 2
	)
	for i := 0; i < nDigits; i++ {
		var digit = int(PAN[i] - 48)
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}

//CheckDate - метод проверки даты
func CheckDate(month, year int) bool {
	if month != 0 && year != 0 && 1 <= month && month <= 12 {
		dateNow := time.Now()
		if (year > dateNow.Year()) || (year == dateNow.Year() && month > int(dateNow.Month())) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

//CheckHolder - метод проверки держателя карты
func CheckHolder(holder string) bool {
	if holder != "" {
		name := strings.Split(holder, " ")
		wrong := 0
		for _, v := range name {
			if regexp.MustCompile(`^[A-Z]+$`).MatchString(v) {
				wrong = 0
			} else {
				wrong++
			}
		}
		if wrong == 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

//CheckCVV - метод проверки CVV
func CheckCvv(cvv int) bool {
	if cvv != 0 && cvv >= 100 && cvv <=999 {
		return true
	} else {
		return false
	}
}

//CheckOrderId - метод проверки номера заказа
func CheckOrderID(orderId string) bool {
	if orderId != "" {
		return true
	} else {
		return false
	}
}

//CheckAmount - метод проверки суммы
func CheckAmount(amount int) bool {
	if amount != 0 && amount <= 100 && amount > 0 {
		return true
	} else {
		return false
	}
}
