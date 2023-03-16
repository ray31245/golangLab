package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// model
type book struct {
	Name string
	ISBN int
}

// DAO
type bookDAO interface {
	Get(int) book
	List() []book
	Insert(book) int64
	Delete(int) int64
}

// DAOImpl
type bookDAOImpl struct {
	db *sql.DB
}

func (b bookDAOImpl) Get(ISBN int) book {
	book := book{}
	rows, err := b.db.Query(`SELECT * FROM books WHERE isbn = %d limit 1;`, ISBN)
	if err != nil {
		return book
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.ISBN, &book.Name)
		if err != nil {
			return book
		}
	}
	return book
}

func (b bookDAOImpl) List() []book {
	books := []book{}
	rows, err := b.db.Query(`SELECT * FROM books;`)
	if err != nil {
		return books
	}
	defer rows.Close()

	book := book{}
	for rows.Next() {
		err = rows.Scan(&book.ISBN, &book.Name)
		if err != nil {
			return books
		}
		books = append(books, book)
	}
	return books
}

func (b bookDAOImpl) Insert(book book) int64 {
	stmt, err := b.db.Prepare(fmt.Sprintf(`INSERT INTO books , name VALUES ("%d","%s");`, book.ISBN, book.Name))
	if err != nil {
		return 0
	}
	defer stmt.Close()

	r, err := stmt.Exec()
	if err != nil {
		return 0
	}

	n, err := r.RowsAffected()
	if err != nil {
		return 0
	}

	return n
}

func (b bookDAOImpl) Delete(ISBN int) int64 {
	stmt, err := b.db.Prepare(fmt.Sprintf(`DELETE FROM books WHERE isbn=%d;`, ISBN))
	if err != nil {
		return 0
	}
	defer stmt.Close()

	r, err := stmt.Exec()
	if err != nil {
		return 0
	}

	n, err := r.RowsAffected()
	if err != nil {
		return 0
	}

	return n
}

type operatorBook struct {
	dao bookDAO
}

func (o operatorBook) FindAllBook() []book {
	return o.dao.List()
}

func main() {
	db, err := sql.Open("mysql", "root:O/14jiQ6LOG4x5r5niA6RHkwKssek/yJ@tcp(172.17.0.3:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	allBooks := operatorBook{bookDAOImpl{db}}.FindAllBook()
	allBooks2 := operatorBook{bookDAOImpl{db}}.dao.List()
	fmt.Println(allBooks)
	fmt.Println(allBooks2)
}
