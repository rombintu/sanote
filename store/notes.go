package store

import (
	"time"

	"github.com/rombintu/sanote/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Store) CreateNote(note models.Note) error {
	ctx, err := s.Open()
	if err != nil {
		return err
	}

	note.ID = primitive.NewObjectID()
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	defer s.Close(ctx)
	if _, err := s.Database.Collection(notesColl).InsertOne(ctx, note); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetNotesByAuthor(author string) ([]models.Note, error) {
	ctx, err := s.Open()
	if err != nil {
		return []models.Note{}, err
	}
	defer s.Close(ctx)

	cur, err := s.Database.Collection(notesColl).Find(
		ctx,
		bson.D{{"author", author}},
	)
	if err != nil {
		return []models.Note{}, err
	}
	defer cur.Close(ctx)

	var notes []models.Note
	if err := cur.All(ctx, &notes); err != nil {
		return []models.Note{}, err
	}
	return notes, nil
}

func (s *Store) UpdateNoteById(note models.Note) error {
	ctx, err := s.Open()
	if err != nil {
		return err
	}
	defer s.Close(ctx)

	if _, err := s.Database.Collection(notesColl).UpdateOne(
		ctx,
		bson.M{"_id": note.ID},
		bson.D{
			{"$set", bson.D{
				{"title", note.Title},
				{"content", note.Content},
				{"public", note.Public},
				{"tags", note.Tags},
				{"updated_at", time.Now()},
			}},
		},
	); err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteNoteById(id primitive.ObjectID) error {
	ctx, err := s.Open()
	if err != nil {
		return err
	}
	defer s.Close(ctx)

	if _, err := s.Database.Collection(notesColl).DeleteOne(
		ctx,
		bson.M{"_id": id},
	); err != nil {
		return err
	}
	return nil
}
