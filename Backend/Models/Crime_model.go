package models

import (
	
	"Backend/Entities"
	"database/sql"
	"log"
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

	
	var idloc int64

	result2, err2 := crimemodel.Db.Query("select Top(1) IdLocacion from Locaciones order by IdLocacion desc")

	if err2 != nil {
		return 0, err2
	} else {
		
		for result2.Next() {
			var id   int64 
	
			err2 = result2.Scan(&id)
			if err2 != nil {
				log.Fatal("paso por aqui")
				return 0, err2
			} else {
				
				idloc = id
			}
		}
	
	}
	
	result, err := crimemodel.Db.Exec("INSERT Delitos (IdCategoria,Id_Locacion,Fecha,Hora,Descripcion,Modo) values (?, ?, ?, ?, ?, ?)", 
	CrimeN.Cat,idloc,CrimeN.Date,CrimeN.Time,CrimeN.Description,CrimeN.Mode)

	CrimeN.Loc = idloc;
	
	if err != nil {
		log.Fatal("Llego hasta aca")
		return 0, err
	} else {
		CrimeN.Id,_ = result.LastInsertId() 
		crimes = append(crimes, CrimeN)
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
		}
		
}
