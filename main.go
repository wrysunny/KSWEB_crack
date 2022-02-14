package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strconv"
	"time"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("这是一个私人消息API服务器."))
	})
	r.Post("/service/serKey/v2/do.php", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(EncryptMsg("pass;2")))
	})
	r.Post("/service/24082016/checkTrial.php", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(EncryptMsg(strconv.FormatInt(time.Now().Unix(), 10) + " 4099737600")))
	})

	http.ListenAndServe(":3000", r)
}
