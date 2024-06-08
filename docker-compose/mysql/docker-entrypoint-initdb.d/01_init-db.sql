create database if not exists gamesdb;
create user games@'%' identified by 'password';
grant all on gamesdb.* to games@'%';
