package server

import (
	"health-app/server/controller"
	"log"
	"net/http"
)

func Run(serverPort string) {
	uh := controller.NewUserHandler()
	http.Handle("/users", uh)
	http.Handle("/users/", uh)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!\n"))
	})
	log.Println("Serving Go server on http://localhost" + serverPort)
	log.Fatal(http.ListenAndServe(serverPort, nil))
}
