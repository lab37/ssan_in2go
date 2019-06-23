package main

import (
	"net/http"
	"time"
)

func main() {
	p("进销存系统", version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", index)
	mux.HandleFunc("/jxc", jxc)
	mux.HandleFunc("/webchat", webchat)
	mux.HandleFunc("/htmlturn", htmlTurn)
	mux.HandleFunc("/chgaccount", chgAccount)
	mux.HandleFunc("/xls2cat", xls2cat)
	mux.HandleFunc("/discount", discount)
	mux.HandleFunc("/tax", tax)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_jxc.go
	mux.HandleFunc("/handle_products", handleProducts)
	mux.HandleFunc("/handle_customers", handleCustomers)
	mux.HandleFunc("/handle_contracts", handleContracts)
	mux.HandleFunc("/handle_instocks", handleInStocks)
	mux.HandleFunc("/handle_outstocks", handleOutStocks)
	mux.HandleFunc("/handle_invoices", handleInvoices)
	mux.HandleFunc("/handle_payments", handlePayments)
	mux.HandleFunc("/get_products_ns", getProductNS)
	mux.HandleFunc("/get_cstmname", getCustomerName)
	mux.HandleFunc("/handle_stocks", handleStocks)
	mux.HandleFunc("/handle_debts", handleDebts)
	mux.HandleFunc("/handle_onway_products", handleOnWayProducts)
	mux.HandleFunc("/handle_onway_invoices", handleOnWayInvoices)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/updateaccount", updateAccount)
	mux.HandleFunc("/updatepassword", updatePassword)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
