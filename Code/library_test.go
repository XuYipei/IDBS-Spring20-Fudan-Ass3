package main

import (
	"testing"
)

func TestCreateTables(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	err := lib.CreateTables()
	if err != nil {
		t.Errorf("can't create tables")
	}
}
func TestAddBook(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	lib.AddBook("Orz Lzh", "Lzh", "0")
	lib.AddBook("Orz Lzh", "Lzh", "1")
	lib.AddBook("Orz Lzh", "Lzh", "2")
	lib.AddBook("Orz Lzh", "Lzh", "3")
	lib.AddBook("Orz Lzh", "Lzh", "4")

	lib.AddBook("Orz Lzh Tql", "Lin", "5")
	lib.AddBook("Orz Lzh Tql", "Lin", "6")
	lib.AddBook("Orz Lzh Tql", "Lin", "7")
	lib.AddBook("Orz Lzh Tql", "Lin", "8")

	lib.AddBook("Lzh Tql", "L", "9")
	lib.AddBook("Lzh Tql", "L", "10")
	lib.AddBook("Lzh Tql", "L", "11")
	lib.AddBook("Lzh Tql", "L", "12")
	lib.AddBook("Lzh Tql", "L", "13")
}

func TestRemoveBook(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	lib.RemoveBook("13")
}

func TestAddAccount(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	lib.AddAccount("Lzh")
	lib.AddAccount("XAs")
	lib.AddAccount("Pry")
	lib.AddAccount("Hld")
	lib.AddAccount("Swt")
	lib.AddAccount("Wb")
}

func CheckBook(ans, res []struct{ title, author, ISBN string }) bool {
	if len(ans) != len(res) {
		return (false)
	}
	for i := 0; i < len(ans); i++ {
		find := 0
		for j := 0; j < len(res); j++ {
			if ans[i] == res[j] {
				find = 1
			}
		}
		if find == 0 {
			return (false)
		}
	}
	return (true)
}

func TestFindBookISBN(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	ans := []struct{ title, author, ISBN string }{
		{"Orz Lzh", "Lzh", "1"},
	}
	res, err := lib.FindBookISBN("1")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	if !CheckBook(ans, res) {
		t.Errorf("WrongAnswer")
	}
}

func TestFindBooktitle(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	ans := []struct{ title, author, ISBN string }{
		{"Orz Lzh Tql", "Lin", "5"},
		{"Orz Lzh Tql", "Lin", "6"},
		{"Orz Lzh Tql", "Lin", "7"},
		{"Orz Lzh Tql", "Lin", "8"},
	}
	res, err := lib.FindBookTitle("Orz Lzh Tql")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	if !CheckBook(ans, res) {
		t.Errorf("WrongAnswer")
	}
}
func TestFindBookAuthor(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	ans := []struct{ title, author, ISBN string }{
		{"Lzh Tql", "L", "9"},
		{"Lzh Tql", "L", "10"},
		{"Lzh Tql", "L", "11"},
		{"Lzh Tql", "L", "12"},
	}
	res, err := lib.FindBookAuthor("L")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	if !CheckBook(ans, res) {
		t.Errorf("WrongAnswer")
	}
}

func TestBorrowBook(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	err := lib.BorrowBook("0", "XAs", "2020-5-5", "2020-5-2")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	err = lib.BorrowBook("1", "XAs", "2020-5-6", "2020-5-2")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	err = lib.BorrowBook("2", "XAs", "2020-5-7", "2020-5-2")
	if err != nil {
		t.Errorf("Runtime Error")
	}
}

func CheckStr(ans, res []string) bool {
	if len(ans) != len(res) {
		return (false)
	}
	for i := 0; i < len(ans); i++ {
		find := 0
		for j := 0; j < len(res); j++ {
			if ans[i] == res[j] {
				find = 1
			}
		}
		if find == 0 {
			return (false)
		}
	}
	return (true)
}

func TestBookNotReturned(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	res, err := lib.QueryBookNotReturned("XAs")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	ans := []string{"0", "1", "2"}
	if !CheckStr(ans, res) {
		t.Errorf("Wrong Answer")
	}
}

func TestModifyDeadline(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	err := lib.ExtendDeadline("0", "2020-5-6")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	err = lib.ExtendDeadline("0", "2020-5-6")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	err = lib.ExtendDeadline("0", "2020-5-6")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	err = lib.ExtendDeadline("0", "2020-5-7")
	if err != nil {
		t.Errorf("Runtime Error")
	}
}

func TestQueryDeadline(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	res, err := lib.QueryDeadline("0")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	ans := []string{"2020-05-06"}
	if !CheckStr(ans, res) {
		t.Errorf("Wrong Answer")
	}
}

func TestReturnBook(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	err := lib.ReturnBook("XAs", "0")
	if err != nil {
		t.Errorf("Runtime Error")
	}
}

func TestQueryOverdue(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	res, err := lib.QueryOverdue("XAs", "2020-5-5")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	ans := []string{"1", "2"}
	if !CheckStr(ans, res) {
		t.Errorf("Wrong Answer")
	}
}

func TestBorrorwHistory(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	res, err := lib.QueryHistory("XAs")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	ans := []string{"0", "1", "2"}
	if !CheckStr(ans, res) {
		t.Errorf("Wrong Answer")
	}
}
