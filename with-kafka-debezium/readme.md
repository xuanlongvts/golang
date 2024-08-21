1. export DEBEZIUM_VERSION=2.7
2. make run
3. make regis-mysql
4. http://localhost:9089/
5. Click nav Topics and observe

6. User dbever connect to mysql with information:
	host: localhost     port: 33061
	user/pass: mysqluser/mysqlpw

	insert into orders (order_number, order_date, purchaser, quantity, product_id)
	values (1, "2024-08-15", 1001, 10, 106)

	update orders set quantity = 15 where order_number = 1

	delete from orders where order_number = 1