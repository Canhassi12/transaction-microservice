USE tmicroservice;

CREATE TABLE IF NOT EXISTS transactions (
    id char(36) PRIMARY KEY, 
    status varchar(255),
    user_id int,
    order_id int,
    payment_id char(36),
    paid_at timestamp
);