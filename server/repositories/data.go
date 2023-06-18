package repositories

import (
	"database/sql"
	"main/models"
	"strconv"
)

type data struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *data {
	return &data{db}
}

func (d data) CreateData(data models.Data) (uint64, error) {
	statement, err := d.db.Prepare("INSERT INTO data (title, user_id, description, completed) VALUES(?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	completed, err := strconv.ParseBool(data.Completed)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(data.Title, data.User_id, data.Description, completed)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (repositories data) FindAll(id int) ([]models.Data, error) {
	lines, err := repositories.db.Query("SELECT * FROM data WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer lines.Close()
	var datas []models.Data
	for lines.Next() {
		var data models.Data
		if err = lines.Scan(
			&data.ID,
			&data.User_id,
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

func (repositories data) CompleteTask(id int) (models.Data, error) {
	lines, err := repositories.db.Query("UPDATE data SET completed = true WHERE id = ?", id)
	if err != nil {
		return models.Data{}, err
	}
	defer lines.Close()

	var data models.Data
	if lines.Next() {
		if err = lines.Scan(
			&data.ID,
			&data.Title,
			&data.User_id,
			&data.Description,
			&data.Completed,
			&data.CreatedAt,
		); err != nil {
			return models.Data{}, err
		}
	}

	return data, nil

}

func (repositories data) UndoTask(id int) (models.Data, error) {
	lines, err := repositories.db.Query("UPDATE data SET completed = false WHERE id = ?", id)
	if err != nil {
		return models.Data{}, err
	}
	defer lines.Close()

	var data models.Data
	if lines.Next() {
		if err = lines.Scan(
			&data.ID,
			&data.Title,
			&data.Description,
			&data.Completed,
			&data.CreatedAt,
		); err != nil {
			return models.Data{}, err
		}
	}

	return data, nil
}

func (repositories data) DeleteTask(id int) error {
	lines, err := repositories.db.Query("DELETE FROM data WHERE id = ?", id)
	if err != nil {
		return err
	}
	defer lines.Close()

	return nil
}
