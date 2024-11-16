SELECT B.id, B.title, B.isbn, A.name AS author, B.publication_date, B.genre, B.description, B.average_rating
FROM books AS B
INNER JOIN book_authors AS BA 
ON B.id = BA.book_id
INNER JOIN authors AS A 
ON A.id = BA.author_id

	WHERE P.id = $1 OR NOT EXISTS (SELECT 1 FROM products WHERE id = $1)


SELECT B.id, B.title, B.isbn, A.name AS author, B.publication_date, B.genre, B.description, B.average_rating
FROM books AS B
INNER JOIN book_authors AS BA 
ON B.id = BA.book_id
INNER JOIN authors AS A 
ON A.id = BA.author_id
INNER JOIN book_list AS BL
ON BL.book_id = B.id
WHERE BL.list_id = 1;

SELECT R.id, R.name, R.description, U.user_name AS created_by, S.name as status 
FROM readList AS R 
INNER JOIN users AS U 
ON R.created_by = U.id 
INNER JOIN status AS S 
ON R.status = S.id

