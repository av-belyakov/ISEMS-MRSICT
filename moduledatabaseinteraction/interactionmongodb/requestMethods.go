package interactionmongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// QueryProcessor интерфейс обработчик запросов к БД
type QueryProcessor interface {
	InsertData([]interface{}) (bool, error)
	UpdateOne(interface{}, interface{}) error
	DeleteOneData(interface{}) error
	DeleteManyData([]interface{}) error
	Find(interface{}) (*mongo.Cursor, error)
	FindAlltoCollection() (*mongo.Cursor, error)
}

// QueryParameters параметры для работы с коллекциями БД
type QueryParameters struct {
	NameDB, CollectionName string
	ConnectDB              *mongo.Client
}

// InsertData добавляет все данные
func (qp *QueryParameters) InsertData(list []interface{}, indexList []mongo.IndexModel) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)

	if _, err := collection.InsertMany(ctx, list); err != nil {
		return false, err
	}

	if qp.CollectionName != "stix_object_collection" {
		return true, nil
	}

	if _, err := collection.Indexes().CreateMany(ctx, indexList); err != nil {
		return false, err
	}

	return true, nil
}

// DeleteOneData удаляет элемент
func (qp *QueryParameters) DeleteOneData(elem interface{}, options *options.DeleteOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	if _, err := collection.DeleteOne(ctx, elem); err != nil {
		return err
	}

	return nil
}

// DeleteManyData удаляет группу элементов
func (qp *QueryParameters) DeleteManyData(list interface{}, opts *options.DeleteOptions) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)

	return collection.DeleteMany(ctx, list, opts)
}

// UpdateOne обновляет параметры в элементе
func (qp *QueryParameters) UpdateOne(searchElem, update interface{}, opts *options.UpdateOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	if _, err := collection.UpdateOne(ctx, searchElem, update, opts); err != nil {
		return err
	}

	return nil
}

// UpdateMany обновляет множественные параметры в элементе
func (qp *QueryParameters) UpdateMany(searchElem, update []interface{}, opts *options.UpdateOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	if _, err := collection.UpdateMany(ctx, searchElem, update, opts); err != nil {
		return err
	}

	return nil
}

// UpdateOneArrayFilters обновляет множественные параметры в массиве
func (qp *QueryParameters) UpdateOneArrayFilters(filter, update interface{}, uo *options.UpdateOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	if _, err := collection.UpdateOne(ctx, filter, update, uo); err != nil {
		return err
	}

	return nil
}

// Find найти всю информацию по заданному элементу
func (qp QueryParameters) Find(elem interface{}) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	options := options.Find().SetAllowDiskUse(true)
	//options := options.Find().SetAllowDiskUse(true).SetSort(bson.D{{Key: "detailed_information_on_filtering.time_interval_task_execution.start", Value: -1}})

	return collection.Find(ctx, elem, options)
}

// FindOne найти информацию по заданному элементу
func (qp QueryParameters) FindOne(elem interface{}) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	options := options.FindOne()
	//options := options.FindOne().SetSort(bson.D{{Key: "detailed_information_on_filtering.time_interval_task_execution.start", Value: -1}})

	return collection.FindOne(ctx, elem, options)
}

// FindAllWithLimitOptions содержит опции поиска для метода FindAllWithLimit
// Offset - смещение в колличестве найденных документов
// LimitMaxSize - максимальное количество возвращаемых документов
// SortField - поле по которому выполняется сортировка (по умолчанию ObjectId)
// SortAscending - порядок сортировки (по умолчанию 'сортировка по убыванию')
type FindAllWithLimitOptions struct {
	Offset        int64
	LimitMaxSize  int64
	SortField     string
	SortAscending bool
}

// FindAllWithLimit найти всю информацию по заданным параметрам, но вывести ограниченное количество найденных документов
func (qp QueryParameters) FindAllWithLimit(elem interface{}, opt *FindAllWithLimitOptions) (*mongo.Cursor, error) {
	const (
		sortAscending  int = 1
		sortDescending int = -1
	)

	var (
		offset    int64
		sortField string = "_id"
		sortOrder int    = sortDescending
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)

	if opt.SortField != "" {
		sortField = opt.SortField
	}

	if opt.SortAscending {
		sortOrder = sortAscending
	}

	if opt.Offset > 0 {
		offset = (opt.Offset - 1) * opt.LimitMaxSize
	}

	options := options.Find().SetAllowDiskUse(true).SetSort(bson.D{{Key: sortField, Value: sortOrder}, {Key: "commonpropertiesobjectstix.id", Value: sortOrder}}).SetSkip(offset).SetLimit(opt.LimitMaxSize)

	//return collection.Find(context.TODO(), elem, options)
	return collection.Find(ctx, elem, options)
}

// FindAlltoCollection найти всю информацию в коллекции
func (qp QueryParameters) FindAlltoCollection() (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	options := options.Find().SetAllowDiskUse(true).SetSort(bson.D{{Key: "_id", Value: -1}})

	return collection.Find(ctx, bson.D{{}}, options)
}

// CountDocuments подсчитать количество документов в коллекции
func (qp QueryParameters) CountDocuments(filter interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	options := options.Count()

	return collection.CountDocuments(ctx, filter, options)
}

// Indexes возвращает представление индекса для этой коллекции
func (qp QueryParameters) Indexes() mongo.IndexView {
	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)

	return collection.Indexes()
}
