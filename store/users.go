package store

import (
	"errors"
	"time"

	"github.com/rombintu/sanote/models"
	"github.com/rombintu/sanote/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Store) CreateUser(user models.User) error {
	ctx, err := s.Open()
	if err != nil {
		return err
	}

	var foundUser models.User
	s.Database.Collection(usersColl).FindOne(
		ctx, bson.D{{"login", user.Login}},
	).Decode(&foundUser)
	if foundUser.Login != "" {
		return errors.New(tools.ErrUserExists)
	}
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	defer s.Close(ctx)
	if _, err := s.Database.Collection(usersColl).InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUsersByLogin(login string) ([]models.User, error) {
	ctx, err := s.Open()
	if err != nil {
		return []models.User{}, err
	}
	defer s.Close(ctx)

	cur, err := s.Database.Collection(usersColl).Find(
		ctx,
		bson.D{{"login", login}},
	)
	if err != nil {
		return []models.User{}, err
	}
	defer cur.Close(ctx)

	var users []models.User
	if err := cur.All(ctx, &users); err != nil {
		return []models.User{}, err
	}
	return users, nil
}

func (s *Store) UpdateUserById(user models.User) error {
	ctx, err := s.Open()
	if err != nil {
		return err
	}
	defer s.Close(ctx)

	if _, err := s.Database.Collection(usersColl).UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		bson.D{
			{"$set", bson.D{
				{"login", user.Login},
				{"password", user.Password},
				{"updated_at", time.Now()},
			}},
		},
	); err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteUserById(id primitive.ObjectID) error {
	ctx, err := s.Open()
	if err != nil {
		return err
	}
	defer s.Close(ctx)

	if _, err := s.Database.Collection(usersColl).DeleteOne(
		ctx,
		bson.M{"_id": id},
	); err != nil {
		return err
	}
	return nil
}
