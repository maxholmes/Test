CREATE TABLE Member (
    id  SERIAL NOT NULL,
    first_name varchar(255) not null,
    last_name varchar(255) ,
    email varchar(255),
    PRIMARY KEY (id)
);

select * from Member;