#!/bin/bash

/etc/init.d/mysql start
echo "CREATE USER 'website'@'%' IDENTIFIED BY ''" | mysql -u root
echo "CREATE DATABASE website" | mysql -u root
echo "GRANT ALL PRIVILEGES ON website.* TO 'website'@'%' WITH GRANT OPTION;" | mysql -u root
cat /sql.sql | mysql -u root website
/etc/init.d/mysql stop

/usr/sbin/mysqld
