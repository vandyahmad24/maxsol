create table if not exists cake
(
    id          int auto_increment primary key,
    title       varchar(255),
    description varchar(255),
    rating      float,
    created_at  datetime(5)  null,
    updated_at  datetime(5)  null
);

