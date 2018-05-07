package main

import (
	"encoding/json"
	"fmt"
	"github.com/MaJloe3Jlo/mapisacard_test/lib"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

//Переменные слайсов запросов Block и Charge
var (
	Block  []*lib.BlockResponse
	Charge []*lib.ChargeResponse
	buf    = make(chan string, 10)
)

//main - задает пути приложения и выводит информацию по работе с приложением
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/block/", block)
	http.HandleFunc("/charge/", charge)
	log.Println("Server started at http://localhost:7000")
	log.Println("POST: methods /block/, /charge/; contentType: application/json")
	log.Println("control json-requests look in path ./jsons")
	log.Fatal(http.ListenAndServe(":7000", nil))
}

//index - метод заглушка всегда возвращает 403
func index(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprint(w, http.StatusForbidden)
}

//block - метод блокирует средства для списание на виртуальной карте
func block(w http.ResponseWriter, req *http.Request) {
	log.Print("Block method. ")
	var reqBlock lib.BlockRequest
	var respBlock *lib.BlockResponse
	var m sync.Mutex

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(w, err.Error())
	}

	state := lib.CheckBody(body, true)

	if state != "" {
		log.Println(state)
		fmt.Fprint(w, state)
		return
	}

	errUnmarshal := json.Unmarshal(body, &reqBlock)

	if errUnmarshal != nil {
		log.Print("JSON isn't correct: " + errUnmarshal.Error() + ". JSON example: " + `{"merchant_contact_id": 1,"card": {"pan": "5469345678901234","e_month": 6,"e_year": 2020,"cvv": 332,"holder": "DMITRIY KLESTOV"},"order_id": "BuyMeTee123","amount": 99}`)
		fmt.Fprint(w, "JSON isn't correct: "+errUnmarshal.Error()+". JSON example: "+`{"merchant_contact_id": 1,"card": {"pan": "5469345678901234","e_month": 6,"e_year": 2020,"cvv": 332,"holder": "DMITRIY KLESTOV"},"order_id": "BuyMeTee123","amount": 99}`)
		return
	}
	req.Body.Close()

	respBlock = lib.Validate(reqBlock)

	if respBlock.DealID != -1 {
		log.Printf("Block status: deal ID: %v, amount: %v, error(if nil operation ok): %v", respBlock.DealID, respBlock.Amount, respBlock.Error)
		m.Lock()
		Block = append(Block, respBlock)
		m.Unlock()
		pretty, err := json.MarshalIndent(respBlock, "", "    ")
		if err != nil {
			log.Println(err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(pretty))
	} else {
		log.Print("JSON request isn't valid")
		pretty, errMI := json.MarshalIndent(respBlock, "", "    ")
		if errMI != nil {
			log.Println(errMI.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(pretty))
	}
}

//charge - метод списания средств с виртуальной карты авторизованных методом block
func charge(w http.ResponseWriter, req *http.Request) {
	log.Print("Charge method.")
	var reqCharge lib.ChargeRequest
	var respCharge lib.ChargeResponse
	var m sync.Mutex

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
	}

	state := lib.CheckBody(body, false)
	if state != "" {
		log.Println(state)
		fmt.Fprint(w, state)
		return
	}

	errUnmarshal := json.Unmarshal(body, &reqCharge)

	if errUnmarshal != nil {
		defer log.Print("JSON isn't correct: " + errUnmarshal.Error() + ". JSON example: " + `{"deal_id": 55779410, "amount": 9}`)
		fmt.Fprint(w, "JSON isn't correct: "+errUnmarshal.Error()+". JSON example: "+`{"deal_id": 55779410, "amount": 9}`)
		return
	}
	req.Body.Close()

	stateFind, amount := findBlock(reqCharge.DealID)

	if stateFind == false {
		respCharge.Status = "error"
		respCharge.Error = "Charge not working. Do not have this dealID"
		m.Lock()
		Charge = append(Charge, &respCharge)
		m.Unlock()
		log.Printf("DealID: %v, charge status: %s, error description: %s", reqCharge.DealID, respCharge.Status, respCharge.Error)
		pretty, err := json.MarshalIndent(respCharge, "", "    ")
		if err != nil {
			log.Println(err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(pretty))
	} else {
		if amount < reqCharge.Amount {
			respCharge.Status = "error"
			respCharge.Error = "Charge not working. Amount of charge is bigger than amount of block"
			m.Lock()
			Charge = append(Charge, &respCharge)
			m.Unlock()
			log.Printf("DealID: %v, charge status: %s, error description: %s", reqCharge.DealID, respCharge.Status, respCharge.Error)
			pretty, err := json.MarshalIndent(respCharge, "", "    ")
			if err != nil {
				log.Println(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(pretty))
		} else {
			go doReq(buf)
			respCharge.Status = <-buf
			if reqCharge.Amount < 0 {
				respCharge.Status = "error"
				respCharge.Error = "Charge not working. Amount < 0"
				log.Printf("DealID: %v, charge status: %s, error description: %s", reqCharge.DealID, respCharge.Status, respCharge.Error)
				pretty, err := json.MarshalIndent(respCharge, "", "    ")
				if err != nil {
					log.Println(err.Error())
				}
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, string(pretty))
				return
			}
			am := amountMinus(reqCharge.DealID, reqCharge.Amount)
			m.Lock()
			Charge = append(Charge, &respCharge)
			m.Unlock()

			log.Printf("DealID: %v, left amount: %v charge status: %s", reqCharge.DealID, am, respCharge.Status)
			pretty, err := json.MarshalIndent(respCharge, "", "    ")
			if err != nil {
				log.Println(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(pretty))
		}
	}
}

//doReq - метод фоновой отправки запроса к ya.ru
func doReq(buf chan string) {

	req, err := http.Get("https://ya.ru")
	if err != nil {
		log.Println(err)
	}
	log.Println("Request to ya.ru " + req.Status)
	buf <- req.Status
}

//findBlock - поиск среди всех совершенных операций Block
func findBlock(dealCharge int) (state bool, amount int) {
	for _, v := range Block {
		if v.DealID == dealCharge {
			return true, v.Amount
		}
	}
	return state, amount
}

//amountMinus - проверка суммы и уменьшение
func amountMinus(dealID, amount int) (am int) {
	for _, v := range Block {
		if v.DealID == dealID {
			v.Amount -= amount
			am = v.Amount
		}
	}
	return am
}
