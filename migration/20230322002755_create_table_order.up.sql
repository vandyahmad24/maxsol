create table if not exists orders
(
    id          int auto_increment primary key,
    cake_id int,
    qty int,
    created_at  datetime(5)  null,
    updated_at  datetime(5)  null
    );

