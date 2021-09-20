package repository

import (
	"blog/infrastructure"
	"blog/models"
	"blog/util"
)

type UserRepository struct {
	db infrastructure.Database
}

func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (u UserRepository) CreateUser(user models.UserRegister) error {
	var dbUser models.User
	dbUser.Email = user.Email
	dbUser.Password = user.Password
	dbUser.LastName = user.LastName
	dbUser.FirstName = user.FirstName
	dbUser.IsActicve = true
	return u.db.DB.Create(&dbUser).Error
}

func (u UserRepository) LoginUser(user models.UserLogin) (*models.User, error) {
	var dbUser models.User
	email := user.Email
	password := user.Password

	err := u.db.DB.Where("email=?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}
	hashErr := util.CheckPasswordHash(password, dbUser.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	return &dbUser, nil

}
func (p UserRepository) FindAllUser(user models.User, keyword string) (*[]models.User, int64, error) {
	var users []models.User
	var totalRows int64 = 0
	queryBuilder := p.db.DB.Order("created_at desc").Model(&models.User{})
	if keyword != "" {

		querykeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			p.db.DB.Where("user.first_name LIKE ? ", querykeyword))
	}
	err := queryBuilder.Where(user).Find(&users).Count(&totalRows).Error
	return &users, totalRows, err
}
