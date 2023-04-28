Operations in Go with RDBMS.

Start a MySQL instance:

  docker run -d --name go-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=p4ssw0rd mysql

Get the Sakila example database and apply the scripts for the Sakila DB:

  https://dev.mysql.com/doc/index-other.html

To run the application:

  go run .
