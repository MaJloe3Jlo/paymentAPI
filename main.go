package main

import (
	"encoding/json"
	"fmt"
	"github.com/MaJloe3Jlo/mapisacard_test/lib"
	"log"
	"net/http"
)

//Переменные слайсов запросов Block и Charge
var (
	Block  []*lib.Block_resp
	Charge []*lib.Charge_resp
)


//main - задает пути приложения и выводит информацию по работе с приложением
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/block/", block)
	http.HandleFunc("/charge/", charge)
	log.Println("POST: methods /block/, /charge/; contentType: application/json")
	log.Println("control json-requests look in path ./jsons")
	log.Println("to test app you can use curl or postman")
	log.Fatal(http.ListenAndServe(":7000", nil))
}


//index - метод заглушка всегда возвращает 403
func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, http.StatusForbidden)
}

//block - метод блокирует средства для списание на виртуальной карте
func block(w http.ResponseWriter, req *http.Request) {
	log.Print("Block method. ")
	decoder := json.NewDecoder(req.Body)

	var reqBlock lib.Block_req
	var respBlock *lib.Block_resp

	err := decoder.Decode(&reqBlock)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	respBlock = lib.Validate(reqBlock)
	if respBlock.DealId != -1 {
		log.Printf("Block status: deal ID: %v, amount: %v, error(if nil operation ok): %v",respBlock.DealId, respBlock.Amount, respBlock.Error)
		Block = append(Block, respBlock)
	}
}

//charge - метод списания средств с виртуальной карты авторизованных методом block
func charge(w http.ResponseWriter, req *http.Request) {
	log.Print("Charge method.")
	decoder := json.NewDecoder(req.Body)

	var reqCharge lib.Charge_req
	var respCharge lib.Charge_resp

	err := decoder.Decode(&reqCharge)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	buf := make(chan string, 10)

	for _, v := range Block {
		if v.DealId != reqCharge.DealId {
			respCharge.Status = "error"
			respCharge.Error = "Charge not working. Do not have this dealID"
			Charge = append(Charge, &respCharge)
			log.Printf("DealID: %v, charge status: %s, error description: %s", v.DealId, respCharge.Status, respCharge.Error)
		} else if v.Amount < reqCharge.Amount {
			respCharge.Status = "error"
			respCharge.Error = "Charge not working. Amount of charge is bigger than amount of block"
			Charge = append(Charge, &respCharge)
			log.Printf("DealID: %v, charge status: %s, error description: %s", v.DealId, respCharge.Status, respCharge.Error)
		} else {
			go doReq(buf)
			respCharge.Status = <-buf
			v.Amount -= reqCharge.Amount
			Charge = append(Charge, &respCharge)
			log.Printf("DealID: %v, charge status: %s, amount balance: %v", v.DealId, respCharge.Status, v.Amount)
		}
	}
}

//doReq - метод фоновой отправки запроса к ya.ru
func doReq(buf chan string) {

	req, err := http.Get("https://ya.ru")
	if err != nil {
		log.Println(err)
	}
	buf <- req.Status
}