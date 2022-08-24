package repository

import (
	"context"
	"github.com/chirag3003/vote-back/config"
	"github.com/chirag3003/vote-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserAccountRepository interface {
	CreateAccount(account *models.UserAccountImportInput) error
	UpdateAccount(userName string, data *models.UserAccountImportInput) error
	FindByUsername(userName string) (*models.UserAccount, error)
}

func NewUserAccountRepository(db *mongo.Database) UserAccountRepository {
	return &userRepo{
		db: db.Collection(config.USER_ACCOUNT),
	}
}

type userRepo struct {
	db *mongo.Collection
}

func (u *userRepo) CreateAccount(data *models.UserAccountImportInput) error {
	_, err := u.db.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) UpdateAccount(userName string, data *models.UserAccountImportInput) error {
	res := &models.UserAccount{}
	err := u.db.FindOneAndUpdate(context.TODO(), bson.D{{"userName", userName}}, bson.D{{"$set", data}}).Decode(res)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) FindByUsername(userName string) (*models.UserAccount, error) {
	data := &models.UserAccount{}
	err := u.db.FindOne(context.TODO(), bson.D{{"userName", userName}}).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *userRepo) UpVote(userName string) error {
	_, err := u.db.UpdateOne(context.TODO(), bson.D{{"userName", userName}}, bson.D{{"$inc", bson.D{{"upVote", 1}}}})
	if err != nil {
		return err
	}
	return nil
}
func (u *userRepo) DownVote(userName string) error {
	_, err := u.db.UpdateOne(context.TODO(), bson.D{{"userName", userName}}, bson.D{{"$inc", bson.D{{"downVote", 1}}}})
	if err != nil {
		return err
	}
	return nil
}
