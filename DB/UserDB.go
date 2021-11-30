package DB

import (
	"cc_gatherer_gin/model"
)

type users struct{}

var Users users

func (t *users) Insert(user model.User) error {
	//插入数据
	stmt, err := db.Prepare(`INSERT INTO users.userinfo(u_tel,u_name,u_pass,u_email,u_words)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING u_tel`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Tel, user.Name, user.Password, user.Email, user.Words)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (t *users) WordsUpdate(tel, words string) error {
	//插入数据
	stmt, err := db.Prepare(`UPDATE users.userinfo SET u_words = $1 WHERE u_tel = $2;`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(words, tel)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (t *users) Query(tel string) (model.User, error) {
	rows, err := db.Query(`select *
	from users.userinfo
	where u_tel=$1`, tel)
	if err != nil {
		return model.User{}, err
	}
	defer rows.Close()

	var user model.User

	for rows.Next() {
		err := rows.Scan(&user.Tel, &user.Name, &user.Password, &user.Email, &user.Words)
		if err != nil {
			return model.User{}, nil
		}
	}

	err = rows.Err()
	if err != nil {
		return model.User{}, nil
	}

	return user, nil
}

func (t *users) List() ([]model.User, error) {
	rows, err := db.Query(`select *
	from users.userinfo`)
	if err != nil {
		return []model.User{}, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Tel, &user.Name, &user.Password, &user.Email)
		if err != nil {
			return []model.User{}, nil
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return []model.User{}, nil
	}

	return users, nil
}

func (t *users) Delete(tel string) error {
	//删除数据
	stmt, err := db.Prepare("DELETE FROM users.userinfo where u_tel=$1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(tel)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}
