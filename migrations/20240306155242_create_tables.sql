-- +goose Up
CREATE TABLE promotions
(
    id         int auto_increment primary key,
    name       varchar(255) not null,
    active     bool         not null default false,
    count      int          not null default 0,
    created_at timestamp    not null,
    updated_at timestamp    not null,
    deleted_at timestamp
);

CREATE TABLE filters
(
    id           int auto_increment primary key,
    name         varchar(255) not null,
    user_name    varchar(255) null,
    book_name    varchar(255) null,
    author_name  varchar(255) null,
    start_at     time         not null,
    next_after   int          not null,
    promotion_id int          not null,
    created_at   timestamp    not null default current_timestamp,
    updated_at   timestamp null,
    deleted_at   timestamp null,
    CONSTRAINT FK_filters_promotion_id
        FOREIGN KEY (promotion_id) REFERENCES promotions (id)
            ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE directories
(
    id           int auto_increment primary key,
    name         varchar(255) not null,
    created_at   timestamp    not null default current_timestamp,
    updated_at   timestamp null,
    promotion_id int          not null,
    CONSTRAINT `FK-directories-promotion_id`
        FOREIGN KEY (promotion_id) REFERENCES promotions (id)
            ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose Down
DROP TABLE directories;
DROP TABLE filters;
DROP TABLE promotions;
