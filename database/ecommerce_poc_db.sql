create table public.tbl_user
(
    id      serial  not null
        constraint users_pkey
            primary key,
    name    varchar not null,
    email   varchar not null,
    address text    not null
);

alter table public.tbl_user
    owner to postgres;

create unique index users_email_idx
    on public.tbl_user (email);

create table public.tbl_product
(
    id          serial       not null
        constraint products_pkey
            primary key,
    code        varchar(200) not null
        constraint products_un
            unique,
    name        varchar(200) not null,
    description text         not null,
    stock       integer      not null,
    price       bigint       not null
);

alter table public.tbl_product
    owner to postgres;

create table public.tbl_cart_item
(
    id         serial                                                not null
        constraint cart_items_pkey
            primary key,
    user_id    smallint                                              not null,
    product_id smallint                                              not null,
    quantity   integer                                               not null,
    date       timestamp(0) with time zone default CURRENT_TIMESTAMP not null
);

alter table public.tbl_cart_item
    owner to postgres;

create table public.tbl_order__cart
(
    id                 serial                                        not null
        constraint tbl_order__cart_pk
            primary key,
    user_id            integer                                       not null,
    total_price        integer,
    order_status       varchar default 'checkout'::character varying not null,
    date               timestamp                                     not null,
    time_limit_payment timestamp(0) with time zone
);

alter table public.tbl_order__cart
    owner to postgres;

create table public.tbl_order__cart_detail
(
    id         serial    not null,
    quantity   integer   not null,
    price      integer   not null,
    order_id   integer   not null,
    date       timestamp not null,
    product_id integer   not null
);

alter table public.tbl_order__cart_detail
    owner to postgres;

