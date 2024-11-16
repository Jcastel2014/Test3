package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *appDependencies) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(a.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(a.notAllowedResponse)

	// GET    /api/v1/books              # List all books with pagination
	router.HandlerFunc(http.MethodGet, "/api/v1/books", a.GetAllBooks)
	// GET    /api/v1/books/{id}         # Get book details
	router.HandlerFunc(http.MethodGet, "/api/v1/books/:id", a.getBook)
	// POST   /api/v1/books              # Add new book
	router.HandlerFunc(http.MethodPost, "/api/v1/books", a.postBook)
	// PUT    /api/v1/books/{id}         # Update book details
	router.HandlerFunc(http.MethodPut, "/api/v1/books/:id", a.PutBook)
	// DELETE /api/v1/books/{id}         # Delete book
	router.HandlerFunc(http.MethodDelete, "/api/v1/books/:id", a.deleteBook)
	// GET    /api/v1/books/search       # Search books by title/author/genre
	router.HandlerFunc(http.MethodGet, "/api/v1/book/search", a.searchBook)

	// GET    /api/v1/lists              # Get all reading lists
	router.HandlerFunc(http.MethodGet, "/api/v1/lists", a.getAllLists)
	// GET    /api/v1/lists/{id}         # Get specific reading list
	router.HandlerFunc(http.MethodGet, "/api/v1/lists/:id", a.getList)
	// POST   /api/v1/lists              # Create new reading list
	router.HandlerFunc(http.MethodPost, "/api/v1/lists", a.postReadingList)
	// PUT    /api/v1/lists/{id}         # Update reading list
	// DELETE /api/v1/lists/{id}         # Delete reading list
	// POST   /api/v1/lists/{id}/books   # Add book to reading list
	router.HandlerFunc(http.MethodPost, "/api/v1/lists/:id/books", a.listAddBook)
	// DELETE /api/v1/lists/{id}/books   # Remove book from reading list

	// router.HandlerFunc(http.MethodPost, "/createProduct", a.createProduct)
	// router.HandlerFunc(http.MethodGet, "/displayProduct/:id", a.displayProduct)
	// router.HandlerFunc(http.MethodDelete, "/deleteProduct/:id", a.deleteProduct)
	// router.HandlerFunc(http.MethodGet, "/displayAllProducts", a.displayAllProducts)
	// router.HandlerFunc(http.MethodPatch, "/updateProduct/:id", a.updateProduct)

	// router.HandlerFunc(http.MethodPost, "/product/:id/createReview", a.createReview)
	// router.HandlerFunc(http.MethodGet, "/product/:id/getReview/:rid", a.getReview)
	// router.HandlerFunc(http.MethodPatch, "/product/:id/updateReview/:rid", a.updateReview)
	// router.HandlerFunc(http.MethodDelete, "/product/:id/deleteReview/:rid", a.deleteReview)

	// router.HandlerFunc(http.MethodGet, "/reviews", a.GetAllReviews)

	return a.recoverPanic(a.rateLimit(router))
}
