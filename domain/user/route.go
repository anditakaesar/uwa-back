package user

import (
	"net/http"

	"github.com/anditakaesar/uwa-back/application/context"
	roleSvc "github.com/anditakaesar/uwa-back/application/services/role"
	"github.com/anditakaesar/uwa-back/application/services/router"
	userSvc "github.com/anditakaesar/uwa-back/application/services/user"
	"github.com/anditakaesar/uwa-back/internal/constants"
)

type Dependecy struct {
	Context     router.Context
	UserService userSvc.UserSeviceInterface
	RoleService roleSvc.RoleSeviceInterface
	AppContext  context.AppContext
}

type Route struct {
	Context     router.Context
	UserService userSvc.UserSeviceInterface
	RoleService roleSvc.RoleSeviceInterface
	AppContext  context.AppContext
}

func NewDomain(d Dependecy) {
	route := Route(d)
	route.InitEndpoints()
}

func (r Route) InitEndpoints() {
	h := NewHandler(HandlerDependency{
		AppContext:  r.AppContext,
		UserService: r.UserService,
		RoleService: r.RoleService,
	})

	r.Context.RegisterEndpoint(r.CreateUser(h))
	r.Context.RegisterEndpoint(r.LoginUser(h))
	r.Context.RegisterEndpoint(r.VerifyUser(h))
}

func (r Route) CreateUser(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod: http.MethodPost,
		URLPattern: "/users",
		Handler:    h.CreateUser(),
		Verifications: []constants.VerificationType{
			constants.VerificationTypeConstants.APIToken,
		},
	}
}

func (r Route) LoginUser(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod: http.MethodPost,
		URLPattern: "/auth/login",
		Handler:    h.LoginUser(),
		Verifications: []constants.VerificationType{
			constants.VerificationTypeConstants.APIToken,
		},
	}
}

func (r Route) VerifyUser(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod: http.MethodGet,
		URLPattern: "/auth/verify",
		Handler:    h.VerifyUser(),
		Verifications: []constants.VerificationType{
			constants.VerificationTypeConstants.JWT,
		},
	}
}
