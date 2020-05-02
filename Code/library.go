package main

import (
	"fmt"

	// mysql connector
	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
)

const (
	User     = "XuYipei"
	Password = "123456"
	DBName   = "ass3"
)

type Library struct {
	db *sqlx.DB
}
type Book struct {
	title, author, ISBN string
}
type Student struct {
	student_id string
}
type Borrow struct {
}

func (lib *Library) ConnectDB() {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", User, Password, DBName))
	if err != nil {
		panic(err)
	}
	lib.db = db
	//	defer db.Close()
}

// CreateTables created the tables in MySQL
func (lib *Library) CreateTables() error {
	s := fmt.Sprintf("CREATE TABLE books(title CHAR(32), author CHAR(32), ISBN CHAR(32), PRIMARY KEY(ISBN)) ")
	stmt, err := lib.db.Prepare(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	stmt.Exec()
	stmt.Close()
	s = fmt.Sprintf("CREATE TABLE students(student_id CHAR(32), PRIMARY KEY(student_id))")
	stmt, err = lib.db.Prepare(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	stmt.Exec()
	stmt.Close()
	s = fmt.Sprintf("CREATE TABLE borrows_current(student_id CHAR(32), ISBN CHAR(32), deadline DATE, extend_times INT(4), PRIMARY KEY (student_id, ISBN), FOREIGN KEY (ISBN) REFERENCES books(ISBN), FOREIGN KEY (student_id) REFERENCES students(student_id))")
	stmt, err = lib.db.Prepare(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	stmt.Exec()
	stmt.Close()
	s = fmt.Sprintf("CREATE TABLE borrows_history(student_id CHAR(32), ISBN CHAR(32), deadline DATE, extend_times INT(4), PRIMARY KEY (student_id, ISBN), FOREIGN KEY (ISBN) REFERENCES books(ISBN), FOREIGN KEY (student_id) REFERENCES students(student_id))")
	stmt, err = lib.db.Prepare(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	stmt.Exec()
	defer stmt.Close()
	return nil
}

// AddBook add a book into the library
func (lib *Library) AddBook(title, author, ISBN string) error {
	s := fmt.Sprintf("INSERT INTO books(title, author, ISBN) VALUES ('%s', '%s', '%s')", title, author, ISBN)
	stmt, err := lib.db.Prepare(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	stmt.Exec()
	defer stmt.Close()
	return nil
}

// Add a student account into the library
func (lib *Library) RemoveBook(ISBN string) error {
	s := fmt.Sprintf("DELETE FROM books WHERE ISBN = '%s'", ISBN)
	stmt, err := lib.db.Prepare(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	stmt.Exec()
	defer stmt.Close()
	return nil
}

// Add a student account into the library
func (lib *Library) AddAccount(student_id string) error {
	s := fmt.Sprintf("INSERT INTO students VALUES ('%s')", student_id)
	stmt, err := lib.db.Prepare(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	stmt.Exec()
	defer stmt.Close()
	return nil
}

//Find a book given ISBN in the library
func (lib *Library) FindBookISBN(ISBN string) ([]struct{ title, author, ISBN string }, error) {
	s := fmt.Sprintf("SELECT * FROM books WHERE ISBN = '%s'", ISBN)
	rows, err := lib.db.Query(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	var res []struct{ title, author, ISBN string }
	for rows.Next() {
		var title, author, ISBN string
		rows.Scan(&title, &author, &ISBN)
		res = append(res, Book{title, author, ISBN})
	}
	rows.Close()
	return res, nil
}

//Find a book given title in the library
func (lib *Library) FindBookTitle(title string) ([]struct{ title, author, ISBN string }, error) {
	s := fmt.Sprintf("SELECT * FROM books WHERE title = '%s'", title)
	rows, err := lib.db.Query(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	var res []struct{ title, author, ISBN string }
	for rows.Next() {
		var title, author, ISBN string
		rows.Scan(&title, &author, &ISBN)
		res = append(res, Book{title, author, ISBN})
	}
	rows.Close()
	return res, nil
}

//Find a book given ISBN in the library
func (lib *Library) FindBookAuthor(ISBN string) ([]struct{ title, author, ISBN string }, error) {
	s := fmt.Sprintf("SELECT * FROM books WHERE author = '%s'", ISBN)
	rows, err := lib.db.Query(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	var res []struct{ title, author, ISBN string }
	for rows.Next() {
		var title, author, ISBN string
		rows.Scan(&title, &author, &ISBN)
		res = append(res, Book{title, author, ISBN})
	}
	rows.Close()
	return res, nil
}

//Borrow a book given ISBN from a student account given student_id at the time date in the library
func (lib *Library) BorrowBook(ISBN, student_id, date, current string) error {
	s := fmt.Sprintf("SELECT COUNT(*) FROM borrows_current WHERE student_id = '%s' AND deadline > %s", student_id, current)
	rows, err := lib.db.Query(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	var overdue int
	for rows.Next() {
		rows.Scan(&overdue)
	}
	rows.Close()
	if overdue > 3 {
		return nil
	}

	s = fmt.Sprintf("INSERT INTO borrows_current(ISBN, student_id, deadline, extend_times) VALUES ('%s', '%s', '%s', 0)", ISBN, student_id, date)
	stmt, err := lib.db.Prepare(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	stmt.Exec()
	defer stmt.Close()
	return nil
}

//Query borrow history from student account given student_id
func (lib *Library) QueryHistory(student_id string) ([]string, error) {
	s := fmt.Sprintf("SELECT ISBN FROM borrows_history WHERE student_id = '%s'", student_id)
	rows, er := lib.db.Query(s)
	if er != nil {
		fmt.Println(s)
		panic(er)
	}
	var res []string
	for rows.Next() {
		var ISBN string
		rows.Scan(&ISBN)
		res = append(res, ISBN)
	}
	rows.Close()

	s = fmt.Sprintf("SELECT ISBN FROM borrows_current WHERE student_id = '%s'", student_id)
	rows, er = lib.db.Query(s)
	if er != nil {
		fmt.Println(s)
		panic(er)
	}
	for rows.Next() {
		var ISBN string
		rows.Scan(&ISBN)
		res = append(res, ISBN)
	}
	rows.Close()
	return res, nil
}

//Query books borrowed but not returned given student account student_id
func (lib *Library) QueryBookNotReturned(student_id string) ([]string, error) {
	s := fmt.Sprintf("SELECT ISBN FROM borrows_current WHERE student_id = '%s'", student_id)
	rows, er := lib.db.Query(s)
	if er != nil {
		fmt.Println(s)
		panic(er)
	}
	var res []string
	for rows.Next() {
		var ISBN string
		rows.Scan(&ISBN)
		res = append(res, ISBN)
	}
	rows.Close()
	return res, nil
}

//Query deadline given a book ISBN
func (lib *Library) QueryDeadline(ISBN string) ([]string, error) {
	s := fmt.Sprintf("SELECT deadline FROM borrows_current WHERE ISBN = '%s'", ISBN)
	rows, er := lib.db.Query(s)
	if er != nil {
		fmt.Println(s)
		panic(er)
	}
	var res []string
	for rows.Next() {
		var deadline string
		rows.Scan(&deadline)
		res = append(res, deadline)
	}
	rows.Close()
	return res, nil
}

//Extend deadline given ISBN
func (lib *Library) ExtendDeadline(ISBN, date string) error {
	s := fmt.Sprintf("UPDATE borrows_current SET deadline = '%s', extend_times = extend_times + 1 WHERE extend_times < 3 AND ISBN = '%s'", date, ISBN)
	stmt, err := lib.db.Prepare(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	stmt.Exec()
	defer stmt.Close()
	return nil
}

//Query the overdue books
func (lib *Library) QueryOverdue(student_id, date string) ([]string, error) {
	s := fmt.Sprintf("SELECT ISBN FROM borrows_current WHERE deadline > '%s' AND student_id = '%s'", date, student_id)
	rows, er := lib.db.Query(s)
	if er != nil {
		fmt.Println(s)
		panic(er)
	}
	var res []string
	for rows.Next() {
		var ISBN string
		rows.Scan(&ISBN)
		res = append(res, ISBN)
	}
	defer rows.Close()
	return res, nil
}

//Return a Book
func (lib *Library) ReturnBook(student_id, ISBN string) error {
	s := fmt.Sprintf("INSERT INTO borrows_history (SELECT * FROM borrows_current WHERE borrows_current.student_id = '%s' AND borrows_current.ISBN = '%s')", student_id, ISBN)
	stmt, err := lib.db.Prepare(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	stmt.Exec()
	stmt.Close()
	s = fmt.Sprintf("DELETE FROM borrows_current WHERE borrows_current.student_id = '%s' AND borrows_current.ISBN = '%s'", student_id, ISBN)
	stmt, err = lib.db.Prepare(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	stmt.Exec()
	defer stmt.Close()
	return nil
}

func main() {
	fmt.Println("Welcome to the Library Management System!")
	/*
		lib := new(Library)
		lib.ConnectDB()
		lib.CreateTables()
		lib.db.Close()
	*/
}
