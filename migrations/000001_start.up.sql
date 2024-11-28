DROP TABLE IF EXISTS status;
CREATE TABLE status (
    id SERIAL PRIMARY KEY,
    name VARCHAR (20) NOT NULL
);

insert into status(name) values ('Completed');
insert into status(name) values ('Currently Reading');

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    username text NOT NULL,
    email VARCHAR (255) NOT NULL,
    password_hash bytea NOT NULL,
    activated bool NOT NULL,
    version integer NOT NULL DEFAULT 1
);

DROP TABLE IF EXISTS readList;
CREATE TABLE readList (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_by INT REFERENCES users(id),
    status INT REFERENCES status(id)

);




DROP TABLE IF EXISTS books;
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


DROP TABLE IF EXISTS book_authors;
CREATE TABLE book_authors (
    id SERIAL PRIMARY KEY,
    book_id INT REFERENCES books(id) ON DELETE CASCADE,
    author_id INT REFERENCES authors(id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS book_list;
CREATE TABLE book_list (
    id SERIAL PRIMARY KEY,
    book_id INT REFERENCES books(id) ON DELETE CASCADE,
    list_id INT REFERENCES readList(id) ON DELETE CASCADE
);


--     - Book Reviews { id, book id, user id, rating, the actual review of the book, review date }

DROP TABLE IF EXISTS book_reviews;
CREATE TABLE book_reviews (
    id SERIAL PRIMARY KEY,
    book_id INT REFERENCES books(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id),
    review text NOT NULL,
    rating DECIMAL(3,2),
    created_at DATE
);
DROP TABLE IF EXISTS tokens;

CREATE TABLE tokens (
    hash bytea PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    expiry timestamp(0) WITH TIME ZONE NOT NULL,
    scope text NOT NULL
);
