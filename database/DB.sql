create DATABASE fp_DB;
use fp_DB;
create table users (
    id          INTEGER PRIMARY KEY AUTO_INCREMENT,
    name        varchar(255),
    telegram_id INT,
    first_name  varchar(255),
    last_name   varchar(255),
    chat_id     INT,
    created_at  datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at  datetime,
    deleted_at  datetime
);

create table messages (
    id           INT PRIMARY KEY AUTO_INCREMENT,
    msg_text     TEXT,
    user_chat_id INTEGER FOREIGN KEY REFERENCES ON users(chat_id),
    created_at   datetime DEFAULT CURRENT_TIMESTAMP 
)