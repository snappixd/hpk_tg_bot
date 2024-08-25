package models

type GroupReplacement struct {
	GroupName    string `json:"groupName"`
	Replacements []Para `json:"replacements"`
}
