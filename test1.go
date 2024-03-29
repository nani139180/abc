package main
import (
    "strings"
    "fmt"
    "net/http"
    "log"
)
func sayHelloName(w http.ResponseWriter, r *http.Request){
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path: ", r.URL.Path)
    fmt.Println("scheme: ", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form{
        fmt.Println("key: ", k)
        fmt.Println("val: ", strings.Join(v, " "))
    }
    fmt.Fprintf(w, "你好go")
}

func Start(){
    http.HandleFunc("/", sayHelloName)
    err := http.ListenAndServe(":9191", nil)
    if err != nil{
        log.Fatal("ListenAndServe: ", err)
    }
}
func main(){
    Start()
}

