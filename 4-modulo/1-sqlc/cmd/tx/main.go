package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/danielfmpc/pos-go-sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTX(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := c.WithTx(tx)

	err = fn(q)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rollbackErr)
		}
		return err
	}

	return tx.Commit()
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, category CategoryParams, course CourseParams) error {
	err := c.callTX(ctx, func(q *db.Queries) error {
		_, err := q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
		if err != nil {
			return err
		}

		_, err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          course.ID,
			CategoryID:  category.ID,
			Name:        course.Name,
			Description: course.Description,
			Price:       course.Price,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// category := CategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Category 1",
	// 	Description: sql.NullString{String: "Description 1", Valid: true},
	// }

	// course := CourseParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Course 1",
	// 	Description: sql.NullString{String: "Description 1", Valid: true},
	// 	Price:       100,
	// }

	// err = NewCourseDB(dbConn).CreateCourseAndCategory(ctx, category, course)

	courses, err := queries.ListCourses(ctx)

	if err != nil {
		log.Fatal(err)
	}

	for _, course := range courses {
		fmt.Printf("ID: %s, Name: %s, Description: %s, Price: %f, CategoryName: %s\n", course.ID, course.Name, course.Description.String, course.Price, course.CategoryName)
	}
}
