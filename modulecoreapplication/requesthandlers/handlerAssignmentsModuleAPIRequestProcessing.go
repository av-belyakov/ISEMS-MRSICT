package requesthandlers

import (
	"encoding/json"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
)

//UnmarshalJSONCommonReq декодирует JSON документ, поступающий от модуля 'moduleapirequestprocessing', только по общим полям
func UnmarshalJSONCommonReq(msgReq *[]byte) (*datamodels.ModAPIRequestProcessingReqJSON, error) {
	var modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON
	err := json.Unmarshal(*msgReq, &modAPIRequestProcessingReqJSON)

	return &modAPIRequestProcessingReqJSON, err
}

//UnmarshalJSONObjectSTIXReq декодирует JSON документ, поступающий от модуля 'moduleapirequestprocessing', который содержит список объектов STIX
func UnmarshalJSONObjectSTIXReq(msgReq datamodels.ModAPIRequestProcessingReqJSON) ([]*datamodels.ListSTIXObject, bool, error) {
	var (
		isFail                             bool
		listResults                        []*datamodels.ListSTIXObject
		listSTIXObjectJSON                 datamodels.ModAPIRequestProcessingReqHandlingSTIXObjectJSON
		commonPropertiesObjectSTIX         datamodels.CommonPropertiesObjectSTIX
		numberSuccessfullyProcessedObjects int
	)

	if err := json.Unmarshal(*msgReq.RequestDetails, &listSTIXObjectJSON); err != nil {
		isFail = true

		return nil, isFail, err
	}

	for _, item := range listSTIXObjectJSON {
		err := json.Unmarshal(*item, &commonPropertiesObjectSTIX)
		if err != nil {
			/*
			   тут бы записать ошибку в лог-файл
			*/

			numberSuccessfullyProcessedObjects++
		}

		r, t, err := commonlibs.DecoderFromJSONToSTIXObject(commonPropertiesObjectSTIX.Type, item)
		if err != nil {
			/*
			   тут бы записать ошибку в лог-файл
			*/

			numberSuccessfullyProcessedObjects++
		}

		listResults = append(listResults, &datamodels.ListSTIXObject{
			DataType: t,
			Data:     r,
		})
	}

	if len(listSTIXObjectJSON) == numberSuccessfullyProcessedObjects {
		isFail = true
	}

	return listResults, isFail, nil
}
