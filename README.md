# dbconfparser

Parse dbconf.yml to string.

## Install

```bash
go get github.com/mm-matsuda/dbconfparser
```

## Example

main.go
```go
package main

import (
        "github.com/jinzhu/gorm"
        _ "github.com/jinzhu/gorm/dialects/mysql"
        "github.com/mm-matsuda/dbconfparser"
)

type Sample struct {
        Name string `json:"name"`
}

func main() {
        driver, open, err := dbconfparser.Parse("db/dbconf.yml", "test")
        if err != nil {
                panic(err)
        }
        db, err := gorm.Open(driver, open)
        if err != nil {
                panic(err)
        }

        sample := Sample{}
        sample.Name = "moo"
        db.First(&sample)
}
```

dbconf.yml

```yml
development:
  driver: mysql
  open: user:password@tcp(127.0.0.1:3306)/database_development?parseTime=true
test:
  driver: mysql
  open: user:password@tcp(127.0.0.1:3306)/database_test?parseTime=true
```
