package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const endpoint = "https://script.google.com/macros/s/AKfycbywwDmlmQrNPYoxL90NCZYjoEzuzRcnRuUmFCPzEqG7VdWBAhU/exec"

type Form struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Target string `json:"target"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/translate", func(w http.ResponseWriter, r *http.Request) {
		var form Form
		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		form.Source = "ja"
		form.Target = "en"

		buf := &bytes.Buffer{}
		if err := json.NewEncoder(buf).Encode(form); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		req, err := http.NewRequest(
			http.MethodPost,
			endpoint,
			buf,
		)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer resp.Body.Close()

		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		m := map[string]string{
			"output": string(result),
		}

		if err := json.NewEncoder(w).Encode(&m); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

	})
	log.Println("start http server :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
