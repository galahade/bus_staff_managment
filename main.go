package main

import (
	"net/http"
	"io"
	"log"
	"fmt"
	"html/template"
	"time"
	"crypto/md5"
	"strconv"
	"github.com/galahade/bus_staff_managment/session"
	. "github.com/galahade/bus_staff_managment/controller"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
//	go globalSessions.GC()

}

func main() {

	mux := http.NewServeMux()

	//mux.Handle("/bus", http.HandlerFunc(handler))
	//mux.Handle("/", http.HandlerFunc(HelloServer))
	mux.HandleFunc("/bus", handler)
	mux.HandleFunc("/", HelloServer)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/drivers", ListDriver)

	log.Fatal(http.ListenAndServe(":8000", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "BUS station")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	sess := globalSessions.SessionStart(w,r)
	r.ParseForm()
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		sess.Get("username")
		log.Println(t.Execute(w, token))

	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

