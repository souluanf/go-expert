package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/souluanf/fullcycle-sqlc/internal/db"
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

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, course CourseParams, category CategoryParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryID:  category.ID,
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

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	queries := db.New(dbConn)
	//courseArgs := CourseParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Go",
	//	Description: sql.NullString{String: "Golang course", Valid: true},
	//	Price:       12.99,
	//}
	//categoryArgs := CategoryParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Backend",
	//	Description: sql.NullString{String: "Backend category", Valid: true},
	//}
	//
	//courseDB := NewCourseDB(dbConn)
	//
	//err = courseDB.CreateCourseAndCategory(ctx, courseArgs, categoryArgs)
	//if err != nil {
	//	panic(err)
	//}

	course, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}
	for _, c := range course {
		fmt.Printf("Name: %s, Description: %s, Price: %f, Category: %s\n", c.Name, c.Description.String, c.Price, c.CategoryName)
	}

}
