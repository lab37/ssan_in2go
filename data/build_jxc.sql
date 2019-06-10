/* 手动操作的表只有产品、客户、合同、发票、入库、出库、回款。其它的表为自动生成 */
drop table contracts;
drop table instocks;
drop table outstocks;
drop table invoices;
drop table payments;
drop table products;
drop table customers;
drop table stocks;
drop table debts;
drop table onway_products;
drop table onway_invoices;

create table products
(
  id integer primary key not null,
  prdtname text not null,
  specific text not null,
  inventor text not null,
  unit text not null,
  ivtype text DEFAULT '-',
  remark text DEFAULT '-'
);

create table customers
(
  id integer primary key not null,
  cstmname text not null UNIQUE,
  cstmtype text DEFAULT '-',
  city text DEFAULT '-',
  area text DEFAULT '-',
  cstmaddr text DEFAULT '-',
  owner_name text DEFAULT '-',
  telephone text DEFAULT '-',
  police  text DEFAULT '-',
  axis text DEFAULT '-',
  remark text DEFAULT '-'
);

create table contracts
(
  id integer primary key not null,
  ccsn text not null,
  vector text not null,
  create_date text DEFAULT current_date,
  cctype text DEFAULT '-',
  ivtype text DEFAULT '-',
  cstmid integer not null references customers(id),
  prdtid text not null references products(id),
  price real not null,
  quantity integer not null,
  remark text DEFAULT '-'
);


create table instocks
(
  id integer primary key not null,
  cstmid integer not null references customers(id),
  prdtid text not null references products(id),
  mac  text DEFAULT '-',
  sn text DEFAULT '-',
  create_date text DEFAULT current_date,
  quantity integer not null,
  remark text DEFAULT '-'
);

create table outstocks
(
  id integer primary key not null,
  cstmid integer not null references customers(id),
  prdtid text not null references products(id),
  mac  text DEFAULT '-',
  sn text DEFAULT '-',
  create_date text DEFAULT current_date,
  quantity integer not null,
  remark text DEFAULT '-'
);

create table invoices
(
  id integer primary key not null,  
  sn text not null,
  ivtype text not null DEFAULT '收据',
  create_date text DEFAULT current_date,
  cstmid text not null references customers(id),
  prdtid text not null references products(id),
  quantity integer not null,
  price integer not null,
  remark text DEFAULT '-'
);

create table payments
(
  id integer primary key not null,
  cstmid text not null references customers(id),
  create_date text DEFAULT current_date,
  pmtype text not null DEFAULT '电汇',
  amount real not null,
  remark text DEFAULT '-'
);

create table stocks
(
  id integer PRIMARY key not null,
  prdtid text not null references products(id),
  mac  text DEFAULT '-',
  sn text DEFAULT '-',
  quantity integer default 0 check(quantity>=0),
  remark text DEFAULT '-'
);

create table debts
(
  id integer PRIMARY key not null,
  cstmid integer not null references customers(id),
  amount real  default 0,
  remark text DEFAULT '-'
);

create table onway_products
(
  id integer PRIMARY key not null,
  cstmid text not null references customers(id),
  prdtid text not null references products(id),
  quantity INTEGER default 0 check(quantity>=0),
  remark text DEFAULT '-'
);

create table onway_invoices
(
  id integer PRIMARY key not null,
  cstmid text not null references customers(id),
  prdtid text not null references products(id),
  amount real  default 0 check(amount>=0),
  remark text DEFAULT '-'
);
