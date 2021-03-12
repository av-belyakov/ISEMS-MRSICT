package routingflowsmoduleapirequestprocessing

import (
	"encoding/json"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
)

//unmarshalJSONObjectSTIXReq декодирует JSON документ, поступающий от модуля 'moduleapirequestprocessing', который содержит список объектов STIX
func unmarshalJSONObjectSTIXReq(msgReq datamodels.ModAPIRequestProcessingReqJSON) ([]*datamodels.ListSTIXObject, struct{ All, Unsuccess int }, error) {
	var (
		listResults                []*datamodels.ListSTIXObject
		listSTIXObjectJSON         datamodels.ModAPIRequestProcessingReqHandlingSTIXObjectJSON
		commonPropertiesObjectSTIX datamodels.CommonPropertiesObjectSTIX
		numberUnsuccessfullyProc   int
	)

	if err := json.Unmarshal(*msgReq.RequestDetails, &listSTIXObjectJSON); err != nil {
		return nil, struct{ All, Unsuccess int }{0, 0}, err
	}

	for _, item := range listSTIXObjectJSON {
		err := json.Unmarshal(*item, &commonPropertiesObjectSTIX)
		if err != nil {
			numberUnsuccessfullyProc++
		}

		r, t, err := commonlibs.DecoderFromJSONToSTIXObject(commonPropertiesObjectSTIX.Type, item)
		if err != nil {
			numberUnsuccessfullyProc++
		}

		listResults = append(listResults, &datamodels.ListSTIXObject{
			DataType: t,
			Data:     r,
		})
	}

	return listResults, struct{ All, Unsuccess int }{len(listSTIXObjectJSON), numberUnsuccessfullyProc}, nil
}
