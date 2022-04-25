#!/bin/sh

CMD_MYSQL="mysql -u root -proot_password -D echo_api_sample"

$CMD_MYSQL -e "create table if not exists products (
  id          bigint unsigned auto_increment primary key,
  title       varchar(100) not null,
  description varchar(2000) not null,
  price       int unsigned not null,
  created_at  datetime not null default current_timestamp,
  updated_at  datetime not null default current_timestamp on update current_timestamp
);"
