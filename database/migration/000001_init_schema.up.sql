
CREATE TABLE public.sign_ups (
    user_name varchar(255) NOT NULL,
    user_email varchar(255) NOT NULL,
    users_id int8 NOT NULL,
    password varchar(255) NOT NULL,
    CONSTRAINT sign_up_email_key UNIQUE (user_email)
);

CREATE TABLE public.ecom_users (
    ecom_id varchar(255) NOT NULL,
    account_name varchar(255) NOT NULL,
    wallet_amount int8 NOT NULL DEFAULT 0,
    delivery_address varchar(500) NULL,
    users_id int8 NOT NULL,
    CONSTRAINT ecom_users_pkey PRIMARY KEY (ecom_id)
);

CREATE TABLE public.card_details (
    encrypted_data varchar(5000) NOT NULL,
    ecom_id varchar(255) NOT NULL
);

CREATE TABLE public.orders (
    order_id varchar(255) NOT NULL,
    order_status varchar(255) NOT NULL,
    order_amount int8 NOT NULL,
    order_date timestamp NOT NULL,
    order_items jsonb NOT NULL,
    order_address jsonb NOT NULL,
    order_payment jsonb NOT NULL,
    ecom_id varchar(255) NOT NULL,
    users_id int8 NOT NULL,
    CONSTRAINT orders_order_id_key UNIQUE (order_id)
);

CREATE TABLE public.wish_list (
    wish_list_id varchar(255) NOT NULL,
    wish_list_items jsonb NOT NULL,
    ecom_id varchar(255) NOT NULL,
    users_id int8 NOT NULL,
    CONSTRAINT wish_list_wish_list_id_key UNIQUE (wish_list_id)
);

CREATE TABLE public.cart_items (
    cart_item_id varchar(255) NOT NULL,
    cart_item jsonb NOT NULL,
    ecom_id varchar(255) NOT NULL,
    users_id int8 NOT NULL,
    CONSTRAINT cart_items_cart_item_id_key UNIQUE (cart_item_id)
);


CREATE TABLE public.items (
    item_id int8 NOT NULL,
    item_title varchar(1000) NOT NULL,
    item_price DOUBLE PRECISION NOT NULL,
    item_description varchar(2000) NOT NULL,
    item_category varchar(1000) NOT NULL,
    item_image varchar(1000) NOT NULL,
    item_rating DOUBLE PRECISION NOT NULL,
    item_count int8 NOT NULL,
    CONSTRAINT items_item_id_key UNIQUE (item_id)
);

