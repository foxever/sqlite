# GORM Sqlite Driver

## USAGE

```go
import (
  "github.com/foxever/sqlite"
  "gorm.io/gorm"
)

// modernc.org/sqlite
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
```

Checkout [https://gorm.io](https://gorm.io) for details.
