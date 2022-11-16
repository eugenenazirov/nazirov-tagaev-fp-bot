-- Изменил название БД в соответствии с нашим соглашением о доступе к БД
create DATABASE mailings;
use mailings;

create table users (
    id          INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
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
    id           INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    msg_text     TEXT,
    created_at   datetime DEFAULT CURRENT_TIMESTAMP 
);

-- Добавил m2m таблицу
create table users_messages (
    users_id     INT UNSIGNED,
    messages_id  INT UNSIGNED,
    PRIMARY KEY (users_id, messages_id),
    FOREIGN KEY (users_id) REFERENCES users(id),
    FOREIGN KEY (messages_id) REFERENCES messages(id)
);