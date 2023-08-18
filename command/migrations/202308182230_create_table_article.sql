-- +migrate Up
CREATE TABLE article_db.article (
                                    id INT auto_increment NOT NULL,
                                    author varchar(100) NULL,
                                    title varchar(250) NULL,
                                    body TEXT NULL,
                                    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
                                    CONSTRAINT article_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;

-- +migrate Down
DROP TABLE article_db.article;