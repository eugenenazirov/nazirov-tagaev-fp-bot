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
    created_at   datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at  datetime,
    deleted_at  datetime
);
