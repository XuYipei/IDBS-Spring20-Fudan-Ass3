# library交互界面使用说明

## 运行方式

​	`go run library.go`

​	运行环境：Windows10 或 Linux 都可运行，go version：go1.14.1，MySQL version: MySQL 8.0.19 

## 命令输入格式

​	通过回车确认  
​	交互指令忽略大小写，数据仍区分大小写

## 登录模式

​	输入用户名、密码、以及选择的数据库名称  
​	![1](D:\DataBase\IDBS-Spring20-Fudan-Ass3\Report\1.png)

## 数据库操作

- AddBook  

  输入书本名称、书本、书本作者，会返回一个该书本的id

  ![addbook](D:\DataBase\IDBS-Spring20-Fudan-Ass3\Report\addbook.png)

- RemoveBook

  输入要删除的书本的id和删除该书本的备注

  ![removebook](D:\DataBase\IDBS-Spring20-Fudan-Ass3\Report\removebook.png)

- AddAccount

  输入账号的名称和密码

  ![addaccount](D:\DataBase\IDBS-Spring20-Fudan-Ass3\Report\addaccount.png)

- FindBook

  输入书本名称、作者和ISBN以及查询方式  
  查新方式：0-给定三个信息中的一个；1-给定两个；2-给定三个   
  不给定的信息用空串表示  
  ![findbook](D:\DataBase\IDBS-Spring20-Fudan-Ass3\Report\findbook.png)

- BorrowBook

  输入书本编号、学生编号和归还期限  
  ![borrowbook](D:\DataBase\IDBS-Spring20-Fudan-Ass3\Report\borrowbook.png)

- QueryBookNotReturned

  输入学生编号  
  ![notreturned](D:\DataBase\IDBS-Spring20-Fudan-Ass3\Report\booknotreturned.png)

- ExtendDeadline

  输入借阅书本编号、延续的时间  
  ![extenddeadline](D:\DataBase\IDBS-Spring20-Fudan-Ass3\Report\extenddeadline.png)

- QueryOverdue

  输入学生编号  
  ![queryoverdue](D:\DataBase\IDBS-Spring20-Fudan-Ass3\Report\queryoverdue.png)

- ReturnBook

  输入借阅书本的编号  
  ![returnbook](D:\DataBase\IDBS-Spring20-Fudan-Ass3\Report\returnbook.png)

- QueryHistory 

  输入学生编号  
  ![queryhistory](D:\DataBase\IDBS-Spring20-Fudan-Ass3\Report\queryhistory.png)

- DeleteAll

  删除所有数据、表  

- CreateTables

  创建所需的表

