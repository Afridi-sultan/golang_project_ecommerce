package database

type User struct{
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsShopOwner bool `json:"is_shop_owner"`
}

var users []User

func (u *User) Store (){
	if u.ID != 0{
		return 
	}
	u.ID = len(users)+1
	users = append(users, *u)
	
}

func FindUser(email, pass string)*User{
	for i := range users{
		if email == users[i].Email && pass == users[i].Password{
			return &users[i]
		}
	}
	return nil
}