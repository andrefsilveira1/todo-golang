package repositories

import (
	"database/sql"
	"fmt"
	"main/models"
)

type data struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *data {
	return &data{db}
}

func (d data) createData(data models.Data) (uint64, error) {
	statement, err := d.db.Prepare("INSERT INTO data (title, description, completed) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(data.Title, data.Description, data.Completed)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (repositories data) FindAll(title string) ([]models.Data, error) {
	title = fmt.Sprintf("%%%s%%", title)
	lines, err := repositories.db.Query("SELECT * FROM info")
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var datas []models.Data
	for lines.Next() {
		var data models.Data
		if err = lines.Scan(
			&data.ID,
			&data.Title,
			&data.Description,
			&data.Completed,
			&data.CreatedAt,
		); err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return datas, nil

}
