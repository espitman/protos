package amenity

import "encoding/json"

func (t *RequestExist) Unmarshal(entry interface{}) RequestExist {
	obj, _ := json.Marshal(entry)
	response := RequestExist{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *RequestCreate) Unmarshal(entry interface{}) RequestCreate {
	obj, _ := json.Marshal(entry)
	response := RequestCreate{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *ResponseDetails) Unmarshal(entry interface{}) ResponseDetails {
	obj, _ := json.Marshal(entry)
	response := ResponseDetails{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *ResponseFindByIds) Unmarshal(entry interface{}) ResponseFindByIds {
	obj, _ := json.Marshal(entry)
	response := ResponseFindByIds{}
	_ = json.Unmarshal(obj, &response)
	return response
}
