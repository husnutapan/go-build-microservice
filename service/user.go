package service

import (
	"github.com/husnutapan/go-build-microservice/pojo"
	"github.com/jinzhu/gorm"
)

type IUser interface {
	SaveUser(db *gorm.DB) (*pojo.User, error)
	FetchAll(db *gorm.DB) (*[]pojo.User, error)
	FindById(db *gorm.DB, uid int) (*pojo.User, error)
	DeleteById(db *gorm.DB, uid int) (int, error)
	UpdateUser(db *gorm.DB, uid int) (*pojo.User, error)
}

type User struct {
}

func (user User) SaveUser(u *pojo.User) (*pojo.User, error) {
	var err error
	//err = db.Debug().Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (user User) FetchAll(db *gorm.DB) (*[]pojo.User, error) {
	users := []pojo.User{}
	err := db.Debug().Model(&User{}).Limit(20).Find(&users).Error
	if err != nil {
		return &[]pojo.User{}, err
	}
	return &users, err
	return nil, nil
}

//todo will impl.
func (user User) FindById(db *gorm.DB, uid int) (*pojo.User, error) {
	return nil, nil
}

//todo will impl.
func (user User) DeleteById(db *gorm.DB, uid int) (int, error) {
	return 1, nil
}

//todo will impl.
func (user User) UpdateUser(db *gorm.DB, uid int) (*pojo.User, error) {
	return nil, nil
}
