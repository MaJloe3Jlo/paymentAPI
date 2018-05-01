package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/MaJloe3Jlo/mapisacard_test/lib"
	"encoding/json"
)

var (
	Block []*lib.Block_req
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

	var req_b lib.Block_req
	var resp_b *lib.Block_resp

	err := decoder.Decode(&req_b)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	resp_b = lib.Validate(req_b)

	fmt.Fprint(w, resp_b)
	fmt.Fprint(w, req_b)

}


func charge(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "charge")
}
