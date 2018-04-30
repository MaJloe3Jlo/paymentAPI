package lib

import (
	"time"
	"regexp"
	"strings"
)

//CheckLuna - метод проверки Луна номера карты
func CheckLuna(PAN string) bool {
	if PAN != "" {
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
	} else {
		return false
	}
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
