package model

type Poll struct {
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Question string         `json:"question"`
	Options  []string       `json:"options"`
	Result   map[string]int `json:"result"`
}
