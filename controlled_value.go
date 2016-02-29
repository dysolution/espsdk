package espsdk

import (
	"encoding/json"

	"github.com/dysolution/sleepwalker"
)

type cv struct {
	// "friendly" description, suitable for an HTML form label
	Description string `json:"description,omitempty"`
	// machine-parseable, suitable for the "value" attribute in an HTML form
	Value string `json:"value,omitempty"`
}

// ControlledValues contains all of the fields whose values are controlled by
// the API, i.e., the ESP API will validate the values against those in this
// collection.
//
// The marshaled form of this struct does not match that of the ESP API's CV
// endpoint. To do so would require that each of the (currently 5) asset types
// be statically represented as struct properties. Populating each of these
// properties requires a switch statement or reflection, neither of which are
// ideal.
//
// Alternatively, the asset types, as well as the controlled fields for
// Releases, are represented as the top-level map keys within ControlledFields.
//
// Example: (the "controlled fields" key is added by the SDK and is not present
// in the raw JSON, which puts the asset types at the same level as the
// "batch types" key)
//     {
//             "batch_types": [
//                     "getty_creative_still",
//                     "getty_creative_video",
//                     "getty_editorial_still",
//                     "getty_editorial_video",
//                     "istock_creative_video"
//             ],
//             "controlled_fields": {
//                     "getty_creative_still": {
//                             "collection_code": [
//                                     {
//                                             "description": "AbleStock.com",
//                                             "value": "ABL"
//                                     },
//                                     ...
//                             ],
//                             ...
//                 },
//                     "istock_creative_video": {
//                             ...
//                     },
//                     ...
//             },
//     }
//
// Example access:
// fmt.Println(allCV.ControlledFields["releases"]["model_ethnicities"])
//
type ControlledValues struct {
	BatchTypes       []string                   `json:"batch_types,omitempty"`
	ControlledFields map[string]map[string][]cv `json:"controlled_fields,omitempty"`
}

func parseCV(result sleepwalker.Result) ControlledValues {
	var allCV ControlledValues
	allCV.ControlledFields = make(map[string]map[string][]cv)

	var payload map[string]interface{}
	json.Unmarshal(result.Payload, &payload)

	var batchTypes []string

	for objectType, data := range payload {
		switch childData := data.(type) {
		case []interface{}:
			for _, batchType := range childData {
				batchTypes = append(batchTypes, batchType.(string))
			}
			allCV.BatchTypes = batchTypes
		case map[string]interface{}:
			fields := parseCVMaps(childData)
			allCV.ControlledFields[objectType] = fields
		}
	}
	return allCV
}

func parseCVMaps(childData map[string]interface{}) map[string][]cv {
	var fields = make(map[string][]cv)
	for fieldName := range childData {
		var values []cv
		for _, field := range childData[fieldName].([]interface{}) {
			switch mapData := field.(type) {
			case map[string]interface{}:
				values = append(values, cv{
					Description: mapData["description"].(string),
					Value:       mapData["value"].(string),
				})
			}
		}
		fields[fieldName] = values
	}
	return fields
}
