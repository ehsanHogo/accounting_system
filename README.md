# accounting_system

- This is an accounting system project, written in Go. For this project, GORM, PostgreSQL, and Go Migrate are used.

### How to run

- download the repository
- create a database (e.g. username : postgres, password : 12551255, host : localhost, port : 5432 , database : accounting)
- open terminel
- run this command based what you set for your database:

```bash
    $ export POSTGRESQL_URL='postgres://postgres:12551255@localhost:5432/accounting?sslmode=disable'
```

- for migration up :

```bash
    $ migrate -database ${POSTGRESQL_URL} -path internal/db/migrations up
```

- for migration down:

```bash
    $ migrate -database ${POSTGRESQL_URL} -path internal/db/migrations down
```
