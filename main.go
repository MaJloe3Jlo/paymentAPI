package main

import (
	"encoding/json"
	"fmt"
	"github.com/MaJloe3Jlo/mapisacard_test/lib"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	log.Println("Server started at http://localhost:7000\nPOST: methods /block/, /charge/; contentType: application/json\ncontrol json-requests look in path ./jsons")
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

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
	}
	err = json.Unmarshal(body, &reqBlock)
	if !strings.Contains(string(body), "merchant_contact_id") {
		log.Print("JSON not correct: field merchant_contact_id doesn't exist")
		fmt.Fprint(w, "JSON not correct: field merchant_contact_id doesn't exist")
		return
	} else if !strings.Contains(string(body), "card") {
		log.Print("JSON not correct: field card doesn't exist")
		fmt.Fprint(w, "JSON not correct: field card doesn't exist")
		return
	} else if !strings.Contains(string(body), "pan") {
		log.Print("JSON not correct: field pan doesn't exist")
		fmt.Fprint(w, "JSON not correct: field pan doesn't exist")
		return
	} else if !strings.Contains(string(body), "e_month") {
		log.Print("JSON not correct: field e_month doesn't exist")
		fmt.Fprint(w, "JSON not correct: field e_monnth doesn't exist")
		return
	} else if !strings.Contains(string(body), "e_year") {
		log.Print("JSON not correct: field e_year doesn't exist")
		fmt.Fprint(w, "JSON not correct: field e_year doesn't exist")
		return
	} else if !strings.Contains(string(body), "cvv") {
		log.Print("JSON not correct: field cvv doesn't exist")
		fmt.Fprint(w, "JSON not correct: field cvv doesn't exist")
		return
	} else if !strings.Contains(string(body), "holder") {
		log.Print("JSON not correct: field holder doesn't exist")
		fmt.Fprint(w, "JSON not correct: field holder doesn't exist")
		return
	} else if !strings.Contains(string(body), "order_id") {
		log.Print("JSON not correct: field order_id doesn't exist")
		fmt.Fprint(w, "JSON not correct: field order_id doesn't exist")
		return
	} else if !strings.Contains(string(body), "amount") {
		log.Print("JSON not correct: field amount doesn't exist")
		fmt.Fprint(w, "JSON not correct: field amount doesn't exist")
		return
	}

	if err != nil {
		defer log.Print("JSON isn't correct: " + err.Error() + ". JSON example: " + `{"merchant_contact_id": 1,"card": {"pan": "5469345678901234","e_month": 6,"e_year": 2020,"cvv": 332,"holder": "DMITRIY KLESTOV"},"order_id": "BuyMeTee123","amount": 99}`)
		fmt.Fprint(w, "JSON isn't correct: "+err.Error()+". JSON example: "+`{"merchant_contact_id": 1,"card": {"pan": "5469345678901234","e_month": 6,"e_year": 2020,"cvv": 332,"holder": "DMITRIY KLESTOV"},"order_id": "BuyMeTee123","amount": 99}`)
	}
	defer req.Body.Close()

	respBlock = lib.Validate(reqBlock)
	if respBlock.DealID != -1 {
		log.Printf("Block status: deal ID: %v, amount: %v, error(if nil operation ok): %v", respBlock.DealID, respBlock.Amount, respBlock.Error)
		Block = append(Block, respBlock)
		pretty, err := json.MarshalIndent(respBlock, "", "    ")
		if err != nil {
			log.Println(err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(pretty))
	} else {
		log.Print("JSON request isn't valid")
		fmt.Fprint(w, "JSON request isn't valid")
		respBlock.Error = append(respBlock.Error, err.Error())
	}
}

//charge - метод списания средств с виртуальной карты авторизованных методом block
func charge(w http.ResponseWriter, req *http.Request) {
	log.Print("Charge method.")
	var reqCharge lib.ChargeRequest
	var respCharge lib.ChargeResponse
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
	}
	err = json.Unmarshal(body, &reqCharge)
	if !strings.Contains(string(body), "deal_id") {
		log.Print("JSON not correct: field deal_id doesn't exist")
		fmt.Fprint(w, "JSON not correct: field deal_id doesn't exist")
		return
	} else if !strings.Contains(string(body), "amount") {
		log.Print("JSON not correct: field amount doesn't exist")
		fmt.Fprint(w, "JSON not correct: field amount doesn't exist")
		return
	}
	if err != nil {
		defer log.Print("JSON isn't correct: " + err.Error() + ". JSON example: " + `{"deal_id": 5577006791947779410, "amount": 9}`)
		fmt.Fprint(w, "JSON isn't correct: "+err.Error()+". JSON example: "+`{"deal_id": 5577006791947779410, "amount": 9}`)
		return
	}
	defer req.Body.Close()

	for _, v := range Block {
		if v.DealID != reqCharge.DealID {
			respCharge.Status = "error"
			respCharge.Error = "Charge not working. Do not have this dealID"
			Charge = append(Charge, &respCharge)
			log.Printf("DealID: %v, charge status: %s, error description: %s", v.DealID, respCharge.Status, respCharge.Error)
			pretty, err := json.MarshalIndent(respCharge, "", "    ")
			if err != nil {
				log.Println(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(pretty))
			return
		} else if v.Amount < reqCharge.Amount {
			respCharge.Status = "error"
			respCharge.Error = "Charge not working. Amount of charge is bigger than amount of block"
			Charge = append(Charge, &respCharge)
			log.Printf("DealID: %v, charge status: %s, error description: %s", v.DealID, respCharge.Status, respCharge.Error)
			pretty, err := json.MarshalIndent(respCharge, "", "    ")
			if err != nil {
				log.Println(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(pretty))
			return
		} else {
			go doReq(buf)
			respCharge.Status = <-buf
			if reqCharge.Amount < 0 {
				respCharge.Status = "error"
				respCharge.Error = "Charge not working. Amount < 0"
				log.Printf("DealID: %v, charge status: %s, error description: %s", v.DealID, respCharge.Status, respCharge.Error)
				pretty, err := json.MarshalIndent(respCharge, "", "    ")
				if err != nil {
					log.Println(err.Error())
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, string(pretty))
				return
			} else {
				v.Amount -= reqCharge.Amount
			}
			Charge = append(Charge, &respCharge)
			log.Printf("DealID: %v, charge status: %s, amount balance: %v", v.DealID, respCharge.Status, v.Amount)
			pretty, err := json.MarshalIndent(respCharge, "", "    ")
			if err != nil {
				log.Println(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(pretty))
			return
		}
	}
	log.Println("Cannot find this DealID. Use operation method block before method charge.")
	fmt.Fprint(w, "Cannot find this DealID. Use operation method block before method charge.")
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
