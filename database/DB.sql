-- Изменил название БД в соответствии с нашим соглашением о доступе к БД
create DATABASE mailings;
use mailings;

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
    created_at   datetime DEFAULT CURRENT_TIMESTAMP 
)

-- Добавил m2m таблицу
create table users_messages (
    user_id int UNSIGNED,
    message_id int UNSIGNED,
    PRIMARY KEY (user_id, message_id),
    FOREIGN KEY (user_id) REFERENCES ON users(id),
    FOREIGN KEY (message_id) REFERENCES ON messages(id)
)