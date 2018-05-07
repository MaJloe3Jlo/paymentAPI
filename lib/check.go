package lib

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func CheckBody(body []byte, flag bool) string {

	switch flag {
	case true:
		if !strings.Contains(string(body), "merchant_contact_id") {
			return "JSON not correct: field merchant_contact_id doesn't exist"
		} else if !strings.Contains(string(body), "card") {
			return "JSON not correct: field card doesn't exist"
		} else if !strings.Contains(string(body), "pan") {
			return "JSON not correct: field pan doesn't exist"
		} else if !strings.Contains(string(body), "e_month") {
			return "JSON not correct: field e_monnth doesn't exist"
		} else if !strings.Contains(string(body), "e_year") {
			return "JSON not correct: field e_year doesn't exist"
		} else if !strings.Contains(string(body), "cvv") {
			return "JSON not correct: field cvv doesn't exist"
		} else if !strings.Contains(string(body), "holder") {
			return "JSON not correct: field holder doesn't exist"
		} else if !strings.Contains(string(body), "order_id") {
			return "JSON not correct: field order_id doesn't exist"
		} else if !strings.Contains(string(body), "amount") {
			return "JSON not correct: field amount doesn't exist"
		}
	case false:
		if !strings.Contains(string(body), "deal_id") {
			return "JSON not correct: field deal_id doesn't exist"
		} else if !strings.Contains(string(body), "amount") {
			return "JSON not correct: field amount doesn't exist"
		}
	}
	return ""
}

//Метод проверки всех полей метода Block
func Validate(b BlockRequest) *BlockResponse {
	var resp BlockResponse

	if CheckMerchantID(b.MerchantContactID) == false {
		resp.Error = append(resp.Error, "Error terminal ID; ")

	}
	if CheckLuhn(b.Card.PAN) == false {
		resp.Error = append(resp.Error, "Error PAN number; ")
	}
	if CheckDate(b.Card.EMonth, b.Card.EYear) == false {
		resp.Error = append(resp.Error, "Error date of card; ")
	}
	if CheckCVV(b.Card.CVV) == false {
		resp.Error = append(resp.Error, "Error CVV; ")
	}
	if CheckHolder(b.Card.Holder) == false {
		resp.Error = append(resp.Error, "Error holder name; ")
	}
	if CheckOrderID(b.OrderID) == false {
		resp.Error = append(resp.Error, "Error order ID; ")
	}
	if CheckAmount(b.Amount) == false {
		resp.Error = append(resp.Error, "Error amount; ")
	}

	if len(resp.Error) == 0 {
		resp.DealID = rand.New(rand.NewSource(time.Now().UnixNano())).Int()
		resp.Amount = b.Amount
	} else {
		resp.DealID = -1
	}
	return &resp
}

//CheckMerchantId - метод проверки номера терминала
func CheckMerchantID(mid int) (state bool) {
	if mid != 0 && mid > 0 {
		state = true
	}
	return
}

//CheckLuhn - метод проверки Луна номера карты
func CheckLuhn(PAN string) (state bool) {
	if PAN == "" {
		return
	}
	for _, v := range PAN {
		if (v >= 'A' && v <= 'Z') || v == ' ' {
			return
		}
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
func CheckDate(month, year int) (state bool) {
	if month != 0 && year != 0 && 1 <= month && month <= 12 {
		dateNow := time.Now()
		if (year > dateNow.Year()) || (year == dateNow.Year() && month >= int(dateNow.Month())) {
			state = true
		} else {
			state = false
		}
	}
	return
}

//CheckHolder - метод проверки держателя карты
func CheckHolder(holder string) (state bool) {
	if holder != "" {
		for _, v := range holder {
			if (v >= 'A' && v <= 'Z') || v == ' ' {
				state = true
			} else {
				state = false
				return
			}
		}
	}
	return state
}

//CheckCVV - метод проверки CVV
func CheckCVV(CVV int) (state bool) {
	if CVV != 0 && (len(strconv.Itoa(CVV)) == 3 || len(strconv.Itoa(CVV)) == 4) {
		return true
	}
	return
}

//CheckOrderId - метод проверки номера заказа
func CheckOrderID(orderID string) (state bool) {
	if orderID != "" {
		return true
	}
	return
}

//CheckAmount - метод проверки суммы
func CheckAmount(amount int) (state bool) {
	if amount != 0 && amount > 0 {
		return true
	}
	return
}
