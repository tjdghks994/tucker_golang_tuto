package main

import (
	"fmt"
	"net/http"

	"github.com/tjdghks994/tucker_golang_tuto/web/myapp"
)

func main() {
	err := http.ListenAndServe(":3000", myapp.NewHttpHandler())
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
