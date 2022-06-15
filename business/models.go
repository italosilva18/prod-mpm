package models

type UserResponse struct {
	Nome    int    `json:"status"`
	Message string `json:"message"`
}

/*
func (this *UserResponse) Add (name string, message string) (user, error) {
	var UserResponse User
	db =

	if err := db.Create(&varuser). Error; err != nil {
	return user{}, err
	} else {
	return varuser, nil
	}
	return varuser, nil
	}
	func (this *user) Update(postValues map[string][]string) error {
	var user user
	ILDB
	return nil
	}
	func (this user) Del(id uint) error {
	var user user
	if err := libs.DB.Where("id= ?", id).Delete(&user). Error; err != nil {
	return err
	}
	return nil
	}
*/
