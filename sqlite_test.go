package sqlite

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func TestDialector(t *testing.T) {
	// This is the DSN of the in-memory SQLite database for these tests.
	const InMemoryDSN = "file:testdatabase?mode=memory&cache=shared"

	rows := []struct {
		description  string
		dialector    *Dialector
		openSuccess  bool
		query        string
		querySuccess bool
	}{
		{
			description: "Default driver",
			dialector: &Dialector{
				DSN: InMemoryDSN,
			},
			openSuccess:  true,
			query:        "SELECT 1",
			querySuccess: true,
		},
		{
			description: "Explicit default driver",
			dialector: &Dialector{
				DriverName: DriverName,
				DSN:        InMemoryDSN,
			},
			openSuccess:  true,
			query:        "SELECT 1",
			querySuccess: true,
		},
		{
			description: "Bad driver",
			dialector: &Dialector{
				DriverName: "not-a-real-driver",
				DSN:        InMemoryDSN,
			},
			openSuccess: false,
		},
		{
			description: "Explicit default driver, custom function",
			dialector: &Dialector{
				DriverName: DriverName,
				DSN:        InMemoryDSN,
			},
			openSuccess:  true,
			query:        "SELECT my_custom_function()",
			querySuccess: false,
		},
	}
	for rowIndex, row := range rows {
		t.Run(fmt.Sprintf("%d/%s", rowIndex, row.description), func(t *testing.T) {
			db, err := gorm.Open(row.dialector, &gorm.Config{})
			if !row.openSuccess {
				if err == nil {
					t.Errorf("Expected Open to fail.")
				}
				return
			}

			if err != nil {
				t.Errorf("Expected Open to succeed; got error: %v", err)
			}
			if db == nil {
				t.Errorf("Expected db to be non-nil.")
			}
			if row.query != "" {
				err = db.Exec(row.query).Error
				if !row.querySuccess {
					if err == nil {
						t.Errorf("Expected query to fail.")
					}
					return
				}

				if err != nil {
					t.Errorf("Expected query to succeed; got error: %v", err)
				}
			}
		})
	}
}
