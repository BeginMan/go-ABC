package main

import (
	"log"
	"net/http"

	"handlers"
	"nets"
)

//Go的包加载机制允许我们在init()函数中做这样的事情, init()会在main()函数之前执行。
func init() {
	nets.Init()
}

func main() {
	mux := http.NewServeMux()
	nets.StaticDirHandler(mux, "/assets/", "./public", 0)

	mux.HandleFunc("/", nets.SafeHandler(handlers.ListHandler))
	mux.HandleFunc("/upload", nets.SafeHandler(handlers.UploadHandler))
	mux.HandleFunc("/view", nets.SafeHandler(handlers.ViewHanlder))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}
