--------------------

-- CREATE SECTION --

--------------------

create table
    customers (
        id serial primary key not null,
        customer_name char(50) not null
    );

create table
    products (
        id serial primary key not null,
        name char(50) not null
    );

create table
    orders (
        order_id serial primary key not null,
        customer_id int references customers(id) not null,
        product_id int references products(id) not null,
        order_date date not null,
        total float not null
    );

--------------------

-- INSERT SECTION --

--------------------

insert into
    customers (customer_name)
values ('Boby'), ('Supratman'), ('Waluyo');

insert into products (name)
values ('eMall'), ('eFish'), ('eFresh');

insert into
    orders (
        customer_id,
        product_id,
        order_date,
        total
    )
values (1, 2, '2022-08-24', 4), (2, 3, '2022-08-25', 6), (3, 1, '2022-08-24', 7), (1, 3, '2022-08-26', 9);

--------------------

-- UPDATE SECTION --

--------------------

update customers set customer_name = 'Susi' where id = 2;

update products set name = 'eFeeder' where id = 3;

update orders set total = 10 where order_id = 4;

--------------------

-- DELETE SECTION --

--------------------

delete from orders where order_id = 3;

--------------------

-- SELECT SECTION --

--------------------

select * from customers;

select * from products;

select * from orders;

------------------

-- JOIN SECTION --

------------------

select
    o.order_id,
    c.customer_name as customer_name,
    p.name as product_name,
    o.order_date,
    o.total
from orders o
    inner join customers c on o.customer_id = c.id
    inner join products p on o.product_id = p.id;