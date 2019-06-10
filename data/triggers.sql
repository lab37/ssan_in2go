/* 调整合同将会影响在途发票、在途货物和应收帐款 */
DROP TRIGGER add_onway_pruducts_invoice_debts;
DROP TRIGGER del_onway_invoice_debts;
DROP TRIGGER upt_onway_pruducts_invoice_debts;
DROP TRIGGER del_onway_invoices;
DROP TRIGGER add_onway_invoices;
DROP TRIGGER upt_onway_invoices;
DROP TRIGGER add_stocks_onwayproducs4in;
DROP TRIGGER del_stocks_onwayproducs4in;
DROP TRIGGER upt_stocks_onwayproducs4in;
DROP TRIGGER add_stocks_onwayproducs4out;
DROP TRIGGER del_stocks_onwayproducs4out;
DROP TRIGGER upt_stocks_onwayproducs4out;
DROP TRIGGER del_debts;
DROP TRIGGER add_debts;
DROP TRIGGER upt_debts;


create TRIGGER add_onway_pruducts_invoice_debts after
  insert
  ON
  contracts
BEGIN
    insert or ignore into onway_products (cstmid, prdtid)  values  (new.cstmid, new.prdtid);
    update onway_products set quantity=(quantity+new.quantity) where cstmid=new.cstmid and prdtid=new.prdtid;

    insert or ignore into onway_invoices (cstmid, prdtid)  values  (new.cstmid, new.prdtid);
    update onway_invoices set amount=(amount + new.quantity*new.price) where cstmid=new.cstmid and prdtid=new.prdtid;

    insert or ignore into debts (cstmid)  values  (new.cstmid);
    update debts set amount=(amount+new.quantity*new.price) where cstmid=new.cstmid;

END;

create TRIGGER del_onway_invoice_debts after
delete
ON
contracts
BEGIN
  update onway_products set quantity=(quantity-old.quantity) where cstmid=old.cstmid and prdtid=old.prdtid;
  update onway_invoices set amount=(amount - old.quantity*old.price) where cstmid=old.cstmid and prdtid=old.prdtid;
  update debts set amount=(amount-old.quantity*old.price) where cstmid=old.cstmid;
END;

create TRIGGER upt_onway_pruducts_invoice_debts after
update
ON
contracts
BEGIN
  update onway_products set quantity=(quantity-old.quantity+new.quantity) where cstmid=old.cstmid and prdtid=old.prdtid;
  update onway_invoices set amount=(amount - old.quantity*old.price + new.quantity*new.price) where cstmid=old.cstmid and prdtid=old.prdtid;
  update debts set amount = (amount - old.quantity*old.price + new.quantity*new.price) where cstmid=new.cstmid;
END;





/* 调整发票将会影响在途发票 */
create TRIGGER del_onway_invoices after
insert
ON
invoices
BEGIN
  update onway_invoices set amount = (amount-new.quantity*new.price) where cstmid=new.cstmid and prdtid=new.prdtid;  
END;

create TRIGGER add_onway_invoices after
delete
ON
invoices
BEGIN
  update onway_invoices set amount = (amount+old.quantity*old.price) where cstmid=old.cstmid and prdtid=old.prdtid;
END;

create TRIGGER upt_onway_invoices after
update
ON
invoices
BEGIN
  update onway_invoices set amount = (amount+old.quantity*old.price-new.quantity*new.price) where cstmid=old.cstmid and prdtid=old.prdtid;
END;


/* 入库将会影响库存和在途货物 */
create TRIGGER add_stocks_onwayproducs4in after
  insert
  ON
  instocks
  BEGIN
    insert or  ignore into stocks (prdtid,mac,sn)  values ( new.prdtid,new.mac,new.sn);
    update stocks set quantity = quantity+new.quantity where prdtid=new.prdtid and mac=new.mac and sn=new.sn;

    update onway_products set quantity=(quantity-new.quantity) where cstmid=new.cstmid and prdtid=new.prdtid; 
END;

create TRIGGER del_stocks_onwayproducs4in after
  delete
  ON
  instocks
  BEGIN
  update stocks set quantity = quantity-old.quantity where prdtid=old.prdtid  and mac=old.mac and sn=old.sn;
  update onway_products set quantity=(quantity+old.quantity) where cstmid=old.cstmid and prdtid=old.prdtid;
END;

create TRIGGER upt_stocks_onwayproducs4in after
  update
  ON
  instocks
  BEGIN
  update stocks set quantity = quantity-old.quantity+new.quantity where prdtid=old.prdtid and mac=old.mac and sn=old.sn;
  update onway_products set quantity=(quantity+old.quantity-new.quantity) where cstmid=old.cstmid and prdtid=old.prdtid;
END;



/* 出库将会影响库存和在途货物 */
create TRIGGER add_stocks_onwayproducs4out after
  insert
  ON
  outstocks
  BEGIN
    update stocks set quantity = quantity-new.quantity where prdtid=new.prdtid and  mac=new.mac and sn=new.sn;

    update onway_products set quantity=(quantity-new.quantity) where cstmid=new.cstmid and  prdtid=new.prdtid; 
END;

create TRIGGER del_stocks_onwayproducs4out after
  delete
  ON
  outstocks
  BEGIN
  update stocks set quantity = quantity+old.quantity where prdtid=old.prdtid and mac=old.mac and sn=old.sn;
  update onway_products set quantity=(quantity+old.quantity) where cstmid=old.cstmid and  prdtid=old.prdtid;
END;

create TRIGGER upt_stocks_onwayproducs4out after 
  update
  ON
  outstocks
  BEGIN
  update stocks set quantity = quantity+old.quantity-new.quantity where prdtid=old.prdtid  and mac=old.mac and sn=old.sn;
  update onway_products set quantity=(quantity+old.quantity-new.quantity) where cstmid=old.cstmid and prdtid=old.prdtid;
END;




/* 回款将会影响欠款 */
create TRIGGER del_debts after
insert 
ON
payments
BEGIN
  update debts set amount = amount-new.amount where cstmid=new.cstmid;
END;

create TRIGGER add_debts after
delete 
ON
payments
BEGIN
  update debts set amount = amount+old.amount where cstmid=old.cstmid;
END;

create TRIGGER upt_debts after
update 
ON
payments
BEGIN
  update debts set amount = amount + old.amount -new.amount where cstmid=old.cstmid;
END;