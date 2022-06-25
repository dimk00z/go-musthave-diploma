CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS "public"."user" (
    user_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    password character varying(300) NOT NULL,
    login character varying(100) NOT NULL,
    balance NUMERIC(20, 2) DEFAULT 0 NOT NULL,
    spend NUMERIC(20, 2) DEFAULT 0 NOT NULL,
    PRIMARY KEY (user_id),
    CONSTRAINT unique_user_name UNIQUE (login)
);

CREATE TABLE IF NOT EXISTS "public"."order" (
    order_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    order_number varchar(50) NOT NULL,
    status varchar(50) DEFAULT 'NEW' NOT NULL,
    uploaded_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    accrual real DEFAULT 0 NOT NULL,
    user_id uuid NOT NULL,
    CONSTRAINT pk_tbl_order_id PRIMARY KEY (order_id),
    CONSTRAINT unq_or_number UNIQUE ("order_number"),
    CONSTRAINT fk_tbl_user FOREIGN KEY (user_id) REFERENCES "public"."user"(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS "public"."withdrawal" (
    withdrawal_id uuid DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    order_number VARCHAR (50) NOT NULL UNIQUE,
    processed_at TIMESTAMP,
    withdraw_sum NUMERIC(20, 2) DEFAULT 0,
    CONSTRAINT pk_tbl_withdrawal_id PRIMARY KEY (withdrawal_id),
    CONSTRAINT unq_wd_number UNIQUE ("order_number"),
    CONSTRAINT fk_tbl_user FOREIGN KEY (user_id) REFERENCES "public"."user"(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);