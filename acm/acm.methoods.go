package acm

import "encoding/json"

func (t *RequestCreate) UnmarshalBatch(entry interface{}) RequestCreate {
	obj, _ := json.Marshal(entry)
	response := RequestCreate{}
	_ = json.Unmarshal(obj, &response)
	return response
}
