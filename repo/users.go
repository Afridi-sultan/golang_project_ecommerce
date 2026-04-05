package repo

import (
	"database/sql"
	"eccomerce/domain"
	"eccomerce/user"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserInterface interface {
	user.UserInterface
}

type userListSruct struct {
	db *sqlx.DB
}

// constructor
func NewUserList(db *sqlx.DB) UserInterface {
	return &userListSruct{
		db: db,
	}
}

func (u *userListSruct) Create(usr *domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (
		first_name,
		email, 
		password, 
		is_shop_owner
		)

		VALUES (
		:first_name, 
		:email, 
		:password, 
		:is_shop_owner
		)
		RETURNING id
	`

	var userId int
	rows, err := u.db.NamedQuery(query, usr)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&userId)
	}
	usr.ID = userId
	return usr, nil
}

func (u *userListSruct) GetUser(email, pass string) (*domain.User, error) {
	query := `
		SELECT id, first_name, email, password, is_shop_owner
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1
	`

	var user domain.User
	err := u.db.Get(&user, query, email, pass)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// func (u *userListSruct) List() ([]*User, error) {
// 	return u.userList, nil
// }

// func (u *userListSruct) Delete(id int) error {
// 	var temp_list []*User
// 	for i := range u.userList {
// 		if id != u.userList[i].ID {
// 			temp_list = append(temp_list, u.userList[i])
// 		}
// 	}
// 	u.userList = temp_list
// 	return nil
// }

// func (u *userListSruct) Update(usr User) (*User, error) {
// 	for i := range u.userList {
// 		if usr.ID == u.userList[i].ID {
// 			u.userList[i] = &usr
// 			return u.userList[i], nil
// 		}
// 	}
// 	return nil, nil
// }
