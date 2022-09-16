package services

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/anditakaesar/uwa-back/domain"
)

const (
	DefaultPageSize    = 10
	DefaultCurrentPage = 1
)

type UserServiceInterface interface {
	GetAllUsers(param GetAllUsersRequest) *GetAllUsersResponse
}

type UserService struct {
	Ctx *Context
}

func NewUserService(ctx *Context) UserServiceInterface {
	return &UserService{
		Ctx: ctx,
	}
}

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

func (usvc *UserService) GetAllUsers(param GetAllUsersRequest) *GetAllUsersResponse {
	response := GetAllUsersResponse{}
	userReponse := []UserResponse{}
	users := []domain.User{}
	var count int64
	usvc.Ctx.DB.Find(&users).
		Offset((param.CurrentPage - 1) * param.PageSize).
		Limit(param.PageSize)

	usvc.Ctx.DB.Model(&users).Count(&count)

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
