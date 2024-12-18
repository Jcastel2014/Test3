include .envrc

.PHONY: run/api
run/api:
	@echo 'Running BookClub API...'
	@go run ./cmd/api -port=3000 -env=production -db-dsn=${BOOKCLUB_DB_DSN}

.PHONY: db/psql
db/psql:
	psql ${BOOKCLUB_DB_DSN}

.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

.PHONY: db/migrations/up
db/migrations/up:
	@echo 'Running up migrations...'
	migrate -path=./migrations -database ${BOOKCLUB_DB_DSN} up

.PHONY: addBook
addBook:
	@echo 'Creating Review'; \
	BODY='{"title":"Animal Farm","isbn":"1","author":"George Orwell","genre":"Political Satire","description":"A satirical novella that allegorizes the events leading up to the Russian Revolution and the rise of Stalinism, told through the story of farm animals overthrowing their human owner.","created_at":"1945-08-17T00:00:00Z"}'; \
	curl -X POST -d "$$BODY" localhost:3000/api/v1/books; \

.PHONY: getAllBooks
getAllBooks:
	@echo 'Displaying Reviews'; \
	curl -i localhost:3000/api/v1/books?${filter}

.PHONY: getBook
getBook:
	@echo 'Displaying Product'; \
	curl -i localhost:3000/api/v1/books/${id} 

.PHONY: putBook
putBook:
	@echo 'Updating Product ${id}'; \
	curl -X PUT localhost:3000/api/v1/books/${id} -d '{"Description":"Updated Description", "genre":"Idk"}'

.PHONY: deleteBook
deleteBook:
	@echo 'Deleting Product'; \
	curl -X DELETE localhost:3000/api/v1/books/${id} 


.PHONY: createList
createList:
	@echo 'Creating List'; \
	BODY='{"name":"test2","description":"test2","created_by":"carlos"}'; \
	curl -X POST -d "$$BODY" localhost:3000/api/v1/lists; \

.PHONY: getAllLists
getAllLists:
	@echo 'Displaying Lists'; \
	curl -i localhost:3000/api/v1/lists?${filter}

.PHONY: listAddBook
listAddBook:
	@echo 'Adding book to list'; \
	BODY='{"bookid":1}'; \
	curl -X POST -d "$$BODY" localhost:3000/api/v1/lists/${id}/books ; \

.PHONY: getList
getList:
	@echo 'Displaying List'; \
	curl -i localhost:3000/api/v1/lists/${id}


.PHONY: updateList
updateList:
	@echo 'Updating List'; \
	curl -X PUT localhost:3000/api/v1/lists/${id} -d '{"status":"Completed", "name":"updateTest2"}'

.PHONY: deleteList
deleteList:
	@echo 'Deleting Product'; \
	curl -X DELETE localhost:3000/api/v1/lists/${id} 

.PHONY: deleteFromList
deleteFromList:
	@echo 'Deleting Product'; \
	curl -X DELETE localhost:3000/api/v1/lists/${id}/books


.PHONY: addBookReview
addBookReview:
	@echo 'Adding book review'; \
	BODY='{"user_id":1, "review":"terrible", "rating":5}'; \
	curl -X POST -d "$$BODY" localhost:3000/api/v1/books/${id}/reviews ; \

.PHONY: getAllReviews
getAllReviews:
	@echo 'Displaying Lists'; \
	curl -i localhost:3000/api/v1/books/${id}/reviews?${filter}

.PHONY: deleteReview
deleteReview:
	@echo 'Deleting Review'; \
	curl -X DELETE localhost:3000/api/v1/reviews/${id}


.PHONY: putReview
putReview:
	@echo 'Updating Product ${id}'; \
	curl -X PUT localhost:3000/api/v1/reviews/${id} -d '{"rating":2.00}'

.PHONY: run/rateLimite/enabled
run/rateLimit,enabled:
	@echo 'Running Product API /w Rate Limit...'
	@go run ./cmd/api -port=3000 -env=development -limiter-burst=5 -limiter-rps=2 -limiter-enabled=true -db-dsn=${PRODUCTS_DB_DSN}

.PHONY: run/rateLimite/disabled
run/rateLimit/disabled:
	@echo 'Running Product API /w Rate Limit...'
	@go run ./cmd/api -port=3000 -env=development -limiter-burst=5 -limiter-rps=2 -limiter-enabled=false -db-dsn=${PRODUCTS_DB_DSN}



.PHONY: displayAllProducts
displayAllProducts:
	@echo 'Deleting Product'; \
	curl -i localhost:3000/displayAllProducts

.PHONY: listProducts
listProducts:
	@echo 'Deleting Product'; \
	curl -i localhost:3000/displayAllProducts?${filter}





	

	



.PHONY: createProduct
createProduct:
	@echo 'Creating Product'; \
    BODY=${CREATEPRODUCT}; \
	curl -i -d "$$BODY" localhost:3000/createProduct ; \
	echo 'create a product'
