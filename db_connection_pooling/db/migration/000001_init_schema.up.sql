CREATE TABLE IF NOT EXISTS products (
    id serial Primary key,
    name varchar(255) not null,
    price numeric(10, 2) not null,
    description text default null
)