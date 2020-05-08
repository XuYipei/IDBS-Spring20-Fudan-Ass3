package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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
	db                                *sqlx.DB
	User, Password, DBName            string
	book_tot, student_tot, remove_tot int
}
type BOOK struct {
	book_id             int
	title, author, ISBN string
}
type Student struct {
	student_id string
}
type BORROW struct {
	book_id int
	ISBN    string
}

func ReadStr() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func Board(op []string) []string {
	var res []string
	for _, __ := range op {
		fmt.Printf("%s:", __)
		str := ReadStr()
		res = append(res, str)
	}
	return (res)
}

func (lib *Library) ConnectDB() {
	inp := Board([]string{"User", "Password", "DBName"})
	lib.User = inp[0]
	lib.Password = inp[1]
	lib.DBName = inp[2]
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", lib.User, lib.Password, lib.DBName))
	if err != nil {
		panic(err)
	}
	lib.db = db
	//	defer db.Close()
}

func Execute(db *sqlx.DB, op []string) error {
	for _, s := range op {
		stmt, err := db.Prepare(s)
		if err != nil {
			fmt.Println(s)
			panic(err)
		}
		stmt.Exec()
		defer stmt.Close()
	}
	return nil
}

// CreateTables created the tables in MySQL
func (lib *Library) CreateTables() error {
	s0 := fmt.Sprintf("CREATE TABLE books_avail(book_id INT, title CHAR(32), author CHAR(32), ISBN CHAR(32), PRIMARY KEY(book_id))")
	s1 := fmt.Sprintf("CREATE TABLE students(student_id INT, name CHAR(32), password CHAR(32), PRIMARY KEY(student_id))")
	s2 := fmt.Sprintf("CREATE TABLE books_borrow(book_id INT, title CHAR(32), author CHAR(32), ISBN CHAR(32), student_id INT, deadline DATE, extend_times INT(4), PRIMARY KEY (book_id), FOREIGN KEY (student_id) REFERENCES students(student_id))")
	s3 := fmt.Sprintf("CREATE TABLE borrow_logs(book_id INT, student_id INT, ISBN CHAR(32), FOREIGN KEY (student_id) REFERENCES students(student_id))")
	s4 := fmt.Sprintf("CREATE TABLE remove_logs(book_id INT, ISBN CHAR(32), detail CHAR(64), PRIMARY KEY(book_id))")
	err := Execute(lib.db, []string{s0, s1, s2, s3, s4})
	return err
}

// AddBook add a book into the library
func (lib *Library) AddBook(title, author, ISBN string) (int, error) {
	lib.book_tot = lib.book_tot + 1
	s := fmt.Sprintf("INSERT INTO books_avail(book_id, title, author, ISBN) VALUES (%d, '%s', '%s', '%s')", lib.book_tot, title, author, ISBN)
	err := Execute(lib.db, []string{s})
	return lib.book_tot, err
}

// Add a student account into the library
func (lib *Library) RemoveBook(book_id int, detail string) error {
	s0 := fmt.Sprintf("INSERT INTO remove_logs(book_id, ISBN, detail) (SELECT book_id, ISBN, @detail := '%s' FROM books_avail WHERE book_id = %d)", detail, book_id)
	s1 := fmt.Sprintf("INSERT INTO remove_logs(book_id, ISBN, detail) (SELECT book_id, ISBN, @detail := '%s' FROM books_borrow WHERE book_id = %d)", detail, book_id)
	s2 := fmt.Sprintf("DELETE FROM books_avail WHERE book_id = %d", book_id)
	s3 := fmt.Sprintf("DELETE FROM books_borrow WHERE book_id = %d", book_id)
	err := Execute(lib.db, []string{s0, s1, s2, s3})
	return err
}

// Add a student account into the library
func (lib *Library) AddAccount(name, password string) (int, error) {
	lib.student_tot = lib.student_tot + 1
	s := fmt.Sprintf("INSERT INTO students(student_id, name, password) VALUES (%d, '%s', '%s')", lib.student_tot, name, password)
	err := Execute(lib.db, []string{s})
	return lib.student_tot, err
}

//Find a book given ISBN in the library
func (lib *Library) FindBook(title, author, ISBN string) ([]struct {
	book_id             int
	title, author, ISBN string
}, []struct {
	book_id             int
	title, author, ISBN string
}, error) {
	var s0, s1 string
	s0 = fmt.Sprintf("SELECT book_id, title, author, ISBN FROM books_avail WHERE title = '%s' and author = '%s' and ISBN = '%s'", title, author, ISBN)
	s1 = fmt.Sprintf("SELECT book_id, title, author, ISBN FROM books_borrow WHERE title = '%s' and author = '%s' and ISBN = '%s'", title, author, ISBN)
	rows, err := lib.db.Query(s0)
	if err != nil {
		fmt.Println(s0)
		panic(err)
	}
	var avail, borrow []struct {
		book_id             int
		title, author, ISBN string
	}
	for rows.Next() {
		var book_id int
		var title, author, ISBN string
		rows.Scan(&book_id, &title, &author, &ISBN)
		avail = append(avail, BOOK{book_id, title, author, ISBN})
	}
	rows.Close()
	rows, err = lib.db.Query(s1)
	if err != nil {
		fmt.Println(s1)
		panic(err)
	}
	for rows.Next() {
		var book_id int
		var title, author, ISBN string
		rows.Scan(&book_id, &title, &author, &ISBN)
		borrow = append(borrow, BOOK{book_id, title, author, ISBN})
	}
	return avail, borrow, nil
}

//Borrow a book given ISBN from a student account given student_id at the time date in the library
func (lib *Library) BorrowBook(book_id, student_id int, deadline string) error {
	current_time := time.Now().Format("2006-01-02")
	s := fmt.Sprintf("SELECT COUNT(*) FROM books_borrow WHERE student_id = %d AND deadline > %s", student_id, current_time)
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
	s0 := fmt.Sprintf("INSERT INTO books_borrow(book_id, title, author, ISBN, student_id, deadline, extend_times) SELECT book_id, title, author, ISBN, @student_id := %d, @deadline := '%s', @extend_times := 0 FROM books_avail WHERE books_avail.book_id = %d", student_id, deadline, book_id)
	s1 := fmt.Sprintf("DELETE FROM books_avail WHERE book_id = %d", book_id)
	Execute(lib.db, []string{s0, s1})
	return nil
}

//Query borrow history from student account given student_id
func (lib *Library) QueryHistory(student_id int) ([]struct {
	book_id int
	ISBN    string
}, error) {
	s := fmt.Sprintf("SELECT book_id, ISBN FROM borrow_logs WHERE student_id = %d", student_id)
	rows, er := lib.db.Query(s)
	if er != nil {
		fmt.Println(s)
		panic(er)
	}
	var nor []struct {
		book_id int
		ISBN    string
	}
	for rows.Next() {
		var book_id int
		var ISBN string
		rows.Scan(&book_id, &ISBN)
		nor = append(nor, BORROW{book_id, ISBN})
	}
	defer rows.Close()
	return nor, nil
}

//Query books borrowed but not returned given student account student_id
func (lib *Library) QueryBookNotReturned(student_id int) ([]struct {
	book_id             int
	title, author, ISBN string
}, error) {
	s := fmt.Sprintf("SELECT book_id, title, author, ISBN FROM books_borrow WHERE student_id = %d", student_id)
	rows, er := lib.db.Query(s)
	if er != nil {
		fmt.Println(s)
		panic(er)
	}
	var res []struct {
		book_id             int
		title, author, ISBN string
	}
	for rows.Next() {
		var book_id int
		var title, author, ISBN string
		rows.Scan(&book_id, &title, &author, &ISBN)
		res = append(res, BOOK{book_id, title, author, ISBN})
	}
	rows.Close()
	return res, nil
}

//Query deadline given a book ISBN
func (lib *Library) QueryDeadline(book_id int) ([]string, error) {
	s := fmt.Sprintf("SELECT deadline FROM books_borrow WHERE book_id = %d", book_id)
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
func (lib *Library) ExtendDeadline(book_id int, date string) error {
	s := fmt.Sprintf("UPDATE books_borrow SET deadline = '%s', extend_times = extend_times + 1 WHERE extend_times < 3 AND book_id = %d", date, book_id)
	err := Execute(lib.db, []string{s})
	return err
}

//Query the overdue books
func (lib *Library) QueryOverdue(student_id int) ([]struct {
	book_id             int
	title, author, ISBN string
}, error) {
	date := time.Now().Format("2006-01-02")
	s := fmt.Sprintf("SELECT book_id, title, author, ISBN FROM books_borrow WHERE deadline > %s AND student_id = %d", date, student_id)
	rows, er := lib.db.Query(s)
	if er != nil {
		fmt.Println(s)
		panic(er)
	}
	var res []struct {
		book_id             int
		title, author, ISBN string
	}
	for rows.Next() {
		var book_id int
		var title, author, ISBN string
		rows.Scan(&book_id, &title, &author, &ISBN)
		res = append(res, BOOK{book_id, title, author, ISBN})
	}
	defer rows.Close()
	return res, nil
}

//Return a Book
func (lib *Library) ReturnBook(book_id int) error {
	s0 := fmt.Sprintf("INSERT INTO borrow_logs(book_id, student_id, ISBN) SELECT book_id, student_id, ISBN FROM books_borrow WHERE books_borrow.book_id = %d", book_id)
	s1 := fmt.Sprintf("INSERT INTO books_avail(book_id, ISBN, author, title) SELECT book_id, ISBN, author, title FROM books_borrow WHERE books_borrow.book_id = %d", book_id)
	s2 := fmt.Sprintf("DELETE FROM books_borrow WHERE book_id = %d", book_id)
	err := Execute(lib.db, []string{s0, s1, s2})
	return err
}

func (lib *Library) DeleteAll() error {
	s0 := fmt.Sprintf("DROP TABLE remove_logs")
	s1 := fmt.Sprintf("DROP TABLE borrow_logs")
	s2 := fmt.Sprintf("DROP TABLE books_borrow")
	s3 := fmt.Sprintf("DROP TABLE students")
	s4 := fmt.Sprintf("DROP TABLE books")
	err := Execute(lib.db, []string{s0, s1, s2, s3, s4})
	lib.book_tot = 0
	lib.student_tot = 0
	return err
}

func main() {
	fmt.Println("Welcome to the Library Management System!")

	lib := new(Library)
	lib.ConnectDB()
	lib.CreateTables()

	fmt.Println("OK")
	for {
		op := ReadStr()
		op = strings.ToLower(op)
		switch op {
		case "addbook":
			inp := Board([]string{"title", "author", "ISBN"})
			book_id, err := lib.AddBook(inp[0], inp[1], inp[2])
			if err != nil {
				panic(err)
			}
			fmt.Printf("%d OK\n", book_id)
		case "removebook":
			inp := Board([]string{"bood_id", "detail"})
			book_id, err := strconv.Atoi(inp[0])
			if err != nil {
				panic(err)
			}
			err = lib.RemoveBook(book_id, inp[1])
			if err != nil {
				panic(err)
			}
			fmt.Println("OK")
		case "addaccount":
			inp := Board([]string{"name", "password"})
			student_id, err := lib.AddAccount(inp[0], inp[1])
			if err != nil {
				panic(err)
			}
			fmt.Printf("%d OK\n", student_id)
		case "findbook":
			inp := Board([]string{"title", "author", "ISBN"})
			ret, nor, err := lib.FindBook(inp[0], inp[1], inp[2])
			if err != nil {
				panic(err)
			}
			fmt.Println("Returned: book_id | title | author | ISBN")
			for _, __ := range ret {
				fmt.Printf("[ %d | %s | %s | %s ]\n", __.book_id, __.title, __.author, __.ISBN)
			}
			fmt.Println("Borrowed: book_id | title | author | ISBN")
			for _, __ := range nor {
				fmt.Printf("[ %d | %s | %s | %s ]\n", __.book_id, __.title, __.author, __.ISBN)
			}
		case "borrowbook":
			var err error
			var book_id, student_id int
			inp := Board([]string{"book_id", "student_id", "deadline"})
			book_id, err = strconv.Atoi(inp[0])
			if err != nil {
				panic(err)
			}
			student_id, err = strconv.Atoi(inp[1])
			if err != nil {
				panic(err)
			}
			err = lib.BorrowBook(book_id, student_id, inp[2])
			if err != nil {
				panic(err)
			}
			fmt.Println("OK")
		case "queryhistory":
			inp := Board([]string{"studend_id"})
			student_id, err := strconv.Atoi(inp[0])
			res, err := lib.QueryHistory(student_id)
			if err != nil {
				panic(err)
			}
			fmt.Println("  book_id | ISBN  ")
			for _, __ := range res {
				fmt.Printf("[ %d | %s ]\n", __.book_id, __.ISBN)
			}
		case "querybooknotreturned":
			inp := Board([]string{"student_id"})
			student_id, err := strconv.Atoi(inp[0])
			res, err := lib.QueryBookNotReturned(student_id)
			if err != nil {
				panic(err)
			}
			fmt.Println("  book_id | title | author | ISBN  ")
			for _, __ := range res {
				fmt.Printf("[ %d | %s | %s | %s ]\n", __.book_id, __.title, __.author, __.ISBN)
			}
		case "querydeadline":
			inp := Board([]string{"book_id"})
			book_id, err := strconv.Atoi(inp[0])
			res, err := lib.QueryDeadline(book_id)
			if err != nil {
				panic(err)
			}
			fmt.Println("  deadline  ")
			for _, __ := range res {
				fmt.Printf("[ %s ]\n", __)
			}
		case "extenddeadline":
			inp := Board([]string{})
			book_id, err := strconv.Atoi(inp[0])
			err = lib.ExtendDeadline(book_id, inp[1])
			if err != nil {
				panic(err)
			}
			fmt.Println("OK")

		case "queryoverdue":
			inp := Board([]string{"student_id"})
			student_id, err := strconv.Atoi(inp[0])
			res, err := lib.QueryOverdue(student_id)
			if err != nil {
				panic(err)
			}
			fmt.Printf("  book_id | title | author | ISBN  ")
			for _, __ := range res {
				fmt.Printf("[ %d | %s | %s | %s ]\n", __.book_id, __.title, __.author, __.ISBN)
			}

		case "returnbook":
			inp := Board([]string{"book_id"})
			book_id, err := strconv.Atoi(inp[0])
			err = lib.ReturnBook(book_id)
			if err != nil {
				panic(err)
			}
			fmt.Println("OK")

		case "deleteall":
			lib.DeleteAll()
			fmt.Println("OK")
		}

		if op == "exit" {
			fmt.Println("Bye")
			break
		}
	}

	lib.db.Close()
}
