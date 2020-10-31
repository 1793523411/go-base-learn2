## docker

 `docker run --name go-mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 -d daocloud.io/mysql`

 `docker exec -it go-mysql bash`

 docker里mysql没有vim

 `apt-get update` +  `apt-get install vim`


 ## MySQL预处理
什么是预处理？

普通SQL语句执行过程：

+ 客户端对SQL语句进行占位符替换得到完整的SQL语句。
+ 客户端发送完整SQL语句到MySQL服务端
+  MySQL服务端执行完整的SQL语句并将结果返回给客户端。

预处理执行过程：
+ 把SQL语句分成两部分，命令部分与数据部分。
+  先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
+   然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
+  MySQL服务端执行完整的SQL语句并将结果返回给客户端。

为什么要预处理？

+ 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
    避免SQL注入问题。


## SQL注入问题

我们任何时候都不应该自己拼接SQL语句！

这里我们演示一个自行拼接SQL语句的示例，编写一个根据name字段查询user表的函数如下：

```go
// sql注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}
```

此时以下输入字符串都可以引发SQL注入问题：

```go
sqlInjectDemo("xxx' or 1=1#")
sqlInjectDemo("xxx' union select * from user #")
sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
```

补充：不同的数据库中，SQL语句使用的占位符语法不尽相同。
数据库 	占位符语法
MySQL 	?
PostgreSQL 	$1, $2等
SQLite 	? 和$1
Oracle 	:name



## Go实现MySQL事务
什么是事务？

事务：一个最小的不可再分的工作单元；通常一个事务对应一个完整的业务(例如银行账户转账业务，该业务就是一个最小的工作单元)，同时这个完整的业务需要执行多次的DML(insert、update、delete)语句共同联合完成。A转账给B，这里面就需要执行两次update操作。

在MySQL中**只有使用了Innodb数据库引擎的数据库或表才支持事务**。事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么全部不执行。
事务的ACID

通常事务必须满足4个条件（ACID）：**原子性（Atomicity，或称不可分割性）、一致性（Consistency）、隔离性（Isolation，又称独立性）、持久性（Durability）**。
条件 	解释
原子性 	一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。
一致性 	在事务开始之前和事务结束以后，数据库的完整性没有被破坏。这表示写入的资料必须完全符合所有的预设规则，这包含资料的精确度、串联性以及后续数据库可以自发性地完成预定的工作。
隔离性 	数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。事务隔离分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（Serializable）。
持久性 	事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失。


### 事务相关方法

Go语言中使用以下三个方法实现MySQL中的事务操作。 开始事务

`func (db *DB) Begin() (*Tx, error)`

提交事务

`func (tx *Tx) Commit() error`

回滚事务

`func (tx *Tx) Rollback() error`
