package infrastructure

import (
	"fmt"
	"food-app/domain/entity"
	"food-app/domain/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"


)

type Repositories struct {
	User repository.UserRepository
	Food    repository.FoodRepository
	db      *gorm.DB
}

func NewServices(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		User:    NewUserRepository(db),
		Food:    NewFoodRepository(db),
		db:      db,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

//Drops all tables and rebuild them
//func (s *Services) DestructiveReset() error {
//	if err := s.db.DropTableIfExists(&User{}, &Gallery{}).Error; err != nil {
//		return err
//	}
//	return s.Automigrate()
//}

//This migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.User{}, &entity.Food{}).Error
}