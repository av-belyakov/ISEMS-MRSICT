package interactionmongodb

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"

	"go.mongodb.org/mongo-driver/bson"
)

//CreateSearchQueriesSTIXObject обработчик формирующий поисковые запросы для осуществления поиска по коллекции содержащей документы в формате STIX
func CreateSearchQueriesSTIXObject(sp *datamodels.SearchThroughCollectionSTIXObjectsType) bson.D {
	fmt.Println("func 'CreateSearchQueriesSTIXObject', START...")

	var (
		documentsType           bson.E
		dataTimeActionForObject bson.E
		createdByRef            bson.E
		specificSearchFields    bson.E
	)

	if len(sp.DocumentsID) > 0 {
		return bson.D{{Key: "commonpropertiesobjectstix.id", Value: bson.D{{Key: "$in", Value: sp.DocumentsID}}}}
	}

	if len(sp.DocumentsType) > 0 {
		documentsType = bson.E{Key: "commonpropertiesobjectstix.type", Value: bson.D{{Key: "$in", Value: sp.DocumentsType}}}
	}

	timeCreateIsExist := sp.Created.Start.Unix() > 0 && sp.Created.End.Unix() > 0
	timeModifiedIsExist := sp.Modified.Start.Unix() > 0 && sp.Modified.End.Unix() > 0

	if timeCreateIsExist && timeModifiedIsExist {
		dataTimeActionForObject = bson.E{Key: "$or", Value: bson.A{
			bson.D{{Key: "commonpropertiesdomainobjectstix.created", Value: bson.D{
				{Key: "$gte", Value: sp.Created.Start},
				{Key: "$lte", Value: sp.Created.End},
			}}},
			bson.D{{Key: "commonpropertiesdomainobjectstix.modified", Value: bson.D{
				{Key: "$gte", Value: sp.Modified.Start},
				{Key: "$lte", Value: sp.Modified.End},
			}}},
		}}
	} else if !timeCreateIsExist && timeModifiedIsExist {
		dataTimeActionForObject = bson.E{Key: "commonpropertiesdomainobjectstix.modified", Value: bson.D{
			{Key: "$gte", Value: sp.Modified.Start},
			{Key: "$lte", Value: sp.Modified.End},
		}}
	} else if timeCreateIsExist && !timeModifiedIsExist {
		dataTimeActionForObject = bson.E{Key: "commonpropertiesdomainobjectstix.created", Value: bson.D{
			{Key: "$gte", Value: sp.Created.Start},
			{Key: "$lte", Value: sp.Created.End},
		}}
	}

	if sp.CreatedByRef != "" {
		createdByRef = bson.E{Key: "commonpropertiesdomainobjectstix.created_by_ref", Value: sp.CreatedByRef}
	}

	if len(sp.SpecificSearchFields) == 0 {
		return bson.D{
			documentsType,
			dataTimeActionForObject,
			createdByRef,
		}
	}

	return bson.D{
		documentsType,
		dataTimeActionForObject,
		createdByRef,
		specificSearchFields,
	}

	/*
		type SpecificSearchFieldsSTIXObjectType struct {
			ObjectName   string                     `json:"object_name"`
			SearchFields SearchFieldsSTIXObjectType `json:"search_fields"`
		}

		type SearchFieldsSTIXObjectType struct {
			Name      string   `json:"name"`
			Aliases   []string `json:"aliases"`
			FirstSeen struct {
				Start time.Time `json:"start"`
				End   time.Time `json:"end"`
			} `json:"first_seen"`
			LastSeen struct {
				Start time.Time `json:"start"`
				End   time.Time `json:"end"`
			} `json:"last_seen"`
			Roles   []string `json:"roles"`
			Country string   `json:"country"`
			City    string   `json:"city"`
			URL     string   `json:"url"`
			Number  int      `json:"number"`
			Value   []string `json:"value"`
		}

		queryTemplate := map[string]bson.E{
						"sourceID":             (bson.E{Key: "source_id", Value: bson.D{{Key: "$eq", Value: sp.ID}}}),
						"filesIsFound":         (bson.E{Key: "detailed_information_on_filtering.number_files_found_result_filtering", Value: bson.D{{Key: "$gt", Value: 0}}}),
						"taskProcessed":        (bson.E{Key: "general_information_about_task.task_processed", Value: bson.D{{Key: "$eq", Value: sp.TaskProcessed}}}),
						"filesIsDownloaded":    (bson.E{Key: "detailed_information_on_downloading.number_files_downloaded", Value: bson.D{{Key: "$gt", Value: 0}}}),
						"filesIsNotDownloaded": (bson.E{Key: "detailed_information_on_downloading.number_files_downloaded", Value: bson.D{{Key: "$eq", Value: 0}}}),
						"allFilesIsDownloaded": (bson.E{Key: "$expr", Value: bson.D{
							{Key: "$eq", Value: bson.A{"$detailed_information_on_downloading.number_files_total", "$detailed_information_on_downloading.number_files_downloaded"}}}}),
						"allFilesIsNotDownloaded": (bson.E{Key: "$expr", Value: bson.D{
							{Key: "$ne", Value: bson.A{"$detailed_information_on_downloading.number_files_total", "$detailed_information_on_downloading.number_files_downloaded"}}}}),
						"sizeAllFiles": (bson.E{Key: "detailed_information_on_filtering.size_files_found_result_filtering", Value: bson.D{
							{Key: "$gte", Value: sp.InformationAboutFiltering.SizeAllFilesMin},
							{Key: "$lte", Value: sp.InformationAboutFiltering.SizeAllFilesMax},
						}}),
						"countAllFiles": (bson.E{Key: "detailed_information_on_filtering.number_files_found_result_filtering", Value: bson.D{
							{Key: "$gte", Value: sp.InformationAboutFiltering.CountAllFilesMin},
							{Key: "$lte", Value: sp.InformationAboutFiltering.CountAllFilesMax},
						}}),
						"dateTimeParameters": (bson.E{Key: "$and", Value: bson.A{
							bson.D{{Key: "filtering_option.date_time_interval.start", Value: bson.D{
								{Key: "$gte", Value: sp.InstalledFilteringOption.DateTime.Start}}}},
							bson.D{{Key: "filtering_option.date_time_interval.end", Value: bson.D{
								{Key: "$lte", Value: sp.InstalledFilteringOption.DateTime.End}}}},
						}}),
						"transportProtocol":      (bson.E{Key: "filtering_option.protocol", Value: sp.InstalledFilteringOption.Protocol}),
						"statusFilteringTask":    (bson.E{Key: "detailed_information_on_filtering.task_status", Value: sp.StatusFilteringTask}),
						"statusFileDownloadTask": (bson.E{Key: "detailed_information_on_downloading.task_status", Value: sp.StatusFileDownloadTask}),
					}

						!!! Тут надо написать формирование поискового запроса на основе полученных параметров !!!
									ПОКА ЭТО ПРОСТО ЗАГЛУШКА

									что бы не забыть: IPv4 нужно искать с учетом числового диапазона, а IPv6 как строку

									и еще надо подумать что делать с полями по которым необходимо выполнять сортировку
									пока, по умолчанию сортировка выполняется по стандартному полу MongoDB ObjectId
									Может быть в запрос ввести параметр для сортировки по полям (тип STIX объекта, время создания или
									модификации), но только что бы не повторяли ДОСЛОВНО название полей из коллекции (или может быть наоборот
									совпадали, но валидировались до того как будут включены в запрос)
	*/
}
