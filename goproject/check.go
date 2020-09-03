package main

import ("log"
        "net/http"
)


func main(){

log.Println("Welcome go")
  
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<H1>hello world!"))
    })

    http.ListenAndServe(":8080", nil)


}
