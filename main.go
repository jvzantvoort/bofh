package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	DEFAUL_PORT int = 2609
)

func GetExcuse(w http.ResponseWriter, r *http.Request) {
	var retv string
	box := packr.NewBox("./templates")
	excuses_text, err := box.FindString("excuses.txt")
	if err != nil {
		log.Fatal(err)
	}
	excuses_lines := strings.Split(strings.Replace(excuses_text, "\r\n", "\n", -1), "\n")
	rand.Seed(time.Now().Unix())
	retv = excuses_lines[rand.Intn(len(excuses_lines))]
	fmt.Fprintf(w, retv+"\n")
}

func main() {
	http.HandleFunc("/", GetExcuse)
	portstr := fmt.Sprintf(":%d", DEFAUL_PORT)
	http.ListenAndServe(portstr, nil)
}
