
CREATE TABLE public.sign_ups (
    id bigserial NOT NULL,
    user_name varchar(255) NOT NULL,
    user_email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    CONSTRAINT sign_up_pkey PRIMARY KEY (id),
    CONSTRAINT sign_up_email_key UNIQUE (user_email)
);


