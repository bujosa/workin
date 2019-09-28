package models

import (
	
	"Backend/Entities"
	"database/sql"
)

type CrimeModel struct {
	Db *sql.DB
}

var crimes []entities.Crime

func (crimemodel CrimeModel) FindAll() ([]entities.Crime, error) {
	rows, err := crimemodel.Db.Query("select * from Delitos")

	if err != nil {
		return nil, err
	} else {
		
		for rows.Next() {
			var id             int64 
			var cat		 	   int64
			var loc	           int64
			var date		   string
			var time		   string
			var description    string
			var mode		   string

			err2 := rows.Scan(&id, &cat, &loc, &date, &time, &description, &mode)
			if err2 != nil {
				return nil, err2
			} else {
				crime := entities.Crime{
					Id:             id,
					Cat: 			cat,
					Loc:         	loc,
					Date:     		date,
					Time:         	time,
					Description:    description,
					Mode:    		mode,
				}

				crimes = append(crimes, crime)
			}
		}
		return crimes, nil
	}

}

func (crimemodel CrimeModel) CreateCrime(CrimeN entities.Crime) (int64, error) {

	result, err := crimemodel.Db.Exec("INSERT Delitos (IdCategoria,Id_Locacion,Fecha,Hora,Descripcion,Modo) values (?, ?, ?, ?, ?, ?)", 
	CrimeN.Cat,CrimeN.Loc,CrimeN.Date,CrimeN.Time,CrimeN.Description,CrimeN.Mode)

	if err != nil {
		return 0, err
	} else {
		CrimeN.Id,_ = result.LastInsertId() 
		crimes = append(crimes, CrimeN)
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
		}
		
}
