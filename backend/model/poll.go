package model

type Poll struct {
	Id       string         `json:"id"`
	Name     string         `json:"name" binding:"required"`
	Question string         `json:"question" binding:"required"`
	Options  []string       `json:"options" binding:"required"`
	Result   map[string]int `json:"result"`
}
