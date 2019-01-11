
create database mytestdb;
 
create table tbl_user (
 	uid int primary key,
 	username varchar(15) not null,
 	age varchar(10) not null,
	memo varchar(10),
	remark varchar(32)
 )

 show variables like 'sql_safe_updates';
 set SQL_SAFE_UPDATES = 0