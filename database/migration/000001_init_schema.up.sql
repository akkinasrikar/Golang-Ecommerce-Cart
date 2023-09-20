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
    cart_items varchar(500) NULL,
    users_id int8 NOT NULL,
    CONSTRAINT ecom_users_pkey PRIMARY KEY (ecom_id)
);

CREATE TABLE public.card_details (
    card_id varchar(255) NOT NULL,
    encrypted_data varchar(5000) NOT NULL,
    ecom_id varchar(255) NOT NULL
);

CREATE TABLE public.delivery_addresses (
    address_id varchar(255) NOT NULL,
    house_no varchar(255) NOT NULL,
    street varchar(255) NOT NULL,
    city varchar(255) NOT NULL,
    state varchar(255) NOT NULL,
    pincode varchar(255) NOT NULL,
    ecom_id varchar(255) NOT NULL
);

CREATE TABLE public.orders (
    order_id varchar(255) NOT NULL,
    order_status varchar(255) NOT NULL,
    order_amount int8 NOT NULL,
    order_date varchar(255) NOT NULL,
    order_name varchar(5000) NOT NULL,
    payment_mode varchar(255) NOT NULL,
    delivery_status varchar(255) NOT NULL,
    delivery_date varchar(255) NULL,
    address_id varchar(500) NOT NULL,
    card_id varchar(255) NOT NULL,
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

