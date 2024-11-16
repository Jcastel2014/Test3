CREATE TABLE status (
    id SERIAL PRIMARY KEY,
    name VARCHAR (20) NOT NULL
);

insert into status(name) values ('Completed');
insert into status(name) values ('Currently Reading');


DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash bytea NOT NULL,
    activated bool NOT NULL
);

DROP TABLE IF EXISTS readList;
CREATE TABLE readList (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_by INT REFERENCES users(id),
    status INT REFERENCES status(id)

);





CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    isbn VARCHAR(255) UNIQUE,
    publication_date DATE,
    genre VARCHAR(100),
    description TEXT,
    average_rating DECIMAL(3,2)
);

CREATE TABLE authors(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE book_authors (
    id SERIAL PRIMARY KEY,
    book_id INT REFERENCES books(id) ON DELETE CASCADE,
    author_id INT REFERENCES authors(id)
);

CREATE TABLE book_list (
    id SERIAL PRIMARY KEY,
    book_id INT REFERENCES books(id) ON DELETE CASCADE,
    list_id INT REFERENCES readList(id)
);