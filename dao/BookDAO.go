package dao

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/giorg/goapi/database"
	"github.com/giorg/goapi/model"
)

//Get database connection
var db, _ = database.GetDb()

//Get All Books
func GetBooks() ([]model.Book, error) {

	rows, err := db.Query("SELECT ID, Isbn, Title, Author FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An book slice to hold data from returned rows.
	var books []model.Book

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Isbn, &book.Title,
			&book.Author); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

//Get Single Book
func GetBook(id string) (model.Book, error) {

	rows, err := db.Query("SELECT ID, Isbn, Title, Author FROM books WHERE id=?", id)
	if err != nil {
		return model.Book{}, err
	}
	defer rows.Close()

	// An book slice to hold data from returned rows.
	var books []model.Book

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Isbn, &book.Title,
			&book.Author); err != nil {
			return book, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return model.Book{}, err
	}
	return books[0], nil
}

//Create Book
func CreateBook(book model.Book) (model.Book, error) {
	query := "INSERT INTO books (Isbn, Title, Author) VALUES (?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return model.Book{}, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, book.Isbn, book.Title, book.Author)
	if err != nil {
		log.Printf("Error %s when inserting row into books table", err)
		return model.Book{}, err
	}

	bookId, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error %s when cast book id in string", err)
		return model.Book{}, err
	}
	strBookId := strconv.FormatInt(bookId, 10)
	return model.Book{ID: strBookId, Isbn: book.Isbn, Title: book.Title, Author: book.Author}, nil
}

//Update Book
func UpdateBook(book model.Book) (model.Book, error) {
	query := "UPDATE books SET Isbn = ?, Title = ?, Author = ? WHERE ID = ?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return model.Book{}, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, book.Isbn, book.Title, book.Author, book.ID)
	if err != nil {
		log.Printf("Error %s when executing SQL statement", err)
		return model.Book{}, err
	}
	return model.Book{ID: book.ID, Isbn: book.Isbn, Title: book.Title, Author: book.Author}, nil
}

//Delete Book
func DeleteBook(id string) error {
	query := "DELETE FROM books WHERE ID = ?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)

	if err != nil {
		log.Printf("Error %s when executing SQL statement", err)
		return err
	}
	return nil
}
