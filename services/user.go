package services

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/domain"
)

const (
	DefaultPageSize    = 10
	DefaultCurrentPage = 1
)

type GetAllUsersResponse struct {
	List []UserResponse `json:"list"`
	domain.Paging
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetAllUsersRequest struct {
	domain.Paging
}

func (param *GetAllUsersRequest) GetParamFromRequest(r *http.Request) error {
	param.PageSize = DefaultPageSize
	param.CurrentPage = DefaultCurrentPage
	pageSizeQ := r.URL.Query().Get("pageSize")
	numberRegex := regexp.MustCompile(`\d+`)
	if numberRegex.MatchString(pageSizeQ) {
		pageSize, err := strconv.Atoi(pageSizeQ)
		if err != nil {
			return err
		}

		param.PageSize = pageSize
	}

	currentPageQ := r.URL.Query().Get("currentPage")
	if numberRegex.MatchString(currentPageQ) {
		currentPage, err := strconv.Atoi(currentPageQ)
		if err != nil {
			return err
		}

		param.CurrentPage = currentPage
	}

	return nil
}

func GetAllUsers(appCtx application.Context, param GetAllUsersRequest) *GetAllUsersResponse {
	response := GetAllUsersResponse{}
	userReponse := []UserResponse{}
	users := []domain.User{}
	var count int64
	appCtx.DB.Find(&users).
		Offset((param.CurrentPage - 1) * param.PageSize).
		Limit(param.PageSize)

	appCtx.DB.Model(&users).Count(&count)

	for _, u := range users {
		userReponse = append(userReponse, UserResponse{
			ID:        u.ID,
			Username:  u.Username,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	response.List = userReponse
	response.Paging = domain.Paging{
		Count:       uint64(count),
		PageSize:    param.PageSize,
		CurrentPage: param.CurrentPage,
	}

	return &response
}
