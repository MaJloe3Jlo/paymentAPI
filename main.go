package main

import (
	"encoding/json"
	"fmt"
	"github.com/MaJloe3Jlo/mapisacard_test/lib"
	"log"
	"net/http"
	"time"
)

var (
	Block  []*lib.Block_resp
	Charge []*lib.Charge_req
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/block/", block)
	http.HandleFunc("/charge/", charge)
	log.Fatal(http.ListenAndServe(":7000", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, http.StatusForbidden)
}

func block(w http.ResponseWriter, req *http.Request) {
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
		Block = append(Block, respBlock)
	}

	fmt.Fprint(w, reqBlock)

	fmt.Fprint(w, respBlock)
	fmt.Fprint(w,"\n")

	for _, v := range Block {
		fmt.Fprint(w, v)
		fmt.Fprint(w,"\n")
	}
}

func charge(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var reqCharge lib.Charge_req
	var respCharge *lib.Charge_resp

	err := decoder.Decode(&reqCharge)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	buf := make(chan string)

	for _, v := range Block {
		if v.DealId == reqCharge.DealId && v.Amount >= reqCharge.Amount {
			go doReq(buf)
			time.Sleep(1 * time.Second)
			status := <-buf
			respCharge.Status = status
		} else {
			respCharge.Status = "error"
			respCharge.Error = "Charge not working"
		}
	}
	fmt.Fprint(w, respCharge)
}

func doReq(buf chan string) {

	req, err := http.Get("https://ya.ru")
	if err != nil {
		log.Println(err)
	}
	buf<-req.Status
}