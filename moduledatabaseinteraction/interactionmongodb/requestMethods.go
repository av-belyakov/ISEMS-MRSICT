package interactionmongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//QueryProcessor интерфейс обработчик запросов к БД
type QueryProcessor interface {
	InsertData([]interface{}) (bool, error)
	UpdateOne(interface{}, interface{}) error
	DeleteOneData(interface{}) error
	DeleteManyData([]interface{}) error
	Find(interface{}) (*mongo.Cursor, error)
	FindAlltoCollection() (*mongo.Cursor, error)
}

//QueryParameters параметры для работы с коллекциями БД
type QueryParameters struct {
	NameDB, CollectionName string
	ConnectDB              *mongo.Client
}

//InsertData добавляет все данные
func (qp *QueryParameters) InsertData(list []interface{}) (bool, error) {
	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)

	if _, err := collection.InsertMany(context.TODO(), list); err != nil {
		return false, err
	}

	if qp.CollectionName != "stix_object_collection" {
		return true, nil
	}

	if _, err := collection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "commonpropertiesobjectstix.type", Value: 1},
				{Key: "commonpropertiesobjectstix.id", Value: 1},
			},
			Options: &options.IndexOptions{},
		}, {
			Keys: bson.D{
				{Key: "source_ref", Value: 1},
			},
			Options: &options.IndexOptions{},
		},
	}); err != nil {
		return false, err
	}

	return true, nil
}

//DeleteOneData удаляет элемент
func (qp *QueryParameters) DeleteOneData(elem interface{}) error {
	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	if _, err := collection.DeleteOne(context.TODO(), elem); err != nil {
		return err
	}

	return nil
}

//DeleteManyData удаляет группу элементов
func (qp *QueryParameters) DeleteManyData(list interface{}) (*mongo.DeleteResult, error) {
	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)

	return collection.DeleteMany(context.TODO(), list)
}

//UpdateOne обновляет параметры в элементе
func (qp *QueryParameters) UpdateOne(searchElem, update interface{}) error {
	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	if _, err := collection.UpdateOne(context.TODO(), searchElem, update); err != nil {
		return err
	}

	return nil
}

//UpdateMany обновляет множественные параметры в элементе
func (qp *QueryParameters) UpdateMany(searchElem, update []interface{}) error {
	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	if _, err := collection.UpdateMany(context.TODO(), searchElem, update); err != nil {
		return err
	}

	return nil
}

//UpdateOneArrayFilters обновляет множественные параметры в массиве
func (qp *QueryParameters) UpdateOneArrayFilters(filter, update interface{}, uo *options.UpdateOptions) error {
	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	if _, err := collection.UpdateOne(context.TODO(), filter, update, uo); err != nil {
		return err
	}

	return nil
}

//Find найти всю информацию по заданному элементу
func (qp QueryParameters) Find(elem interface{}) (*mongo.Cursor, error) {
	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	options := options.Find().SetAllowDiskUse(true)
	//options := options.Find().SetAllowDiskUse(true).SetSort(bson.D{{Key: "detailed_information_on_filtering.time_interval_task_execution.start", Value: -1}})

	return collection.Find(context.TODO(), elem, options)
}

//FindOne найти информацию по заданному элементу
func (qp QueryParameters) FindOne(elem interface{}) *mongo.SingleResult {
	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	options := options.FindOne()
	//options := options.FindOne().SetSort(bson.D{{Key: "detailed_information_on_filtering.time_interval_task_execution.start", Value: -1}})

	return collection.FindOne(context.TODO(), elem, options)
}

//FindAllWithLimitOptions содержит опции поиска для метода FindAllWithLimit
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

//FindAllWithLimit найти всю информацию по заданным параметрам, но вывести ограниченное количество найденных документов
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

	if opt.SortField != "" {
		sortField = opt.SortField
	}

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

	options := options.Find().SetAllowDiskUse(true).SetSort(bson.D{{Key: sortField, Value: sortOrder}}).SetSkip(offset).SetLimit(opt.LimitMaxSize)

	return collection.Find(context.TODO(), elem, options)
}

//FindAlltoCollection найти всю информацию в коллекции
func (qp QueryParameters) FindAlltoCollection() (*mongo.Cursor, error) {
	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	options := options.Find().SetAllowDiskUse(true).SetSort(bson.D{{Key: "_id", Value: -1}})

	return collection.Find(context.TODO(), bson.D{{}}, options)
}

//CountDocuments подсчитать количество документов в коллекции
func (qp QueryParameters) CountDocuments(filter interface{}) (int64, error) {
	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
	options := options.Count()

	return collection.CountDocuments(context.TODO(), filter, options)
}

//Indexes возвращает представление индекса для этой коллекции
func (qp QueryParameters) Indexes() mongo.IndexView {
	collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)

	return collection.Indexes()
}
