package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handlePresensi(w http.ResponseWriter, r *http.Request) {
	url := "https://bristars.bri.co.id/bristars/presensi/online"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	// Set your form fields here
	formFields := map[string]string{
		"save":             "save",
		"mt":               "1",
		"latitude":         "-6.3028584",
		"longtitude":       "106.8216535",
		"timestamp":        "1700009873123",
		"tanggal_presensi": "2023-11-15 07:57:53",
		"jenis_presensi":   "datang",
	}

	for key, value := range formFields {
		_ = writer.WriteField(key, value)
	}

	err := writer.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Add("Cookie", "visid_incap_2649695=sJ1RoDMMTyqgg14fG4fG7Sux0GQAAAAAQUIPAAAAAABYVL70J6dXOy3mfA1A1Dyz; _ga_L4FXYDYCV6=GS1.1.1691398453.1.1.1691398465.0.0.0; visid_incap_2649669=CfKr0U/+RgqFbNQndx4+GU/S5WQAAAAAQUIPAAAAAAC12kncR682UXOizImjNPI2; visid_incap_2611317=2iSCNcENTymEIMXzyVKIK+A46GQAAAAAQUIPAAAAAAAHoMLWNaeEdqs7kW5Ti8wI; _ga_NE8WYYCWD6=GS1.1.1693364169.22.0.1693364169.0.0.0; _ga_E1N3J7LSCH=GS1.3.1694577801.2.0.1694577801.0.0.0; _ga_LFFEQ22SW8=GS1.1.1694577801.2.0.1694579366.0.0.0; _ga_4D8HKV1L3T=GS1.1.1697084641.24.0.1697084641.60.0.0; _ga=GA1.3.810752856.1690868603; PHPSESSID=9ebhuq5j45gr2bda5299h5q7k7; PHPSESSID=a2km5ld5a032rjq7ccb1lfgcr4")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	w.Write(body)
}

func main() {
	router := mux.NewRouter()

	// Define your route
	router.HandleFunc("/presensi", handlePresensi).Methods("POST")

	// Start the server
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Server is running on :8080")
	log.Fatal(server.ListenAndServe())
}
