package main

import (
	"net/http"
	guru "./guru"
	siswa "./siswa"
)

func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/guru/", guru.List)
	mux.HandleFunc("/guru/insert", guru.Insert)
	mux.HandleFunc("/guru/update", guru.Update)
	mux.HandleFunc("/guru/delete", guru.Delete)

	mux.HandleFunc("/siswa/", siswa.List)
	mux.HandleFunc("/siswa/insert", siswa.Insert)
	mux.HandleFunc("/siswa/update", siswa.Update)
	mux.HandleFunc("/siswa/delete", siswa.Delete)


	server := new (http.Server)

	server.Addr = (":3000")
	server.ListenAndServe()
}