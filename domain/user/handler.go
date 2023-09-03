package user

import (
	"net/http"

	"github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/application/dto"
	roleSvc "github.com/anditakaesar/uwa-back/application/services/role"
	userSvc "github.com/anditakaesar/uwa-back/application/services/user"
	"github.com/anditakaesar/uwa-back/internal/errs"
	"github.com/anditakaesar/uwa-back/internal/handler"
	"github.com/anditakaesar/uwa-back/internal/json"
	"go.uber.org/zap"
)

type HandlerDependency struct {
	AppContext  context.AppContext
	UserService userSvc.UserSeviceInterface
	RoleService roleSvc.RoleSeviceInterface
}

type Handler struct {
	Resp        handler.ResponseInterface
	AppContext  context.AppContext
	UserService userSvc.UserSeviceInterface
	RoleService roleSvc.RoleSeviceInterface
}

func NewHandler(d HandlerDependency) Handler {
	return Handler{
		Resp:        handler.NewResponse(handler.Dep{Log: d.AppContext.Logger}),
		AppContext:  d.AppContext,
		UserService: d.UserService,
		RoleService: d.RoleService,
	}
}

func (h Handler) CreateUser() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		var request dto.UserRegistrationRequest

		err := json.Decode(&request, r.Body)
		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusBadRequest,
				err, errs.GenerateErrorCode(errs.HandlerError, errs.ErrDecoder), err.Error())
		}

		entity := CreateEntity(EntityDependency{
			UserService: h.UserService,
			AppContext:  h.AppContext,
			RoleService: h.RoleService,
		})

		err = entity.ValidateCreateUser(r.Context(), request)
		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusBadRequest,
				err, errs.GenerateErrorCode(errs.ValidationError, errs.ErrInvalidValidation), err.Error())
		}

		err = entity.CreateUser(r.Context(), request)
		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusInternalServerError,
				err, errs.GenerateErrorCode(errs.ServiceError, errs.ErrUnknown), err.Error())
		}

		return h.Resp.SetOk(nil)
	}
}

func (h Handler) LoginUser() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		var request dto.UserLoginRequest

		err := json.Decode(&request, r.Body)
		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusBadRequest,
				err, errs.GenerateErrorCode(errs.HandlerError, errs.ErrDecoder), err.Error())
		}

		entity := CreateEntity(EntityDependency{
			UserService: h.UserService,
			AppContext:  h.AppContext,
			RoleService: h.RoleService,
		})

		err = entity.ValidateLoginUser(r.Context(), request)
		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusBadRequest,
				err, errs.GenerateErrorCode(errs.ValidationError, errs.ErrInvalidValidation), err.Error())
		}

		loginResponse, err := entity.LoginUser(r.Context(), request)
		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusBadRequest,
				err, errs.GenerateErrorCode(errs.HandlerError, errs.ErrUnknown), err.Error())
		}

		return h.Resp.SetOkWithStatus(http.StatusOK, loginResponse)
	}
}

func (h Handler) VerifyUser() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		ctx := r.Context()
		claims, err := h.AppContext.UtilInterface.GetClaimsFromContext(ctx)
		if err != nil {
			h.AppContext.Logger.Error("error getting claims from context", err, zap.Error(err))
			return h.Resp.SetErrorWithStatus(http.StatusInternalServerError,
				err, errs.GenerateErrorCode(errs.HandlerError, errs.ErrUnknown), err.Error())
		}

		return h.Resp.SetOkWithStatus(http.StatusOK, claims)
	}
}
