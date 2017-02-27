package controller

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
	"github.com/galahade/bus_staff_managment/util"
	"github.com/galahade/bus_staff_managment/service"
)

func ListDriver(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/driver.tmpl")
		util.CheckErr(err)
		staffs :=  service.GetAllDrivers()
		w.Header().Set("Content-Type", "text/html")
		log.Println(t.Execute(w, staffs))

	} else {
		http.Redirect(w, r, "/", 302)
	}
}

