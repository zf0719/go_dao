
create database mytestdb;
 
create table user (
 	uid int primary key,
 	username varchar(10) not null,
 	age varchar(10) not null
 )

 show variables like 'sql_safe_updates';
 set SQL_SAFE_UPDATES = 0