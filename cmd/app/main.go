package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"


	"github.com/go-chi/chi"
)

func main() {
	// 最低限動くwebサーバー
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	svr := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("server started")

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	log.Println("server stopped")
}
