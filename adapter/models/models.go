package models

const (
	DefaultPageSize    = 10
	DefaultCurrentPage = 1
)

type ChangesInfo struct {
	CreatedBy      string `json:"createdBy"`
	LastModifiedBy string `json:"lastModifiedBy"`
	DeletedBy      string `json:"deletedBy"`
}

type Paging struct {
	Count       int `json:"count"`
	PageSize    int `json:"pageSize"`
	CurrentPage int `json:"currentPage"`
}
