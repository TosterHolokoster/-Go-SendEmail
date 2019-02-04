package main

import (
	"encoding/json"
	"net/http"
	"net/smtp"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/proccess-form", SendMail)
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("resourses/image"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

//SendMail send email use SMTP
func SendMail(w http.ResponseWriter, r *http.Request) {
	type Anser struct {
		Result bool `json:result`
	}

	r.ParseForm()
	gmail := "GoSendEmailTest@gmail.com"
	pass := "passForTest0000"
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	mess := r.PostFormValue("message")
	to := "alexnemchinov3541@gmail.com"
	msg := "From: " + name + " <" + email + ">\n" +
		"To: " + to + "\n" +
		"Subject: It's Work!\n\n" +
		"Name: " + name + "\n" +
		"Email: " + email + "\n" +
		"Message: " + mess + "\n"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", gmail, pass, "smtp.gmail.com"),
		gmail, []string{to}, []byte(msg))

	if err != nil {
		data := Anser{Result: false}
		json.NewEncoder(w).Encode(data)
		return
	}
	data := Anser{Result: true}
	json.NewEncoder(w).Encode(data)
}
