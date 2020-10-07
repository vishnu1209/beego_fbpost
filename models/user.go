package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// User represents a person in the system
type User struct {
	Id        int
	FirstName string
	LastName  string
}

func init() {
	orm.RegisterModel(new(User))
}

// GetAllUsers function gets all users
func GetAllUsers() (users []*User, err error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("User").All(&users)
	//fmt.Println(num, users)
	if num > 0 && err != orm.ErrNoRows {
		return users, nil
	} else {
		return nil, err
	}
}

// InsertOneUser inserts a single new User record
func InsertOneUser(user *User) (id int64, err error) {
	o := orm.NewOrm()
	fmt.Println(*user)
	id, err = o.Insert(user)
	if err != nil {
		return 0, err
	}

	//qs := o.QueryTable(new(User))
	//
	//// get prepared statement
	//i, _ := qs.PrepareInsert()
	//
	//var u User
	//fmt.Println(u.Id)

	// hash password
	//user.Password, _ = hashPassword(user.Password)

	// get now datetime
	//user.RegDate = time.Now()

	// Insert
	//user = User{Id:1, FirstName:"Vishnu"}

	//if err == nil {
	//	// successfully inserted
	//	//u = User{Id: int(id)}
	//	//err := o.Read(&u)
	//	//if err == orm.ErrNoRows {
	//	//	return nil
	//	//}
	//} else {
	//	return 0
	//}

	return
}

// UpdateUser updates an existing user
func UpdateUser(user User) *User {
	o := orm.NewOrm()
	u := User{Id: user.Id}
	var updatedUser User

	// get existing user
	if o.Read(&u) == nil {

		// updated user
		// hash new password
		//user.Password, _ = hashPassword(user.Password)

		// Keep the old registration date in case user tries to update it
		//user.RegDate = u.RegDate
		u = user
		_, err := o.Update(&u)

		// read updated user
		if err == nil {
			// update successful
			updatedUser = User{Id: user.Id}
			o.Read(&updatedUser)
		}
	}

	return &updatedUser
}

// DeleteUser deletes a user
func DeleteUser(id int) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&User{Id: id})
	if err == nil {
		// successfull
		return true
	}

	return false
}
