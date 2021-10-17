package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
)

//func sayhelloGolang(w http.ResponseWriter, r *http.Request) {
//      r.ParseForm() //解析参数，默认是不会解析的
//      fmt.Println("path", r.URL.Path)
//      w.Write([]byte("Hello Golang"))
//}

func setHeader(w http.ResponseWriter, r *http.Request) {
        //      fmt.Println("打印Header参数列表：")
        ver := os.Getenv("VERSION")
        if len(r.Header) > 0 {
                for k, v := range r.Header {
                        w.Header().Set(k, v[0])
                }
                fmt.Println(r.RemoteAddr)

        }
        w.Header().Add("VERSION", ver)

}
func healthz(w http.ResponseWriter, r *http.Request) {
        ver := os.Getenv("VERSION")
        if len(r.Header) > 0 {
                for k, v := range r.Header {
                        w.Header().Set(k, v[0])
                }
        }
        w.Header().Add("VERSION", ver)
        fmt.Println(r.RemoteAddr)
        w.Write([]byte("200 OK"))
}

func main() {
        http.HandleFunc("/", setHeader) //设置访问的路由
        http.HandleFunc("/healthz", healthz)
        err := http.ListenAndServe(":8080", nil) //设置监听的端口
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }

}