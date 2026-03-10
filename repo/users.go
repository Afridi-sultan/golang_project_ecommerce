package repo

type User struct{
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsShopOwner bool `json:"is_shop_owner"`
}

type UserInterface interface{
	Create(usr *User)(*User,error)
}

type userListSruct struct{
	userList []*User
}

//constructor 
func NewUserList()UserInterface{
	return &userListSruct{}
}

func (u *userListSruct) Create(usr *User)(*User,error){
	usr.ID = len(u.userList)+1
	u.userList = append(u.userList, usr)
	return usr,nil
}