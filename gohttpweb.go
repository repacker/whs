package main

import(
	"io"
	"log"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"Hello,World!I am whs!")
}

func main(){
	http.HandleFunc("/hellowhs",helloHandler)
	fmt.Print("starting...")
	err:=http.ListenAndServe(":8080",nil)
	fmt.Print("end...")
	if err != nil{
	    log.Fatal("ListenAndServer:",err.Error())
	    fmt.Print("error")
	}
}






