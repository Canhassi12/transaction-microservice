CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS transactions;

CREATE TABLE IF NOT EXISTS transactions (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    status varchar(255),
    user_id int,
    order_id int,
    payment_id char(36),
    paid_at timestamp
);
