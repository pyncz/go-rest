package base

import (
	"context"
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service[
	T any,
	TFilters any,
	TCreateForm any,
	TKeyValue any,
] struct {
	Env            *models.AppEnv
	CollectionName string
	KeyField       string
}

func CreateService[
	T any,
	TFilters any,
	TCreateForm any,
	TKeyValue any,
](
	env *models.AppEnv,
	collectionName string,
	keyField string,
) *Service[T, TFilters, TCreateForm, TKeyValue] {
	return &Service[T, TFilters, TCreateForm, TKeyValue]{
		Env:            env,
		CollectionName: collectionName,
		KeyField:       keyField,
	}
}

// Methods

func (s *Service[_, _, _, _]) collection() *mongo.Collection {
	return s.Env.DB.Collection(s.CollectionName)
}

func (s *Service[T, TFilters, _, _]) Read(
	filters *TFilters,
) (*models.ListResults[T], error) {
	c := s.collection()

	var records []T

	// Get count
	count, err := c.CountDocuments(context.TODO(), filters)
	if err != nil {
		return nil, err
	}

	// Build options
	opts := options.Find()

	// Extract results
	cursor, err := c.Find(context.TODO(), filters, opts)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &records); err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Form results
	return &models.ListResults[T]{
		Count:   count,
		Results: records,
	}, nil
}

func (s *Service[T, TFilters, _, _]) ReadPaginated(
	filters *TFilters,
	pagination *models.PaginationQuery,
) (*models.PaginatedListResults[T], error) {
	c := s.collection()

	var records []T

	// Get count
	count, err := c.CountDocuments(context.TODO(), filters)
	if err != nil {
		return nil, err
	}

	// Build options
	opts := options.Find()
	opts = opts.SetLimit(pagination.Limit).SetSkip(pagination.Offset)

	// Extract results
	cursor, err := c.Find(context.TODO(), filters, opts)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &records); err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Form results
	return &models.PaginatedListResults[T]{
		Count:   count,
		Limit:   pagination.Limit,
		Offset:  pagination.Offset,
		Cursor:  utils.GetNextOffset(pagination.Offset, count, int64(len(records))),
		Results: records,
	}, nil
}

func (s *Service[T, _, TCreateForm, _]) Create(
	form *TCreateForm,
) (*T, error) {
	c := s.collection()

	inserted, err := c.InsertOne(context.TODO(), form)
	if err != nil {
		return nil, err
	}

	return findByKey[T](
		c,
		inserted.InsertedID,
		"_id",
	)
}

func findByKey[T any, TKeyValue any](
	c *mongo.Collection,
	key TKeyValue,
	keyField string,
) (*T, error) {
	// Default key field's name
	if keyField == "" {
		keyField = "_id"
	}

	// Form Find options
	filters := bson.M{
		(keyField): key,
	}

	var found T
	err := c.FindOne(context.TODO(), &filters).Decode(&found)
	if err != nil {
		return nil, err
	}

	return &found, nil
}

func (s *Service[T, TFilters, _, TKeyValue]) FindByKey(
	key TKeyValue,
) (*T, error) {
	c := s.collection()

	return findByKey[T, TKeyValue](
		c,
		key,
		s.KeyField,
	)
}
