CREATE SCHEMA grey;

CREATE TABLE grey.user (
    id serial primary key,
    username varchar(100) not null unique,
    firstname varchar(100) not null,
    lastname varchar(100) not null,
    fullname varchar(200) not null,
    age int not null CHECK (age > 17),
    is_married boolean not null default false,
    password varchar(255) not null
);

CREATE TABLE grey.product (
    id serial primary key,
    name varchar(200) not null unique,
    description text,
    quantity int not null,
    price_date timestamp NOT NULL
);

CREATE TABLE grey.tag (
    id serial primary key,
    name varchar(100) not null unique
);

CREATE TABLE grey.product_tag (
    product_id int references grey.product (id) on delete cascade not null,
    tag_id int references grey.tag (id) on delete cascade not null,
    CONSTRAINT product_tag_pkey PRIMARY KEY (product_id, tag_id)
);

CREATE TABLE grey.price (
    id serial primary key,
    product_id int references grey.product (id) on delete cascade not null,
    price numeric CHECK (price > 0),
    date timestamp not null
);

CREATE TABLE grey.cart (
    id serial primary key,
    user_id int references grey.user (id) on delete cascade not null,
    CONSTRAINT user_id_unq UNIQUE(user_id)
);

CREATE TABLE grey.cart_item (
    product_id int references grey.product (id) on delete cascade not null,
    cart_id int references grey.cart (id) on delete cascade not null,
    price_id int references grey.price (id) on delete cascade not null,
    quantity int CHECK (quantity > 0),
    CONSTRAINT cart_item_pkey PRIMARY KEY (product_id, cart_id, price_id)
);

CREATE TABLE grey.order (
    id serial primary key,
    user_id int references grey.user (id) on delete cascade not null
);

CREATE TABLE grey.order_item (
    product_id int references grey.product (id) on delete cascade not null,
    order_id int references grey.order (id) on delete cascade not null,
    price_id int references grey.price (id) on delete cascade not null,
    quantity int CHECK (quantity > 0),
    CONSTRAINT order_item_pkey PRIMARY KEY (product_id, order_id, price_id)
);