package common

type BlogTypeResp struct {
	TypeId   int    `json:"type_id"`
	TypeName string `json:"type_name"`
	Count    int    `json:"count"`
}
