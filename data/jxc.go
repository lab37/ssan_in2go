package data

import "log"


type Product struct {
	Id       int
	PrdtName string
	Specific string
	Inventor string
	Unit     string
	IvType   string
	Remark   string
}

type Customer struct {
	Id       int
	CstmName  string
	CstmType  string
	City      string
	Area      string
	Address   string
	OwnerName string
	Telephone string
	Police    string
	Axis      string
	Remark    string
}

type Contract struct {
	Id         int
	Ccsn       string
	Vector     string
	CreateDate string
	CcType     string
	CstmName   string
	PrdtName   string
	Specific   string
	Price      float64
	Quantity   int
	IvType     string
	Remark     string
}
type InStock struct {
	Id         int
	CstmName   string
	PrdtName   string
	Specific   string
	Mac        string
	Sn         string
	CreateDate string
	Quantity   int
	Remark     string
}
type OutStock struct {
	Id         int
	CstmName   string
	PrdtName   string
	Specific   string
	Mac        string
	Sn         string
	CreateDate string
	Quantity   int
	Remark     string
}
type Invoice struct {
	Id          int
	Sn          string
	IvType      string
	CreateDate  string
	CstmName    string
	PrdtName    string
	Specific    string
	Quantity    int
	Price       float64
	Remark      string
}

type Payment struct {
	Id         int
	CstmName   string
	CreateDate string
	PmType     string
	Amount     float64
	Remark     string
}

type Stock struct {
	Id       int
	PrdtName string
	Specific string
	Mac      string
	Sn       string
	Quantity int
	Remark   string
}

type Debt struct {
	Id       int
	CstmName string
	Amount   float64
	Remark   string
}

type OnWayProduct struct {
	Id       int
	CstmName string
	PrdtName string
	Specific string
	Quantity int
	Remark   string
}

type OnWayInvoice struct {
	Id       int
	CstmName string
	PrdtName string
	Specific string
	Amount   float64
	Remark   string
}

func InsertProduct(r *Product) (err error) {
	statement := "insert into products (prdtname, specific, inventor, unit, ivtype, remark) values (?,?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.PrdtName, r.Specific, r.Inventor, r.Unit, r.IvType, r.Remark)
	return
}

func SelectProducts(args map[string]string) (products []Product, err error) {
	rows, err := Db.Query("SELECT * FROM products WHERE prdtname LIKE ? AND specific LIKE ? AND inventor LIKE ?", "%"+args["prdtName"]+"%", "%"+args["specific"]+"%", "%"+args["inventor"]+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	product := Product{}
	for rows.Next() {
		
		if err = rows.Scan(&product.Id, &product.PrdtName, &product.Specific, &product.Inventor, &product.Unit, &product.IvType, &product.Remark); err != nil {
			return
		}
		products = append(products, product)
	}

	return
}

func GetAllProducts() (products []Product, err error) {
	rows, err := Db.Query("SELECT * FROM products")
	defer rows.Close()
	if err != nil {
		return
	}
	product := Product{}
	for rows.Next() {
		
		if err = rows.Scan(&product.Id, &product.PrdtName, &product.Specific, &product.Inventor, &product.Unit, &product.IvType, &product.Remark); err != nil {

			return
		}
		products = append(products, product)
	}

	return
}

func UpdateProduct(r *Product) (err error) {
	_, err = Db.Exec("UPDATE products set prdtname=?,specific=?,inventor=?,unit=? ,ivtype=?,remark=? where id=?", r.PrdtName, r.Specific, r.Inventor, r.Unit, r.IvType, r.Remark, r.Id)
	return
}

func DeleteProduct(id int) (err error) {
	_, err = Db.Exec("DELETE FROM products WHERE id=?", id)
	return
}



func GetAllCustomers() (customers []Customer, err error) {

	rows, err := Db.Query("SELECT * FROM customers")
	defer rows.Close()
	if err != nil {
		return
	}
	customer := Customer{}
	for rows.Next() {
		
		if err = rows.Scan(&customer.Id, &customer.CstmName, &customer.CstmType, &customer.City, &customer.Area, &customer.Address, &customer.OwnerName, &customer.Telephone, &customer.Police, &customer.Axis, &customer.Remark); err != nil {
			return
		}
		customers = append(customers, customer)
	}

	return
}


func SelectCustomers(args map[string]string) (customers []Customer, err error) {
	rows, err := Db.Query("SELECT * FROM customers WHERE cstmname LIKE ? AND city LIKE ? AND area LIKE ? AND police LIKE ?", "%"+args["cstmName"]+"%", "%"+args["city"]+"%", "%"+args["area"]+"%", "%"+args["police"]+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	customer := Customer{}
	for rows.Next() {
		
		if err = rows.Scan(&customer.Id, &customer.CstmName, &customer.CstmType, &customer.City, &customer.Area, &customer.Address, &customer.OwnerName, &customer.Telephone, &customer.Police, &customer.Axis, &customer.Remark); err != nil {
			return
		}
		customers = append(customers, customer)
	}

	return
}

func InsertCustomer(r *Customer) (err error) {
	statement := "insert into customers (cstmname, cstmtype, city, area, cstmaddr, owner_name, telephone, police, axis, remark) values (?,?,?,?,?,?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.CstmName, r.CstmType, r.City, r.Area, r.Address, r.OwnerName, r.Telephone, r.Police, r.Axis, r.Remark)
	return
}

func UpdateCustomer(r *Customer) (err error) {
	_, err = Db.Exec("UPDATE customers set cstmname=?,cstmtype=?,city=?, area=?, cstmaddr=?, owner_name=?, telephone=?, police=?, axis=?, remark=? where id=?", r.CstmName, r.CstmType, r.City, r.Area, r.Address, r.OwnerName, r.Telephone, r.Police, r.Axis, r.Remark, r.Id)
	return
}

func DeleteCustomer(id int) (err error) {
	_, err = Db.Exec("DELETE FROM customers WHERE id=?", id)
	return
}







func GetAllContracts() (contracts []Contract, err error) {

	rows, err := Db.Query("SELECT id, ccsn , vector, create_date, cctype, ivtype, (SELECT cstmname FROM customers WHERE customers.id = cstmid limit 1) AS cstmname,(SELECT prdtname FROM products WHERE products.id = prdtid limit 1) AS prdtname, (SELECT specific FROM products WHERE products.id = prdtid limit 1) AS specific,price, quantity, remark FROM contracts  order by create_date desc")
	defer rows.Close()
	if err != nil {
		return
	}
	contract := Contract{}
	for rows.Next() {
		
		if err = rows.Scan(&contract.Id, &contract.Ccsn, &contract.Vector, &contract.CreateDate, &contract.CcType, &contract.IvType, &contract.CstmName, &contract.PrdtName,&contract.Specific, &contract.Price, &contract.Quantity, &contract.Remark); err != nil {

			return
		}
		contracts = append(contracts, contract)
	}

	return
}

func SelectContracts(args map[string]string) (contracts []Contract, err error) {
	rows, err := Db.Query("SELECT id, ccsn , vector, create_date, cctype, ivtype, (SELECT cstmname FROM customers WHERE customers.id = cstmid limit 1) AS cstmname,(SELECT prdtname FROM products WHERE products.id = prdtid limit 1) AS prdtname, (SELECT specific FROM products WHERE products.id = prdtid limit 1) AS specific,price, quantity, remark FROM contracts WHERE ccsn LIKE ? AND vector = ? AND create_date LIKE ? AND cstmid in (select id from customers where cstmname LIKE ?) AND prdtid in (select id from products where prdtname LIKE ?) order by create_date desc", "%"+args["ccsn"]+"%",args["vector"], "%"+args["createDate"]+"%", "%"+args["cstmName"]+"%", "%"+args["prdtName"]+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	contract := Contract{}
	for rows.Next() {
		
		if err = rows.Scan(&contract.Id, &contract.Ccsn, &contract.Vector, &contract.CreateDate, &contract.CcType, &contract.IvType, &contract.CstmName, &contract.PrdtName,&contract.Specific, &contract.Price, &contract.Quantity, &contract.Remark); err != nil {

			return
		}
		contracts = append(contracts, contract)
	}
	log.Println(contracts);

	return




}

func  InsertContract(r *Contract) (err error) {
	statement := "insert into contracts(ccsn,vector, create_date, cctype, ivtype, cstmid, prdtid, price, quantity, remark) values (?,?,?,?,?,(select id from customers where customers.cstmname = ? limit 1),(select id from products where products.prdtname = ? and products.specific = ? limit 1),?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {

		return
	}

	_, err = stmt.Exec(r.Ccsn, r.Vector, r.CreateDate, r.CcType, r.IvType, r.CstmName, r.PrdtName, r.Specific, r.Price, r.Quantity, r.Remark)
	return
}


func UpdateContract(r *Contract) (err error) {
	_, err = Db.Exec("UPDATE contracts set ccsn=?, vector=?, create_date=?,cctype=?, ivtype=?,cstmid=(select id from customers where customers.cstmname = ? limit 1),prdtid=(select id from products where products.prdtname = ? and products.specific = ? limit 1),price=?,quantity=?,remark=? where id=?", r.Ccsn, r.Vector, r.CreateDate, r.CcType, r.IvType, r.CstmName, r.PrdtName, r.Specific, r.Price, r.Quantity, r.Remark, r.Id)
	log.Println(err)
	return
}

func  DeleteContract(id int) (err error) {
	_, err = Db.Exec("DELETE FROM contracts WHERE id=?", id)
	return
}


func GetAllInStocks() (inStocks []InStock, err error) {
	rows, err := Db.Query("SELECT id, (SELECT cstmname FROM customers WHERE customers.id=instocks.cstmid) AS cstmname, (SELECT prdtname FROM products WHERE products.id = instocks.prdtid) AS prdtname, (SELECT specific FROM products WHERE products.id = instocks.prdtid) AS specific, mac, sn, create_date, quantity, remark FROM instocks  order by create_date desc")
	defer rows.Close()
	if err != nil {
		return
	}
	inStock := InStock{}
	for rows.Next() {
		
		if err = rows.Scan(&inStock.Id, &inStock.CstmName, &inStock.PrdtName, &inStock.Specific, &inStock.Mac, &inStock.Sn, &inStock.CreateDate, &inStock.Quantity, &inStock.Remark); err != nil {
			return
		}
		inStocks = append(inStocks, inStock)
	}

	return
}


func SelectInStocks(args map[string]string) (inStocks []InStock, err error) {
	rows, err := Db.Query("SELECT id, (SELECT cstmname FROM customers WHERE customers.id=instocks.cstmid) AS cstmname, (SELECT prdtname FROM products WHERE products.id = instocks.prdtid) AS prdtname, (SELECT specific FROM products WHERE products.id = instocks.prdtid) AS specific, mac, sn, create_date, quantity, remark FROM instocks WHERE cstmid in (select id from customers where customers.cstmname like ?) AND create_date like ? AND prdtid in (select id from products where products.prdtname like ? AND specific like ?) AND mac like ?order by create_date desc", "%"+args["cstmName"]+"%", "%"+args["createDate"]+"%", "%"+args["prdtName"]+"%", "%"+args["specific"]+"%", "%"+args["mac"]+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	inStock := InStock{}
	for rows.Next() {
		
		if err = rows.Scan(&inStock.Id, &inStock.CstmName, &inStock.PrdtName, &inStock.Specific, &inStock.Mac, &inStock.Sn, &inStock.CreateDate, &inStock.Quantity, &inStock.Remark); err != nil {
			return
		}
		inStocks = append(inStocks, inStock)
	}

	return
}




func InsertInStock(r *InStock) (err error) {
	statement := "insert into instocks(cstmid,prdtid, mac, sn, create_date, quantity, remark) values ((select id from customers where customers.cstmname like ? limit 1),(select id from products where products.prdtname like ? and products.specific like ? limit 1),?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec("%"+r.CstmName+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", r.Mac, r.Sn, r.CreateDate, r.Quantity,  r.Remark)
	return
}



func UpdateInStock(r *InStock) (err error) {
	_, err = Db.Exec("UPDATE instocks set cstmid=(select id from customers where customers.cstmname like ? limit 1),prdtid=(select id from products where products.prdtname like ? and products.specific like ? limit 1), mac=?, sn=?, create_date=?,quantity=?,remark=? where id=?", "%"+r.CstmName+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", r.Mac, r.Sn, r.CreateDate, r.Quantity,  r.Remark, r.Id)
	return
}

func DeleteInStock(id int) (err error) {
	_, err = Db.Exec("DELETE FROM instocks WHERE id=?", id)
	return
}




func GetAllOutStocks() (outStocks []OutStock, err error) {
	rows, err := Db.Query("SELECT id, (SELECT cstmname FROM customers WHERE customers.id=outstocks.cstmid) AS cstmname, (SELECT prdtname FROM products WHERE products.id = outstocks.prdtid) AS prdtname, (SELECT specific FROM products WHERE products.id = outstocks.prdtid) AS specific, mac, sn, create_date, quantity, remark FROM outstocks  order by create_date desc")
	defer rows.Close()
	if err != nil {
		return
	}
	outStock := OutStock{}
	for rows.Next() {
		
		if err = rows.Scan(&outStock.Id, &outStock.CstmName, &outStock.PrdtName, &outStock.Specific, &outStock.Mac, &outStock.Sn, &outStock.CreateDate, &outStock.Quantity, &outStock.Remark); err != nil {
			return
		}
		outStocks = append(outStocks, outStock)
	}

	return
}


func SelectOutStocks(args map[string]string) (outStocks []OutStock, err error) {
	rows, err := Db.Query("SELECT id, (SELECT cstmname FROM customers WHERE customers.id=outStocks.cstmid) AS cstmname, (SELECT prdtname FROM products WHERE products.id = outStocks.prdtid) AS prdtname, (SELECT specific FROM products WHERE products.id = outStocks.prdtid) AS specific, mac, sn, create_date, quantity, remark FROM outStocks WHERE cstmid in (select id from customers where customers.cstmname like ?) AND create_date like ? AND prdtid in (select id from products where products.prdtname like ? AND specific like ?) AND mac like ?order by create_date desc", "%"+args["cstmName"]+"%", "%"+args["createDate"]+"%", "%"+args["prdtName"]+"%", "%"+args["specific"]+"%", "%"+args["mac"]+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	outStock := OutStock{}
	for rows.Next() {
		
		if err = rows.Scan(&outStock.Id, &outStock.CstmName, &outStock.PrdtName, &outStock.Specific, &outStock.Mac, &outStock.Sn, &outStock.CreateDate, &outStock.Quantity, &outStock.Remark); err != nil {
			return
		}
		outStocks = append(outStocks, outStock)
	}

	return
}




func InsertOutStock(r *OutStock) (err error) {
	statement := "insert into outStocks(cstmid,prdtid, mac, sn, create_date, quantity, remark) values ((select id from customers where customers.cstmname like ? limit 1),(select id from products where products.prdtname like ? and products.specific like ? limit 1),?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec("%"+r.CstmName+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", r.Mac, r.Sn, r.CreateDate, r.Quantity,  r.Remark)
	return
}



func UpdateOutStock(r *OutStock) (err error) {
	_, err = Db.Exec("UPDATE outstocks set cstmid=(select id from customers where customers.cstmname like ? limit 1),prdtid=(select id from products where products.prdtname like ? and products.specific like ? limit 1), mac=?, sn=?, create_date=?,quantity=?,remark=? where id=?", "%"+r.CstmName+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", r.Mac, r.Sn, r.CreateDate, r.Quantity,  r.Remark, r.Id)
	return
}

func DeleteOutStock(id int) (err error) {
	_, err = Db.Exec("DELETE FROM outstocks WHERE id=?", id)
	return
}


func GetAllInvoices() (invoices []Invoice, err error) {
	rows, err := Db.Query("SELECT id, sn, ivtype, create_date, (select cstmname from customers where customers.id = cstmid limit 1) as cstmname, (SELECT prdtname FROM products where products.id = prdtid limit 1)  AS prdtname, (SELECT specific FROM products where products.id = prdtid limit 1)  AS specific, quantity, price, remark FROM invoices  order by create_date desc")
	defer rows.Close()
	if err != nil {
		return
	}
	invoice := Invoice{}
	for rows.Next() {
		
		if err = rows.Scan(&invoice.Id, &invoice.Sn, &invoice.IvType, &invoice.CreateDate, &invoice.CstmName, &invoice.PrdtName, &invoice.Specific, &invoice.Quantity, &invoice.Price, &invoice.Remark); err != nil {
			return
		}
		invoices = append(invoices, invoice)
	}

	return
}

func SelectInvoices(args map[string]string) (invoices []Invoice, err error) {
	rows, err := Db.Query("SELECT id, sn, ivtype, create_date, (select cstmname from customers where customers.id = cstmid limit 1) as cstmname, (SELECT prdtname FROM products where products.id = prdtid limit 1)  AS prdtname, (SELECT specific FROM products where products.id = prdtid limit 1)  AS specific, quantity, price, remark FROM invoices WHERE sn LIKE ? AND create_date LIKE ? AND cstmid in (select id from customers where customers.cstmname LIKE ?) AND prdtid in (select id from products where products.prdtname LIKE ? AND products.specific LIKE ?) ", "%"+args["sn"]+"%", "%"+args["createDate"]+"%", "%"+args["cstmname"]+"%", "%"+args["prdtName"]+"%", "%"+args["specific"]+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	invoice := Invoice{}
	for rows.Next() {
		
		if err = rows.Scan(&invoice.Id, &invoice.Sn, &invoice.IvType, &invoice.CreateDate, &invoice.CstmName, &invoice.PrdtName, &invoice.Specific, &invoice.Quantity, &invoice.Price, &invoice.Remark); err != nil {
			return
		}
		invoices = append(invoices, invoice)
	}

	return
}

func InsertInvoice(r *Invoice) (err error) {
	statement := "insert into invoices(sn, ivtype, create_date, cstmid, prdtid, quantity, price, remark) values(?,?,?,(select id from customers where customers.cstmname LIKE ? limit 1),(select id from products where products.prdtname LIKE ? AND products.specific LIKE ? limit 1),?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.Sn, r.IvType,r.CreateDate, "%"+r.CstmName+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", r.Quantity, r.Price, r.Remark)
	return
}


func UpdateInvoice(r *Invoice) (err error) {
	_, err = Db.Exec("UPDATE invoices set sn=?,ivtype=?, create_date=?,cstmid=(select id from customers where customers.cstmname LIKE ? limit 1),prdtid=(select id from products where products.prdtname LIKE ? AND products.specific LIKE ? limit 1),quantity=?,price=?,remark=? where id=?", r.Sn,r.IvType, r.CreateDate, "%"+r.CstmName+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", r.Quantity, r.Price, r.Remark, r.Id)
	return
}

func DeleteInvoice(id int) (err error) {
	_, err = Db.Exec("DELETE FROM invoices WHERE id=?", id)
	return
}




func InsertPayment(r *Payment) (err error) {
	statement := "insert into payments (cstmid, create_date, amount, remark) values ((select id from customers where customers.cstmname = ?),?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.CstmName, r.CreateDate, r.Amount, r.Remark)
	return
}

func GetAllPayments() (payments []Payment, err error) {
	rows, err := Db.Query("SELECT id ,(SELECT cstmname FROM customers where customers.id = cstmid limit 1) AS cstmname, create_date, pmtype, amount, remark FROM payments order by create_date desc")
	defer rows.Close()
	if err != nil {
		return
	}
	payment := Payment{}
	for rows.Next() {
		
		if err = rows.Scan(&payment.Id,  &payment.CstmName, &payment.CreateDate,&payment.PmType, &payment.Amount, &payment.Remark); err != nil {
			return
		}
		payments = append(payments, payment)
	}

	return
}


func SelectPayments(args map[string]string) (payments []Payment, err error) {
	rows, err := Db.Query("SELECT id ,(SELECT cstmname FROM customers where customers.id = cstmid limit 1) AS cstmname, create_date, pmtype, amount, remark FROM payments WHERE cstmid in (select id from customers where customers.cstmname LIKE ?) AND create_date LIKE ? order by create_date desc", "%"+args["cstmName"]+"%",  "%"+args["createDate"]+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	payment := Payment{}
	for rows.Next() {
		
		if err = rows.Scan(&payment.Id,  &payment.CstmName, &payment.CreateDate,&payment.PmType, &payment.Amount, &payment.Remark); err != nil {
			return
		}
		payments = append(payments, payment)
	}

	return
}


func UpdatePayment(r *Payment) (err error) {
	_, err = Db.Exec("UPDATE payments set cstmid=(select id from customers where customers.cstmname = ? limit 1) ,create_date=?, pmtype=?, amount=?,remark=? where id=?", r.CstmName, r.CreateDate, r.PmType, r.Amount, r.Remark, r.Id)
	return
}

func DeletePayment(id int) (err error) {
	_, err = Db.Exec("DELETE FROM payments WHERE id=?", id)
	return
}












func GetAllStocks() (stocks []Stock, err error) {
	rows, err := Db.Query("select id, (select prdtname from products where products.id=stocks.prdtid limit 1) as prdtname, (select specific from products where products.id=stocks.prdtid limit 1) as specific, mac, sn, quantity, remark from stocks")
	defer rows.Close()
	if err != nil {
		return
	}
	stock := Stock{}
	for rows.Next() {
		
		if err = rows.Scan(&stock.Id, &stock.PrdtName, &stock.Specific, &stock.Mac, &stock.Sn, &stock.Quantity, &stock.Remark); err != nil {
			return
		}
		stocks = append(stocks, stock)
	}

	return
}


func SelectStocks(args map[string]string) (stocks []Stock, err error) {
	rows, err := Db.Query("select id, (select prdtname from products where products.id=stocks.prdtid limit 1) as prdtname, (select specific from products where products.id=stocks.prdtid limit 1) as specific, mac, sn, quantity, remark from stocks WHERE prdtid in (select id from products where products.prdtname LIKE ? AND products.specific LIKE ?) AND mac LIKE ? ", "%"+args["prdtName"]+"%", "%"+args["specific"]+"%", "%"+args["mac"]+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	stock := Stock{}
	for rows.Next() {
		
		if err = rows.Scan(&stock.Id, &stock.PrdtName, &stock.Specific, &stock.Mac, &stock.Sn, &stock.Quantity, &stock.Remark); err != nil {
			return
		}
		stocks = append(stocks, stock)
	}
	return
}


func GetAllDebts() (debts []Debt, err error) {
	rows, err := Db.Query("select id, (select cstmname from customers where customers.id=debts.cstmid limit 1) as cstmname, amount, remark from debts")
	defer rows.Close()
	if err != nil {
		return
	}
	debt := Debt{}
	for rows.Next() {
		
		if err = rows.Scan(&debt.Id, &debt.CstmName, &debt.Amount, &debt.Remark); err != nil {
			return
		}
		debts = append(debts, debt)
	}

	return
}


func SelectDebts(args map[string]string) (debts []Debt, err error) {
	rows, err := Db.Query("select id, (select cstmname from customers where customers.id=debts.cstmid limit 1) as cstmname, amount, remark from debts WHERE cstmid in (select id from customers where customers.cstmname LIKE ?)", "%"+args["cstmName"]+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	debt := Debt{}
	for rows.Next() {
		
		if err = rows.Scan(&debt.Id, &debt.CstmName, &debt.Amount, &debt.Remark); err != nil {
			return
		}
		debts = append(debts, debt)
	}

	return
}





func GetAllOnWayProducts() (onWayProducts []OnWayProduct, err error) {
	rows, err := Db.Query("select id, (select cstmname from customers where customers.id=onway_products.cstmid limit 1) as cstmname,(select prdtname from products where products.id=onway_products.prdtid limit 1) as prdtname,(select specific from products where products.id=onway_products.prdtid limit 1) as specific, quantity, remark from onway_products")
	defer rows.Close()
	if err != nil {
		return
	}
	onWayProduct := OnWayProduct{}
	for rows.Next() {
		
		if err = rows.Scan(&onWayProduct.Id, &onWayProduct.CstmName, &onWayProduct.PrdtName,&onWayProduct.Specific,&onWayProduct.Quantity, &onWayProduct.Remark); err != nil {
			return
		}
		onWayProducts = append(onWayProducts, onWayProduct)
	}

	return
}


func SelectOnWayProducts(args map[string]string) (onWayProducts []OnWayProduct, err error) {
	rows, err := Db.Query("select id, (select cstmname from customers where customers.id=onway_products.cstmid limit 1) as cstmname,(select prdtname from products where products.id=onway_products.prdtid limit 1) as prdtname,(select specific from products where products.id=onway_products.prdtid limit 1) as specific, quantity, remark from onway_products WHERE cstmid in (select id from customers where customers.cstmname LIKE ?) AND prdtid in (select id from products where products.prdtname LIKE ? AND products.specific LIKE ?) ", "%"+args["cstmName"]+"%","%"+args["prdtName"]+"%","%"+args["specific"]+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	onWayProduct := OnWayProduct{}
	for rows.Next() {
		
		if err = rows.Scan(&onWayProduct.Id, &onWayProduct.CstmName, &onWayProduct.PrdtName,&onWayProduct.Specific,&onWayProduct.Quantity, &onWayProduct.Remark); err != nil {
			return
		}
		onWayProducts = append(onWayProducts, onWayProduct)
	}


	return
}

func SelectOnWayInvoices(args map[string]string) (onWayInvoices []OnWayInvoice, err error) {
	rows, err := Db.Query("select id, (select cstmname from customers where customers.id=onway_invoices.cstmid limit 1) as cstmname,(select prdtname from products where products.id=onway_invoices.prdtid limit 1) as prdtname,(select specific from products where products.id=onway_invoices.prdtid limit 1) as specific, amount, remark from onway_invoices WHERE cstmid in (select id from customers where customers.cstmname LIKE ?) AND prdtid in (select id from products where products.prdtname LIKE ? AND products.specific LIKE ?) ", "%"+args["cstmName"]+"%","%"+args["prdtName"]+"%","%"+args["specific"]+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	onWayInvoice := OnWayInvoice{}
	for rows.Next() {
		
		if err = rows.Scan(&onWayInvoice.Id, &onWayInvoice.CstmName, &onWayInvoice.PrdtName,&onWayInvoice.Specific,&onWayInvoice.Amount, &onWayInvoice.Remark); err != nil {
			return
		}
		onWayInvoices = append(onWayInvoices, onWayInvoice)
	}


	return
}

func GetAllOnWayInvoices() (onWayInvoices []OnWayInvoice, err error) {
	rows, err := Db.Query("select id, (select cstmname from customers where customers.id=onway_invoices.cstmid limit 1) as cstmname,(select prdtname from products where products.id=onway_invoices.prdtid limit 1) as prdtname,(select specific from products where products.id=onway_invoices.prdtid limit 1) as specific, amount, remark from onway_invoices")
	defer rows.Close()
	if err != nil {
		return
	}
	onWayInvoice := OnWayInvoice{}
	for rows.Next() {
		
		if err = rows.Scan(&onWayInvoice.Id, &onWayInvoice.CstmName, &onWayInvoice.PrdtName,&onWayInvoice.Specific,&onWayInvoice.Amount, &onWayInvoice.Remark); err != nil {
			return
		}
		onWayInvoices = append(onWayInvoices, onWayInvoice)
	}


	return
}



func GetProductNS() (products []Product, err error) {

	rows, err := Db.Query("SELECT prdtname, specific FROM products")
	defer rows.Close()
	if err != nil {
		return
	}
	product := Product{}
	for rows.Next() {
		
		if err = rows.Scan(&product.PrdtName, &product.Specific); err != nil {
			return
		}
		products = append(products, product)
	}

	return
}




func GetCustomerName() (customers []Customer, err error) {

	rows, err := Db.Query("SELECT cstmname FROM customers")
	defer rows.Close()
	if err != nil {
		return
	}
	customer := Customer{}
	for rows.Next() {
		
		if err = rows.Scan(&customer.CstmName); err != nil {
			return
		}
		customers = append(customers, customer)
	}

	return
}
