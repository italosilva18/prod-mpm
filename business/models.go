package models




var VarUser *user

type user struct {
gorm.Model	
}

func (this *user) Add (name string, message string) (user, error) {
	var UserResponse User
	db = repository.setup

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
