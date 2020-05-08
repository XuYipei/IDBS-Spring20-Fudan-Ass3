# assignment3 实验报告

## 实验描述

​	使用Golang和MySQL实现一个简易图书馆管理系统

## 数据库模式

​	共创建五个模式，所包含的属性分别如下所示：

- books_avail(book_id, title, author, ISBN)，主键为book_id，分别表示书本编号、书名、出版社名 、ISBN编码，记录在架图书.
- students(student_id, name, password)，主键为student_id，本别表示学生账号、姓名、密码.
- books_borrow(book_id, title, author, ISBN, student_id, deadline, extend-_times)，主键为book_id，分别表示书本编号、书名、出版社名 、ISBN编码，学生编号，归还期限、延期次数，记录出借图书.
- borrow_logs(book_id, student_id, ISBN)，主键为books_id，分别表示书本编号、学生账号、书本ISBN码，记录历史借阅信息.
- remove_logs(book_id, ISBN, detail)，主键为book_id，分别表示书本编号、ISBN编码、备注，记录删去的书本

## 功能及实现说明

### CreateTables

​	创建上述模式  

### AddBook	title, author, ISBN

​	增加一本书本  
​	向books_avail中添加新元组，系统会自动分配书本编号

### RemoveBook	book_id, detail

​	删除一本书，并添加备注  
​	先向remove_logs中添加相关书本，再在可能出现该本书的books_avail、books_-borrow中和该书本有关的元组

### AddAccount	name, password

​	添加一个学生账户  
​	向students中添加相关元组，系统会自动分配学生的账号名 

### FindBook	title, author, ISBN

​	查找相关书本信息  
​	书本有是否被借阅两种情况，分别在books_avail和books_borrow中查找相关条件，返回两个集合，表示被借阅和没被借阅的书本

### BorrowBook	book_id, student_id, deadline

​	某账号借阅某书本  
​	先用count得到逾期书本的数量，若逾期书本数小于3，就在books_borrow中添加相关元组，在books_avail中删除相关元组

### QueryHistory	student_id

​	查找某账户的借阅历史  
​	在borrow_logs中查找

### QueryBookNotReturned	student_id

​	查找账户未归还的书本  
​	在books_borrow中查找

### QueryDeadline	book_id

​	查找书本的归还期限  
​	在books_borrow中查找

### ExtendDeadline	book_id, date

​	书本的借阅期限延期  
​	修改extend_times小于3的书本的deadline，并令其extend_times加1

### QueryOverdue	studeng_id

​	查询账户上逾期书本  
​	在books_borros中查找deadline小于当前日期的元组

### ReturnBook	book_id

​	归还某本图书  
​	把该书的借阅记录添加到borrow_logs中，把书本添加到books_avail中，并从bo-ks_borrow中删除相关元组

### DeleteAll

​	删除数据库中所有内容（包括模式）

## 命令行交互

​	首先进行MySQL的用户登录  
​	通过输入上述功能名并按回车，进入相关功能界面  
​	按照界面提示输入相关参数，通过回车确认  
​	可在windows/linux下运行  
​	运行示例：  
​	![1](D:\DataBase\IDBS-Spring20-Fudan-Ass3\Report\1.jpg)

## Test文件

​	文件为library_test.go，使用命令go test -v进行测试

## 项目链接

​	https://github.com/XuYipei/IDBS-Spring20-Fudan-Ass3

