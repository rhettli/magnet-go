package model

import (
	"encoding/json"
	"time"
)

type SearchHash struct {
	Id string  `primary_key"`
	Name string
	Requests string
	LastSeen time.Time
}

func (M *SearchHash) TableName() string {
	return "search_hash"
}

type MiltTbHashDetai struct {
	FileList string
	InfoHash string
	Name string
	CreateTime string
	LastSeen string
	Length string
	Requests string


}

//注意json转换是双向的

type SPhinxHash struct {
	Error         string  `json:"error"`
	Warning   string        `json:"warning"`
	Status   json.Number    `json:"status"`
	Fields []string `json:"fields"`
	Matches []struct{
		Id json.Number  `json:"id"`
		Weight string   `json:"weight"`
		Attrs struct{
			Hash_id json.Number     `json:"hash_id"`
			Category json.Number    `json:"category"`
			Length json.Number      `json:"length"`
			Create_time json.Number `json:"create_time"`
			Last_seen json.Number   `json:"last_seen"`
			Name string
			Requests string
		}
	}

	Total string            `json:"total"`
	Total_found string      `json:"total_found"`
	Time string                     `json:"time"`

}






