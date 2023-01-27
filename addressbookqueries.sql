create database AddressBook;
use AddressBook;

CREATE TABLE contact(
  first_name VARCHAR(128) NOT NULL,
  last_name varchar(128) not null,
  address varchar(200),
  city varchar(100),
  state varchar(100),
  phone_number varchar(10) not null,
  email VARCHAR(255),
  PRIMARY KEY (first_name,last_name)  #change PK
);

select * from contact;
#delete from contact where first_name = 'H' and last_name= 'M';

