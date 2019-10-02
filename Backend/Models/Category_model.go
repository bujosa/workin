package models

import (
	entities "Backend/Entities"
	"database/sql"
)

type CategoryModel struct {
	Db *sql.DB
}

var categories []entities.Category

func (categorymodel CategoryModel) FindAll() ([]entities.Category, error) {
	rows, err := categorymodel.Db.Query("select * from Categorias")

	if err != nil {
		return nil, err
	} else {

		for rows.Next() {
			var id int64
			var categorie string

			err2 := rows.Scan(&id, &categorie)
			if err2 != nil {
				return nil, err2
			} else {
				category := entities.Category{
					Id:        id,
					Categorie: categorie,
				}

				categories = append(categories, category)
			}
		}

		return categories, nil
	}

}

func (categorymodel CategoryModel) CreateCategory(category entities.Category) (int64, error) {

	result, err := categorymodel.Db.Exec("INSERT Categorias (Categoria) values (?)", category.Categorie)
	if err != nil {
		return 0, err
	} else {
		category.Id, _ = result.LastInsertId()
		categories = append(categories, category)
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
	}

}

func (categorymodel CategoryModel) DeleteCategory(category entities.Category) (int64, error) {

	result, err := categorymodel.Db.Exec("delete from Categorias where Categoria = ?", category.Categorie)
	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
	}

}

func (categorymodel CategoryModel) UpdateCategory(category entities.Category) (int64, error) {

	result, err := categorymodel.Db.Exec("Update Categorias set Categoria = ? where IdCategoria = ?", category.Categorie, category.Id)
	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
	}

}
