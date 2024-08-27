package models

type Group struct {
	GroupName string `json:"groupName"`
	Days      []Day  `json:"days"`
}


