package siswa

import (
	"net/http"
	"encoding/json"
	"strconv"
)

type Siswa struct{
	Id 		int 	`json:"id"`
	Nama 	string 	`json:"nama"`
	Kelas 	string 	`json:"kelas"`
	Jurusan string 	`json:"jurusan"`
}

type Message struct{
	Status 	bool 	`json:"status"`
	Message string 	`json:"message"`
}

var db = []*Siswa{}

func init(){
	db = []*Siswa{
		&Siswa{1, "Fajar Fahrurozi", "XII IPA 4", "IPA"},
		&Siswa{2, "Lukman Khakim", "XII IPA 4", "IPA"},
		&Siswa{3, "Dean Febri H.S", "XII IPA 4", "IPA"},
		&Siswa{4, "Gilang Wisnu S", "XII IPA 4", "IPA"},
		&Siswa{5, "Andri Setiawan", "XII IPA 1", "IPA"},
		&Siswa{6, "Bondan Tutus", "XII IPA 1", "IPA"},
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
	var siswa 	Siswa

	siswa.Id = GenerateId()
	siswa.Nama = r.FormValue("nama")
	siswa.Kelas = r.FormValue("kelas")
	siswa.Jurusan = r.FormValue("jurusan")

	db = append(db, &siswa)

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
	kelas 	:= r.FormValue("kelas")
	jurusan := r.FormValue("jurusan")

	for _, value := range db{
		if idx == value.Id {
			value.Nama = nama
			value.Kelas = kelas
			value.Jurusan = jurusan
		}
	}

	message.Status = true
	message.Message = "Sukses update Data"

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