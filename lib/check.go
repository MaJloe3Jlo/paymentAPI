package lib

import (
	"regexp"
	"strings"
	"time"
	"math/rand"
)

//Метод проверки всех полей метода Block
func Validate(b Block_req) (*Block_resp) {
	var resp Block_resp

	if CheckMercantId(b.MerchantContactId) == false {
		resp.error = append(resp.error, "Error terminal ID; ")
	}
	if CheckLuna(b.Card.Pan) == false {
		resp.error = append(resp.error, "Error PAN number; ")
	}
	if CheckDate(b.Card.EMonth, b.Card.EYear) == false {
		resp.error = append(resp.error, "Error date of card; ")
	}
	if CheckCvv(b.Card.Cvv) == false {
		resp.error = append(resp.error, "Error CVV; ")
	}
	if CheckHolder(b.Card.Holder) == false {
		resp.error = append(resp.error, "Error holder name; ")
	}
	if CheckOrderID(b.OrderId) == false {
		resp.error = append(resp.error, "Error order ID; ")
	}
	if CheckAmount(b.Amount) == false {
		resp.error = append(resp.error, "Error amount; ")
	}

	if len(resp.error) == 0 {
		resp.dealId = rand.New(rand.NewSource(99)).Int()
	} else {
		resp.dealId = -1
	}

	return &resp
}

//CheckMercantId- метод проверки номера терминала
func CheckMercantId(mid int) bool {
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
	if month != 0 || year != 0 {
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
	if cvv == 0 {
		return false
	} else {
		return true
	}
}

//CheckOrderId - метод проверки номера заказа
func CheckOrderID(orderId string) bool {
	if orderId == "" {
		return false
	} else {
		return true
	}
}

//CheckAmount - метод проверки суммы
func CheckAmount(amount int) bool {
	if amount == 0 {
		return false
	} else {
		return true
	}
}
