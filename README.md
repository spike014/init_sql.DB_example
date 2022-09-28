# init sql.DB example

```
├── LICENSE
├── README.md
├── config
│   └── config.go (mock external config)
├── dao
│   ├── internal
│   │   ├── db.go
│   │   └── db_sqlmock.go
│   ├── model
│   │   └── user.go (include sql DDL)
│   ├── user.go
│   └── user_test.go
├── go.mod
├── go.sum
└── main.go

4 directories, 11 files
```

相关文章：https://spike014.github.io/posts/how-to-initialize-sql.db-in-go/
