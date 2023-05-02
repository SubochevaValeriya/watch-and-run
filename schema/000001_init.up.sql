CREATE TABLE event

(
    id serial not null primary key unique,
    path varchar(255) not null unique,
    type varchar(255) not null unique,
    time date not null,
);