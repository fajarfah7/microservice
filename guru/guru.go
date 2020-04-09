package guru

import (
	"net/http"
	"encoding/json"
	"strconv"
)

type Guru struct{
	Id 		int 	`json:"id"`
	Nama 	string 	`json:"nama"`
	Mapel 	string 	`json:"mapel"`
}

type Message struct{
	Status 	bool 	`json:"status"`
	Message string 	`json:"message"`
}

var db = []*Guru{}

func init(){
	db = []*Guru{
		&Guru{1, "Santi Nurmalitasari", "Bahasa Indonesia"},
		&Guru{2, "Aditya A.P", "Sejarah"},
		&Guru{3, "Arif Frebianto", "Seni Musik"},
		&Guru{4, "Ika Listiani", "Pendidikan Kewarganegaraan"},
	}
}

// func main() {
// 	mux := http.DefaultServeMux

// 	mux.HandleFunc("/", List)
// 	mux.HandleFunc("/insert", Insert)
// 	mux.HandleFunc("/update", Update)
// 	mux.HandleFunc("/delete", Delete)

// 	server := new (http.Server)

// 	server.Addr = (":3000")
// 	server.ListenAndServe()
// }

func List(w http.ResponseWriter, r *http.Request){
	data, _ := json.Marshal(db)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Insert(w http.ResponseWriter, r *http.Request){
	var message Message
	var guru 	Guru

	guru.Id = GenerateId()
	guru.Nama = r.FormValue("nama")
	guru.Mapel = r.FormValue("mapel")

	db = append(db, &guru)

	message.Status = true
	message.Message = "Sukses input data"

	data, _ := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Update(w http.ResponseWriter, r *http.Request){
	var message Message

	idForm := r.FormValue("id")
	idx, _ := strconv.Atoi(idForm)
	nama 	:= r.FormValue("nama")
	mapel 	:= r.FormValue("mapel")

	for _, value := range db{
		if idx == value.Id {
			value.Nama = nama
			value.Mapel = mapel
		}
	}

	message.Status = true
	message.Message = "Sukses Update Data"

	data, _ := json.Marshal(message)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Delete(w http.ResponseWriter, r *http.Request){
	var message Message

	idForm := r.FormValue("id")
	idx, _ := strconv.Atoi(idForm)
	var k int

	for key, val := range db{
		if val.Id == idx {
			k = key
		}
	}
	db[k] = db[len(db)-1]
	db = db[:(len(db)-1)]

	message.Status = true
	message.Message = "Sukses hapus data"

	data, _ := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GenerateId()int{
	var idx = 0

	for _, value := range db{
		if idx < value.Id {
			idx = value.Id
		}
	}

	return idx+1
}