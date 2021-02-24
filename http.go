package main

import (
				"fmt"
				"log"
				"math"
				"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
				value :=math.Pow(10,2)
				//fmt.Println(w,value)
				//fmt.Println(value)
				fmt.Fprintln(w,value)
				fmt.Fprintln(w,"first web pages")
}				

func main() {
				log.Println("Starting http server on localhost:8080")
				http.HandleFunc("/",handler)
				http.ListenAndServe(":8080",nil)

}