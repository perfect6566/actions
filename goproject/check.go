package main

import ("log"
"net/http"
"html/template"
)
type Info struct{
Name string
}
func sayhelloGolang(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //解析参数，默认是不会解析的
    log.Println("path", r.URL.Path)
    w.Write([]byte("Hello Golang"))
}


func sayhello(w http.ResponseWriter, r *http.Request) {
   r.ParseForm()  //解析参数，默认是不会解析的
   info:=&Info{}
   info.Name="Frank"
   t, err := template.ParseFiles("static/index.html")
   if err != nil {
      http.Error(w, err.Error(),http.StatusInternalServerError)
      return
   }
   t.Execute(w, info)
   return
}

func main(){

log.Println("Welcome go")
http.HandleFunc("/", sayhelloGolang) //设置访问的路由
    
http.HandleFunc("/hello", sayhello) //设置访问的路由
   err := http.ListenAndServe(":9090", nil) //设置监听的端口
   if err != nil {
      log.Fatal("ListenAndServe: ", err)
   }
}
  
