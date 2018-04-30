package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"io/ioutil"
	"encoding/json"
	"github.com/MaJloe3Jlo/mapisacard_test/lib"
)

func main() {
	rr := newPathResolver()
	rr.Add("/", index)
	rr.Add("POST /block", block)
	rr.Add("GET /charge", charge)
	http.ListenAndServe(":7000", rr)
}

func newPathResolver() *regexResolver {
	return &regexResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

type regexResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp
}

func (r *regexResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, err := regexp.Compile(regex)
	if err != nil {
		log.Fatal(err)
	}
	r.cache[regex] = cache
}

func (r *regexResolver) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) == true {
			handlerFunc(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, http.StatusForbidden)
}

func block(w http.ResponseWriter, req *http.Request) {
	query, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	bl := lib.Block_req{}
	json.Unmarshal(query, bl)
	fmt.Println(bl)
}

func charge(w http.ResponseWriter, req *http.Request) {

}
