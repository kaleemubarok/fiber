CREATE EXTENTION IF NOT EXIST "uui-ossp";

SET TIMEZONE="Asia/Jakarta";

CREATE TABLE books (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW,
    updated_at TIMESTAMP NULL,
    title VARCHAR (255) NOT NULL,
    author VARCHAR (255) NOT NULL,
    book_status INT NOT NULL,
    book_attrs JSONB NOT NULL
);

CREATE INDEX active_books ON books (tit;e) where book_status = 1;