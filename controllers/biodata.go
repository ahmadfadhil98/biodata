package controllers

import (
	"biodata/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func GetBiodata(param string) models.Biodata {

	absen, err := strconv.ParseInt(param, 10, 0)

	if err != nil {
		fmt.Println("Parameter must be integer")
		os.Exit(1)
	}

	if absen < 1 {
		fmt.Println("Parameter must be greater than 0")
		os.Exit(1)
	}

	var result models.Biodata

	jsonFile, err := os.Open("../biodata/datas/biodata.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var biodata models.Biodatas

	json.Unmarshal(byteValue, &biodata)

	result = biodata.Biodata[absen-1]

	fmt.Println(biodata.Biodata[absen-1].Nama)
	fmt.Println(biodata.Biodata[absen-1].Alamat)
	fmt.Println(biodata.Biodata[absen-1].Pekerjaan)
	fmt.Println(biodata.Biodata[absen-1].Alasan)

	return result
}
