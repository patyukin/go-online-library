-- +goose Up
CREATE TABLE promotions
(
	id          int auto_increment primary key,
	name        varchar(255) not null,
	description varchar(255) not null,
	comment     varchar(255) null,
	status      tinyint      not null default 0,
	type        enum('oncetime', 'permanent') not null default 'oncetime',
	created_at  datetime     not null,
	updated_at  datetime null,
	deleted_at  datetime null
);

CREATE TABLE filters
(
	id              int auto_increment primary key,
	min_age         int null,
	max_age         int null,
	register_date   date null,
	last_activity   datetime null,
	notify_datetime datetime null,
	created_at      datetime not null,
	updated_at      datetime null,
	deleted_at      datetime null
);

CREATE TABLE promotions_filters
(
	promotion_id int not null,
	filter_id    int not null,
	primary key (promotion_id, filter_id),
	CONSTRAINT `FK-promotion_filters-promotion_id`
		FOREIGN KEY (promotion_id) REFERENCES promotions (id)
			ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT `FK-promotion_filters-filter_id`
		FOREIGN KEY (filter_id) REFERENCES filters (id)
			ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE directories
(
	id         int auto_increment primary key,
	name       varchar(255) not null,
	created_at datetime     not null,
	updated_at datetime null,
	deleted_at datetime null
);

CREATE TABLE promotions_directories
(
	promotion_id int not null,
	directory_id int not null,
	primary key (promotion_id, directory_id),
	CONSTRAINT `FK-promotion_directories-promotion_id`
		FOREIGN KEY (promotion_id) REFERENCES promotions (id)
			ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT `FK-promotion_directories-directory_id`
		FOREIGN KEY (directory_id) REFERENCES directories (id)
			ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE notify
(
	id           int auto_increment primary key,
	notify_time  datetime not null,
	promotion_id int      not null,
	created_at   datetime not null,
	updated_at   datetime null,
	deleted_at   datetime null,
	CONSTRAINT `FK-notify-promotion_id`
		FOREIGN KEY (promotion_id) REFERENCES promotions (id)
			ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose Down
DROP TABLE directories;
DROP TABLE filters;
DROP TABLE promotions;
