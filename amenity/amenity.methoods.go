package amenity

import "encoding/json"

func (t *RequestExist) Unmarshal(entry interface{}) RequestExist {
	obj, _ := json.Marshal(entry)
	response := RequestExist{}
	_ = json.Unmarshal(obj, &response)
	return response
}
