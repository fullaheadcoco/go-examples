package repository

import (
	"errors"
	"fmt"
	"golang-transaction/model"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

// UserRepository : represent the user's repository contract
type UserRepository interface {
	Save(model.User) (model.User, error)
	GetAll() ([]model.User, error)
	IncrementMoney(uint, float64) error
	DecrementMoney(uint, float64) error
	WithTrx(*gorm.DB) userRepository
	Migrate() error
}

// NewUserRepository -> returns new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) Migrate() error {
	log.Print("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&model.User{})
}

func (u userRepository) Save(user model.User) (model.User, error) {
	log.Print("[UserRepository]...Save")
	err := u.DB.Create(&user).Error
	return user, err

}

func (u userRepository) GetAll() (users []model.User, err error) {
	log.Print("[UserRepository]...Get All")
	err = u.DB.Find(&users).Error
	return users, err

}

func (u userRepository) WithTrx(trxHandle *gorm.DB) userRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle
	return u
}

func (u userRepository) IncrementMoney(receiver uint, amount float64) error {
	log.Print("[UserRepository]...Increment Money")
	return u.DB.Model(&model.User{}).Where("id=?", receiver).Update("wallet", gorm.Expr("wallet + ?", amount)).Error
}

func (u userRepository) DecrementMoney(giver uint, amount float64) error {
	log.Print("[UserRepository]...Decrement Money")
	user := &model.User{}
	if err := u.DB.Take(user, "id=?", giver).Where("id=?", giver).Error; err != nil {
		return err
	}
	if user.Wallet-amount < 0 {
		return errors.New(fmt.Sprintf("money is insufficient. (id=%s)", giver))
	}
	return u.DB.Model(user).Update("wallet", gorm.Expr("wallet - ?", amount)).Error
}
