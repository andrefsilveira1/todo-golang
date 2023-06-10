package repositories

import (
	"main/models"
)

func (d data) CreateUser(user models.User) (uint64, error) {
	statement, err := d.db.Prepare("INSERT INTO users (name, email, password) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (d data) Login(user models.User) (models.User, error) {
	statement, err := d.db.Query("SELECT id,password FROM users WHERE email = ?", user.Email)
	if err != nil {
		return models.User{}, err
	}
	defer statement.Close()

	var localuser models.User
	if statement.Next() {
		if err = statement.Scan(&localuser.Id, &localuser.Password); err != nil {
			return models.User{}, err
		}
	}

	return localuser, nil
}
