books(book_id, title, author, ISBN)
students(student_id, name, password)
books_borrow(book_id, title, author, ISBN student_id, deadline, extend_times)
borrow_logs(book_id, student_id, ISBN)
remove_logs(book_id, ISBN, detail)

0.	CREATE TABLE books_avail(
		book_id INT,
		title CHAR(32),
		author CHAR(32),
		ISBN CHAR(32),
		PRIMARY KEY(book_id)
	)
	CREATE TABLE students(
		student_id INT,
		name CHAR(32),
		password CHAR(32),
		PRIMARY KEY(student_id)
	)
	CREATE RABLE books_borrow(
		book_id INT,
		title CHAR(32),
		author CHAR(32),
		ISBN CHAR(32),
		student_id INT,
		deadline DATE,
		extend_times INT(4),
		PRIMARY KEY (book_id)
		FOREIGN KEY (student_id) REFERENCES students(student_id)
	)
	CREATE TABLE borrow_logs(
		book_id INT,
		student_id INT,
		ISBN CHAR(32),
		FOREIGN KEY (student_id) REFERENCES students(student_id)		
	)
	CREATE TABLE remove_logs(
		book_id INT
		ISBN CHAR(32)
		detail CHAR(64)
		PRIMARY KEY (book_id)
	)



1.  INSERT INTO books_avail(book_id, title, author, ISBN) VALUES (%d, '%s', '%s', '%s'), book_id, title, author, ISBN

2.  INSERT INTO remove_logs(book_id, ISBN, detail) 
		(SELECT book_id, ISBN, detail = '%s' FROM books_avail WHERE book_id = %d), detail, book_id
	INSERT INTO remove_logs(book_id, ISBN, detail) 
		(SELECT book_id, ISBN, detail = '%s' FROM books_borrow WHERE book_id = %d), detail, book_id
	DELETE FROM books_avail WHERE book_id = %d, book_id
	DELETE FROM books_borrow WHERE book_id = %d, book_id
	
3.  INSERT INTO students(student_id, name, password) VALUES (student_id, name, password)

4.  SELECT * FROM books_avail
        WHERE ISBN = '%s' or title = '%s' or author = '%s', ISBN, title, author
	SELECT * FROM books_borrow
        WHERE ISBN = '%s' or title = '%s' or author = '%s', ISBN, title, author
	
5.  SELECT COUNT (*) FROM books_borrow
        WHERE student_id = %d AND deadline > '%s', student_id, current_time 
    INSERT INTO books_borrow(books_id, title, author, ISBN, student_id, deadline, extend_times) 
		SELECT book_id, title, author, ISBN, student_id = %d, deadline = '%s', extend_times = 0 
			FROM books_avail WHERE books_avail.books_id = %d, student_id, deadline, book_id
	DELETE FROM books_avail WHERE book_id = %d, book_id

6.  SELECT (book_id, ISBN) FROM borrow_logs
        WHERE student_id = %d, student_id
    SELECT (book_id, ISBN) FROM books_borrow
        WHERE student_id = %d, student_id

7.  SELECT (book_id, title, author, ISBN) FROM books_borrow
        WHERE student_id = %d, student_id

8.  SELECT deadline FROM books_borrow
        WHERE book_id = %d, book_id
	SELECT deadline FROM books_borrow
		WHERE ISBN = '%s', ISBN

9.  UPDATE books_borrow SET deadline = '%s', extend_times = extend_times + 1
        WHERE extend_times < 3 AND book_id = %d, book_id

10. SELECT (book_id, title, author, ISBN) FROM books_borrow
        WHERE deadline > current_time AND student_id = %d, student_id

11. INSERT INTO borrow_logs(book_id, student_id, ISBN)(
        SELECT (book_id, student_id, ISBN) FROM books_borrow
            WHERE books_borrow.book_id = %d), book_id
    INSERT INTO books_avail(book_id, ISBN, author, title)(
        SELECT (ook_id, student_id, ISBN, author, title FROM books_borrow
            WHERE books_borrow.book_id = %d), book_id
	DELETE FROM books_borrow
        WHERE book_id = %d, book_id

