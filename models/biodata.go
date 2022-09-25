package models

type Biodatas struct {
	Biodata []Biodata `json:"biodata"`
}

type Biodata struct {
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Pekerjaan string `json:"pekerjaan"`
	Alasan    string `json:"alasan_memiilih_golang"`
}
