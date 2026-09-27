package dto

type Testimony struct {
	UserName string `json:"userName"`
	Message    string `json:"message"`
}

type SetTestimony struct {
	Message    string `json:"message"`
}
