CREATE TABLE IF NOT EXISTS public."user" (
    user_id uuid NOT NULL,
    password character varying(300) NOT NULL,
    login character varying(100) NOT NULL,
    PRIMARY KEY (user_id),
    CONSTRAINT unique_user_name UNIQUE (login)
);

-- CREATE TABLE IF NOT EXISTS order(
-- id serial PRIMARY KEY,
-- source VARCHAR(255),
-- destination VARCHAR(255),
-- original VARCHAR(255),
-- translation VARCHAR(255)
-- );