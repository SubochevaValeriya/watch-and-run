CREATE TABLE event

(
    id serial not null primary key unique,
    path varchar(255) not null,
    file_name varchar(255) not null,
    type varchar(255) not null,
    time timestamp not null
);

CREATE TABLE launch

(
    id serial not null primary key unique,
    command varchar(255) not null,
    start_time timestamp not null,
    end_time timestamp not null,
    result varchar(255) not null,
    event_id int not null references event(id)
);
