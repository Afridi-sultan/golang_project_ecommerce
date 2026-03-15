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
	Get(email,pass string)(*User, error)
	List()([]*User,error)
	Delete(id int) error
	Update(usr User)(*User, error)
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

func (u *userListSruct)Get(email,pass string)(*User, error){
	for i := range u.userList{
		if email == u.userList[i].Email && pass == u.userList[i].Password{
			return u.userList[i],nil
		}
	}
	return nil,nil
}

func (u *userListSruct)List()([]*User,error){
	return u.userList,nil
}

func (u *userListSruct)Delete(id int)  error{
	var temp_list []*User
	for i := range u.userList{
		if id != u.userList[i].ID{
			temp_list = append(temp_list, u.userList[i])
		}
	}
	u.userList = temp_list
	return nil
}

func (u *userListSruct)Update(usr User)(*User, error){
	for i := range u.userList{
		if usr.ID == u.userList[i].ID{
			u.userList[i] = &usr
			return u.userList[i],nil
		}
	}
	return nil,nil
}