package main

import (
	"ssat/data"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

func handleProducts(writer http.ResponseWriter, request *http.Request){
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		operation := strings.TrimFunc(request.PostFormValue("operation"), unicode.IsSpace)
		tt := strings.TrimFunc(request.PostFormValue("tt"), unicode.IsSpace)
		switch operation {
		case "select":
			writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
			if tt == "all" {
				rsts, err := data.GetAllProducts()
				if err != nil {
					danger("When select all products get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				args:=make(map[string]string)
				args["prdtName"] = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
				args["specific"] = strings.TrimFunc(request.PostFormValue("Specific"), unicode.IsSpace)
				args["inventor"] = strings.TrimFunc(request.PostFormValue("Inventor"), unicode.IsSpace)
				rsts, err := data.SelectProducts(args)
				if err != nil {
					danger("When select some products get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "insert":
			product := data.Product{}
			product.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			product.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			product.Inventor = strings.TrimFunc(request.PostFormValue("inventor"), unicode.IsSpace)
			product.Unit = strings.TrimFunc(request.PostFormValue("unit"), unicode.IsSpace)
			product.IvType = strings.TrimFunc(request.PostFormValue("ivtype"), unicode.IsSpace)
			product.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.InsertProduct(&product)
			if err != nil {
				danger("When insert the product get an error:", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "update":
		    product := data.Product{}
			product.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			product.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			product.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			product.Inventor = strings.TrimFunc(request.PostFormValue("inventor"), unicode.IsSpace)
			product.Unit = strings.TrimFunc(request.PostFormValue("unit"), unicode.IsSpace)
			product.IvType = strings.TrimFunc(request.PostFormValue("ivtype"), unicode.IsSpace)
			product.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.UpdateProduct(&product)
			if err != nil {
				danger("When update the product get an error", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "delete":
			id, _ := strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := data.DeleteProduct(id)
			if err != nil {
				danger("When delete the product get an error", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		}
	}
}




func handleCustomers(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		operation := strings.TrimFunc(request.PostFormValue("operation"), unicode.IsSpace)
		tt := strings.TrimFunc(request.PostFormValue("tt"), unicode.IsSpace)
		switch operation {
		case "select":
			writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
			if tt == "all" {
				rsts, err := data.GetAllCustomers()
				if err != nil {
					danger("When select all customers get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				args:=make(map[string]string)
				args["cstmName"] = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				args["city"] = strings.TrimFunc(request.PostFormValue("city"), unicode.IsSpace)
				args["area"] = strings.TrimFunc(request.PostFormValue("area"), unicode.IsSpace)
				args["police"] = strings.TrimFunc(request.PostFormValue("police"), unicode.IsSpace)
				rsts, err := data.SelectCustomers(args)
				if err != nil {
					danger("When select some customers get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "insert":
			customer := data.Customer{}
			customer.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			customer.CstmType = strings.TrimFunc(request.PostFormValue("cstmtype"), unicode.IsSpace)
			customer.City = strings.TrimFunc(request.PostFormValue("city"), unicode.IsSpace)
			customer.Area = strings.TrimFunc(request.PostFormValue("area"), unicode.IsSpace)
			customer.Address = strings.TrimFunc(request.PostFormValue("address"), unicode.IsSpace)
			customer.OwnerName = strings.TrimFunc(request.PostFormValue("owner_name"), unicode.IsSpace)
			customer.Telephone = strings.TrimFunc(request.PostFormValue("telephone"), unicode.IsSpace)
			customer.Police = strings.TrimFunc(request.PostFormValue("police"), unicode.IsSpace)
			customer.Axis = strings.TrimFunc(request.PostFormValue("axis"), unicode.IsSpace)
			customer.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.InsertCustomer(&customer)
			if err != nil {
				danger("When insert the customer get error:", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "update":
		    customer := data.Customer{}
			customer.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			customer.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			customer.CstmType = strings.TrimFunc(request.PostFormValue("cstmtype"), unicode.IsSpace)
			customer.City = strings.TrimFunc(request.PostFormValue("city"), unicode.IsSpace)
			customer.Area = strings.TrimFunc(request.PostFormValue("area"), unicode.IsSpace)
			customer.Address = strings.TrimFunc(request.PostFormValue("address"), unicode.IsSpace)
			customer.OwnerName = strings.TrimFunc(request.PostFormValue("owner_name"), unicode.IsSpace)
			customer.Telephone = strings.TrimFunc(request.PostFormValue("telephone"), unicode.IsSpace)
			customer.Police = strings.TrimFunc(request.PostFormValue("police"), unicode.IsSpace)
			customer.Axis = strings.TrimFunc(request.PostFormValue("axis"), unicode.IsSpace)
			customer.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.UpdateCustomer(&customer)
			if err != nil {
				danger("When update the customer get error:", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "delete":
			id, _ := strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := data.DeleteCustomer(id)
			if err != nil {
				danger("When delete the customer get error", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		}
	}
}


func handleContracts(writer http.ResponseWriter, request *http.Request){
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		operation := strings.TrimFunc(request.PostFormValue("operation"), unicode.IsSpace)
		tt := strings.TrimFunc(request.PostFormValue("tt"), unicode.IsSpace)
		switch operation {
		case "select":
			writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
			if tt == "all" {
				rsts, err := data.GetAllContracts()
				if err != nil {
					danger("When select all products get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				args:=make(map[string]string)
				args["ccsn"] = strings.TrimFunc(request.PostFormValue("ccsn"), unicode.IsSpace)
				args["vector"] = strings.TrimFunc(request.PostFormValue("vector"), unicode.IsSpace)
				args["createDate"] = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
				args["cstmName"] = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				args["specific"] = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
				args["prdtName"] = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			    fmt.Println(args);
				rsts, err := data.SelectContracts(args)
				if err != nil {
					danger("When select the contracts get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "insert":
			contract := data.Contract{}
			contract.Ccsn = strings.TrimFunc(request.PostFormValue("ccsn"), unicode.IsSpace)
			contract.CreateDate = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
			contract.CcType = strings.TrimFunc(request.PostFormValue("cctype"), unicode.IsSpace)
			contract.Price, _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("price"), unicode.IsSpace), 64)
			contract.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			contract.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			contract.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			contract.Vector = strings.TrimFunc(request.PostFormValue("vector"), unicode.IsSpace)
			contract.IvType = strings.TrimFunc(request.PostFormValue("ivtype"), unicode.IsSpace)
			contract.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			contract.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.InsertContract(&contract)
			if err != nil {
				danger("When insert the contract get error:", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "update":
			contract := data.Contract{}
			contract.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			contract.Ccsn = strings.TrimFunc(request.PostFormValue("ccsn"), unicode.IsSpace)
			contract.CreateDate = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
			contract.CcType = strings.TrimFunc(request.PostFormValue("cctype"), unicode.IsSpace)
			contract.Price, _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("price"), unicode.IsSpace), 64)
			contract.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			contract.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			contract.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			contract.Vector = strings.TrimFunc(request.PostFormValue("vector"), unicode.IsSpace)
			contract.IvType = strings.TrimFunc(request.PostFormValue("ivtype"), unicode.IsSpace)
			contract.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			contract.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			fmt.Println(contract)
			err := data.UpdateContract(&contract)
			if err != nil {
				danger("When update the contract get error:", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "delete":
			id, _ := strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := data.DeleteContract(id)
			if err != nil {
				danger("When delete the contract get error:", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		}
	}
}


func handleInStocks(writer http.ResponseWriter, request *http.Request){
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		operation := strings.TrimFunc(request.PostFormValue("operation"), unicode.IsSpace)
		tt := strings.TrimFunc(request.PostFormValue("tt"), unicode.IsSpace)
		switch operation {
		case "select":
			writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
			if tt == "all" {
				rsts, err := data.GetAllInStocks()
				if err != nil {
					danger("When select all instocks get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				args:=make(map[string]string)
				args["cstmName"] = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				args["createDate"] = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
				args["prdtName"] = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
				args["specific"] = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
				args["mac"] = strings.TrimFunc(request.PostFormValue("mac"), unicode.IsSpace)
				args["prdtName"] = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			
				rsts, err := data.SelectInStocks(args)
				if err != nil {
					danger("When select the instocks get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "insert":
			inStock := data.InStock{}
			inStock.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			inStock.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			inStock.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			inStock.Mac = strings.TrimFunc(request.PostFormValue("mac"), unicode.IsSpace)
			inStock.Sn = strings.TrimFunc(request.PostFormValue("sn"), unicode.IsSpace)
			inStock.CreateDate = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
			inStock.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			inStock.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.InsertInStock(&inStock)
			if err != nil {
				danger("When insert the instock get error:", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "update":
			inStock := data.InStock{}
			inStock.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			inStock.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			inStock.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			inStock.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			inStock.Mac = strings.TrimFunc(request.PostFormValue("mac"), unicode.IsSpace)
			inStock.Sn = strings.TrimFunc(request.PostFormValue("sn"), unicode.IsSpace)
			inStock.CreateDate = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
			inStock.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			inStock.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.UpdateInStock(&inStock)
			if err != nil {
				danger("When update the instock get error:", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "delete":
			id, _ := strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := data.DeleteInStock(id)
			if err != nil {
				danger("When delete the instock get error:", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		}
	}
}



func handleOutStocks(writer http.ResponseWriter, request *http.Request){
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		operation := strings.TrimFunc(request.PostFormValue("operation"), unicode.IsSpace)
		tt := strings.TrimFunc(request.PostFormValue("tt"), unicode.IsSpace)
		switch operation {
		case "select":
			writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
			if tt == "all" {
				rsts, err := data.GetAllOutStocks()
				if err != nil {
					danger("When select all outstocks get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				args:=make(map[string]string)
				args["cstmName"] = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				args["createDate"] = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
				args["prdtName"] = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
				args["specific"] = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
				args["mac"] = strings.TrimFunc(request.PostFormValue("mac"), unicode.IsSpace)
				args["prdtName"] = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			
				rsts, err := data.SelectOutStocks(args)
				if err != nil {
					danger("When select the outstocks get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "insert":
			outStock := data.OutStock{}
			outStock.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			outStock.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			outStock.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			outStock.Mac = strings.TrimFunc(request.PostFormValue("mac"), unicode.IsSpace)
			outStock.Sn = strings.TrimFunc(request.PostFormValue("sn"), unicode.IsSpace)
			outStock.CreateDate = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
			outStock.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			outStock.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.InsertOutStock(&outStock)
			if err != nil {
				danger("When insert the outStock get error:", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "update":
			outStock := data.OutStock{}
			outStock.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			outStock.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			outStock.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			outStock.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			outStock.Mac = strings.TrimFunc(request.PostFormValue("mac"), unicode.IsSpace)
			outStock.Sn = strings.TrimFunc(request.PostFormValue("sn"), unicode.IsSpace)
			outStock.CreateDate = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
			outStock.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			outStock.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.UpdateOutStock(&outStock)
			if err != nil {
				danger("When update the outstock get error:", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "delete":
			id, _ := strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := data.DeleteOutStock(id)
			if err != nil {
				danger("When delete the outstock get error:", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		}
	}
}

func handleInvoices(writer http.ResponseWriter, request *http.Request){
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		operation := strings.TrimFunc(request.PostFormValue("operation"), unicode.IsSpace)
		tt := strings.TrimFunc(request.PostFormValue("tt"), unicode.IsSpace)
		switch operation {
		case "select":
			writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
			if tt == "all" {
				rsts, err := data.GetAllInvoices()
				if err != nil {
					danger("When select all outstocks get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				args:=make(map[string]string)
				args["sn"] = strings.TrimFunc(request.PostFormValue("sn"), unicode.IsSpace)
				args["createDate"] = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
				args["cstmName"] = strings.TrimFunc(request.PostFormValue("cstmName"), unicode.IsSpace)
				args["specific"] = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
				args["prdtName"] = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			
				rsts, err := data.SelectInvoices(args)
				if err != nil {
					danger("When select the invoices get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "insert":
			invoice := data.Invoice{}
			invoice.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			invoice.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			invoice.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			invoice.Sn = strings.TrimFunc(request.PostFormValue("sn"), unicode.IsSpace)
			invoice.IvType = strings.TrimFunc(request.PostFormValue("ivtype"), unicode.IsSpace)
			invoice.Price, _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("price"), unicode.IsSpace), 64)
			invoice.CreateDate = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
			invoice.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			invoice.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.InsertInvoice(&invoice)
			if err != nil {
				danger("When insert the invoice get error:", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "update":
			invoice := data.Invoice{}
			invoice.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			invoice.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			invoice.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			invoice.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			invoice.Sn = strings.TrimFunc(request.PostFormValue("sn"), unicode.IsSpace)
			invoice.IvType = strings.TrimFunc(request.PostFormValue("ivtype"), unicode.IsSpace)
			invoice.Price, _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("price"), unicode.IsSpace), 64)
			invoice.CreateDate = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
			invoice.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			invoice.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.UpdateInvoice(&invoice)
			if err != nil {
				danger("When update the invoice get error:", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "delete":
			id, _ := strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := data.DeleteInvoice(id)
			if err != nil {
				danger("When delete the invoice get error:", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		}
	}
}


func handlePayments(writer http.ResponseWriter, request *http.Request){
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		operation := strings.TrimFunc(request.PostFormValue("operation"), unicode.IsSpace)
		tt := strings.TrimFunc(request.PostFormValue("tt"), unicode.IsSpace)
		switch operation {
		case "select":
			writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
			if tt == "all" {
				rsts, err := data.GetAllPayments()
				if err != nil {
					danger("When select all payments get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				args:=make(map[string]string)
				args["sn"] = strings.TrimFunc(request.PostFormValue("sn"), unicode.IsSpace)
				args["createDate"] = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
				args["cstmName"] = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			
				rsts, err := data.SelectPayments(args)
				if err != nil {
					danger("When select the payments get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "insert":
			payment := data.Payment{}
			payment.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			payment.CreateDate = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
			payment.PmType = strings.TrimFunc(request.PostFormValue("pmtype"), unicode.IsSpace)
			payment.Amount,  _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("amount"), unicode.IsSpace), 64)
			payment.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.InsertPayment(&payment)
			if err != nil {
				danger("When insert the payment get error:", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "update":
			payment := data.Payment{}
			payment.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			payment.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			payment.CreateDate = strings.TrimFunc(request.PostFormValue("create_date"), unicode.IsSpace)
			payment.PmType = strings.TrimFunc(request.PostFormValue("pmtype"), unicode.IsSpace)
			payment.Amount,  _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("amount"), unicode.IsSpace), 64)
			payment.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := data.UpdatePayment(&payment)
			if err != nil {
				danger("When update the payment get error:", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "delete":
			id, _ := strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := data.DeletePayment(id)
			if err != nil {
				danger("When delete the payment get error:", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		}
	}
}


func handleStocks(writer http.ResponseWriter, request *http.Request){
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		// operation := strings.TrimFunc(request.PostFormValue("operation"), unicode.IsSpace)
		tt := strings.TrimFunc(request.PostFormValue("tt"), unicode.IsSpace)
		writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
		
		if tt == "all" {
				rsts, err := data.GetAllStocks()
				if err != nil {
					danger("When select all stocks get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
		} else {
				args:=make(map[string]string)
				args["prdtName"] = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
				args["specific"] = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
				args["mac"] = strings.TrimFunc(request.PostFormValue("mac"), unicode.IsSpace)
			
				rsts, err := data.SelectStocks(args)
				if err != nil {
					danger("When select the stocks get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
	}
}


func handleDebts(writer http.ResponseWriter, request *http.Request){
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		// operation := strings.TrimFunc(request.PostFormValue("operation"), unicode.IsSpace)
		tt := strings.TrimFunc(request.PostFormValue("tt"), unicode.IsSpace)
		writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
		
		if tt == "all" {
				rsts, err := data.GetAllDebts()
				if err != nil {
					danger("When select all debts get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
		} else {
				args:=make(map[string]string)
				args["cstmName"] = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			
				rsts, err := data.SelectDebts(args)
				if err != nil {
					danger("When select the debts get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
		}
	}
}





func handleOnWayProducts(writer http.ResponseWriter, request *http.Request){
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		// operation := strings.TrimFunc(request.PostFormValue("operation"), unicode.IsSpace)
		tt := strings.TrimFunc(request.PostFormValue("tt"), unicode.IsSpace)
		writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
		
		if tt == "all" {
				rsts, err := data.GetAllOnWayProducts()
				if err != nil {
					danger("When select all onwayproducts get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
		} else {
				args:=make(map[string]string)
				args["cstmName"] = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				args["prdtName"] = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
				args["specific"] = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			
				rsts, err := data.SelectOnWayProducts(args)
				if err != nil {
					danger("When select the onwayproducts get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
		}
	}
}



func handleOnWayInvoices(writer http.ResponseWriter, request *http.Request){
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		// operation := strings.TrimFunc(request.PostFormValue("operation"), unicode.IsSpace)
		tt := strings.TrimFunc(request.PostFormValue("tt"), unicode.IsSpace)
		writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
		
		if tt == "all" {
				rsts, err := data.GetAllOnWayInvoices()
				if err != nil {
					danger("When select all onwayproducts get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
		} else {
				args:=make(map[string]string)
				args["cstmName"] = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				args["prdtName"] = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
				args["specific"] = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			
				rsts, err := data.SelectOnWayInvoices(args)
				if err != nil {
					danger("When select the onwayinvoices get error:", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
		}
	}
}



func getCustomerName(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		rsts, err := data.GetCustomerName()
		if err != nil {
			danger(err, "Cannot get all customername")
		}
		resp, _ := json.Marshal(rsts)
		fmt.Fprintf(writer, string(resp))
	}
}

func getProductNS(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		rsts, err := data.GetProductNS()
		if err != nil {
			danger(err, "Cannot get all prdtns")
		}
		resp, _ := json.Marshal(rsts)
		fmt.Fprintf(writer, string(resp))
	}
}
