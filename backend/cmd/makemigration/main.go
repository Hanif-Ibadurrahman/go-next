package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const migrationTemplate = `package migrations

import (
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, Migration{
		Name: "{{.Timestamp}}_{{.Name}}",
		Up: func(db *gorm.DB) error {
			err := db.Exec(` + "`" + `
				CREATE TABLE IF NOT EXISTS {{.TableName}} (
					id BIGSERIAL PRIMARY KEY,
					{{range .Fields}}{{.Name}} {{.Type}} {{if .NotNull}}NOT NULL{{end}}{{if .Unique}} UNIQUE{{end}},
					{{end}}created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					deleted_at TIMESTAMP
				)` + "`" + `).Error
			if err != nil {
				return err
			}
			err = db.Exec(` + "`" + `
				CREATE INDEX IF NOT EXISTS idx_{{.TableName}}_deleted_at ON {{.TableName}} (deleted_at)
			` + "`" + `).Error
			if err != nil {
				return err
			}
			return nil
		},
		Down: func(db *gorm.DB) error {
			err := db.Exec("DROP TABLE IF EXISTS {{.TableName}}").Error
			if err != nil {
				return err
			}
			return nil
		},
	})
}
`

type Field struct {
	Name    string
	Type    string
	NotNull bool
	Unique  bool
}

type MigrationData struct {
	Timestamp string
	Name      string
	TableName string
	Fields    []Field
}

func main() {
	name := flag.String("name", "", "Migration name (e.g., create_users_table)")
	fields := flag.String("fields", "", "Comma-separated fields (e.g., name:varchar(255):notnull:unique,email:varchar(255):notnull)")
	flag.Parse()

	if *name == "" {
		fmt.Println("Error: Migration name is required. Use -name=create_xxx_table")
		os.Exit(1)
	}

	timestamp := time.Now().UTC().Format("20060102150405")
	tableName := strings.TrimSuffix(strings.TrimPrefix(*name, "create_"), "_table")
	fileName := fmt.Sprintf("%s_%s.go", timestamp, *name)

	migrationDir := "../../migrations"
	if err := os.MkdirAll(migrationDir, 0755); err != nil {
		fmt.Printf("Error creating migrations directory: %v\n", err)
		os.Exit(1)
	}

	filePath := filepath.Join(migrationDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating migration file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var fieldList []Field
	if *fields != "" {
		for _, f := range strings.Split(*fields, ",") {
			parts := strings.Split(f, ":")
			if len(parts) < 2 {
				continue
			}
			field := Field{Name: parts[0], Type: parts[1]}
			for _, opt := range parts[2:] {
				if opt == "notnull" {
					field.NotNull = true
				}
				if opt == "unique" {
					field.Unique = true
				}
			}
			fieldList = append(fieldList, field)
		}
	}

	tmpl, err := template.New("migration").Parse(migrationTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		os.Exit(1)
	}

	data := MigrationData{
		Timestamp: timestamp,
		Name:      *name,
		TableName: tableName,
		Fields:    fieldList,
	}
	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error writing migration file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Migration created: %s\n", filePath)
}
