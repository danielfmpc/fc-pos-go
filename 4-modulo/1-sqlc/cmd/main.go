package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/danielfmpc/pos-go-sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// _, err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:   uuid.New().String(),
	// 	Name: "Electronica",
	// 	Description: sql.NullString{
	// 		String: "Electronica",
	// 		Valid:  true,
	// 	},
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:   "24315b27-2452-4187-be96-444c6e6238e5",
	// 	Name: "Electronica 2",
	// 	Description: sql.NullString{
	// 		String: "Electronica 2",
	// 		Valid:  true,
	// 	},
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Categoria actualizada")

	err = queries.DeleteCategory(ctx, "24315b27-2452-4187-be96-444c6e6238e5")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Categoria eliminada")

	categories, err := queries.ListAllCategories(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, category := range categories {
		fmt.Printf("ID: %s, Name: %s, Description: %s\n", category.ID, category.Name, category.Description.String)
	}
}
