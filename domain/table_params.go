package domain

type Paging struct {
	Count       uint64 `json:"count"`
	PageSize    int    `json:"pageSize"`
	CurrentPage int    `json:"currentPage"`
}
