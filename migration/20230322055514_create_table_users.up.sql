create table if not exists users
(
    id          int auto_increment primary key,
    name       varchar(255),
    password varchar(255),
    created_at  datetime(5)  null,
    updated_at  datetime(5)  null
    );

