package main

import (
	"testing"
)

func TestCreateTables(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	lib.DeleteAll()
	err := lib.CreateTables()
	if err != nil {
		t.Errorf("can't create tables")
	}
}

func TestAddBook(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	lib.AddBook("Orz Lzh", "Lzh", "0") //1
	lib.AddBook("Orz Lzh", "Lzh", "0") //2
	lib.AddBook("Orz Lzh", "Lzh", "0") //3
	lib.AddBook("Orz Lzh", "Lzh", "1") //4
	lib.AddBook("Orz Lzh", "Lzh", "1") //5

	lib.AddBook("Orz Lzh Tql", "Lin", "2") //6
	lib.AddBook("Orz Lzh Tql", "Lin", "3") //7
	lib.AddBook("Orz Lzh Tql", "Lin", "3") //8
	lib.AddBook("Orz Lzh Tql", "Lin", "3") //9

	lib.AddBook("Lzh Tql", "L", "4") //10
	lib.AddBook("Lzh Tql", "L", "4") //11
	lib.AddBook("Lzh Tql", "L", "5") //12
	lib.AddBook("Lzh Tql", "L", "6") //13
	lib.AddBook("Lzh Tql", "L", "7") //14
}

func TestRemoveBook(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	lib.RemoveBook(13, "Fire")
}

func TestAddAccount(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	lib.AddAccount("Lzh", "1") //1
	lib.AddAccount("XAs", "1") //2
	lib.AddAccount("Pry", "1") //3
	lib.AddAccount("Hld", "1") //4
	lib.AddAccount("Swt", "1") //5
	lib.AddAccount("Wb", "1")  //6
}

func CheckBook(ans, res []struct {
	book_id             int
	title, author, ISBN string
}) bool {
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

func TestFindBook(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	ans0 := []struct {
		book_id             int
		title, author, ISBN string
	}{
		{1, "Orz Lzh", "Lzh", "0"},
		{2, "Orz Lzh", "Lzh", "0"},
		{3, "Orz Lzh", "Lzh", "0"},
	}
	ans1 := []struct {
		book_id             int
		title, author, ISBN string
	}{}

	ret, nor, err := lib.FindBook("Orz Lzh", "Lzh", "0")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	if !CheckBook(ret, ans0) {
		t.Errorf("WrongAnswer")
	}
	if !CheckBook(nor, ans1) {
		t.Errorf("WrongAnswer")
	}
}

func TestBorrowBook(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	err := lib.BorrowBook(1, 3, "2020-5-5")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	err = lib.BorrowBook(2, 3, "2020-5-6")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	err = lib.BorrowBook(3, 3, "2020-5-9")
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
	res, err := lib.QueryBookNotReturned(3)
	if err != nil {
		t.Errorf("Runtime Error")
	}
	ans := []struct {
		book_id             int
		title, author, ISBN string
	}{
		{1, "Orz Lzh", "Lzh", "0"},
		{2, "Orz Lzh", "Lzh", "0"},
		{3, "Orz Lzh", "Lzh", "0"},
	}
	if !CheckBook(ans, res) {
		t.Errorf("Wrong Answer")
	}
}

func TestModifyDeadline(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	err := lib.ExtendDeadline(1, "2020-5-6")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	err = lib.ExtendDeadline(1, "2020-5-6")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	err = lib.ExtendDeadline(1, "2020-5-6")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	err = lib.ExtendDeadline(1, "2020-5-7")
	if err != nil {
		t.Errorf("Runtime Error")
	}
}

func TestQueryDeadline(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	res, err := lib.QueryDeadline(1)
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
	err := lib.ReturnBook(3)
	if err != nil {
		t.Errorf("Runtime Error")
	}
}

func TestQueryOverdue(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	res, err := lib.QueryOverdue(3)
	if err != nil {
		t.Errorf("Runtime Error")
	}
	ans := []struct {
		book_id             int
		title, author, ISBN string
	}{
		{1, "Orz Lzh", "Lzh", "0"},
		{2, "Orz Lzh", "Lzh", "0"},
	}
	if !CheckBook(ans, res) {
		t.Errorf("Wrong Answer")
	}
}

func CheckBorrow(ans, res []struct {
	book_id int
	ISBN    string
}) bool {
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

func TestBorrorwHistory(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	res, err := lib.QueryHistory(3)
	if err != nil {
		t.Errorf("Runtime Error")
	}
	ans := []struct {
		book_id int
		ISBN    string
	}{
		{3, "0"},
	}
	if !CheckBorrow(ans, res) {
		t.Errorf("Wrong Answer")
	}
}

func TestBorrowAgain(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	err := lib.BorrowBook(5, 3, "2020-5-5")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	err = lib.BorrowBook(6, 3, "2020-5-6")
	if err != nil {
		t.Errorf("Runtime Error")
	}
	nor, er := lib.QueryBookNotReturned(3)
	if er != nil {
		t.Errorf("Runtime Error")
	}
	ans := []struct {
		book_id             int
		title, author, ISBN string
	}{
		{1, "Orz Lzh", "Lzh", "0"},
		{2, "Orz Lzh", "Lzh", "0"},
		{5, "Orz Lzh", "Lzh", "1"},
		{6, "Orz Lzh Tql", "Lin", "2"},
	}
	if !CheckBook(ans, nor) {
		t.Errorf("Wrong Answer")
	}
}
