package main

import (
	"fmt"
	"net/http"

	helper "./helpers"
)

func main() {
	uName, email, pwd, pwdConfirm := "", "", "", ""
	mux := http.NewServeMux()
	//Signup
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")
		uNameCheck := helper.IsEmpty(uName)
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		pwdConfirmCheck := helper.IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "boş bırakılan alan")
			return
		}
		if pwd == pwdConfirm {
			//normalde burada açık olan veritabanı bağlantımız üzerinden gelen datayı dbye kaydedecektik
			fmt.Fprintf(w, "Registration successfull.")
		} else {
			fmt.Fprintf(w, "Password information must be same")
		}
	})

	//Login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email = r.FormValue("email")
		pwd = r.FormValue("pwd")
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "boş alan bırakmayınız")
			return
		}
		//normalde burada else blokumuz olup db sorgusunu yapacaktık
		//daha sonrasında gerekli url'ye yönlendirme yapacaktık
		//ama biz burada simulasyonvari yapıyoruz
		dbPwd := "12345"
		dbEmail := "mustafa.yildirim93@erzurum.edu.tr"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintf(w, "giriş başarılı")
		} else {
			fmt.Fprintf(w, "giriş başarısız")
		}
	})
	http.ListenAndServe(":8080", mux)

}

/*
	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}
*/
