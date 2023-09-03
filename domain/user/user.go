package user

import (
	"context"

	userModel "github.com/anditakaesar/uwa-back/adapter/models/user"
	appCtx "github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/application/dto"
	roleSvc "github.com/anditakaesar/uwa-back/application/services/role"
	userSvc "github.com/anditakaesar/uwa-back/application/services/user"
	"github.com/anditakaesar/uwa-back/internal/crypter"
	"github.com/anditakaesar/uwa-back/internal/errs"
	"github.com/anditakaesar/uwa-back/internal/xjwt"
	"go.uber.org/zap"
)

type EntityDependency struct {
	UserService userSvc.UserSeviceInterface
	RoleService roleSvc.RoleSeviceInterface
	AppContext  appCtx.AppContext
}

type UserEntity struct {
	UserService userSvc.UserSeviceInterface
	RoleService roleSvc.RoleSeviceInterface
	AppContext  appCtx.AppContext
}

func CreateEntity(d EntityDependency) *UserEntity {
	return &UserEntity{
		UserService: d.UserService,
		AppContext:  d.AppContext,
		RoleService: d.RoleService,
	}
}

func (e *UserEntity) ValidateCreateUser(ctx context.Context, request dto.UserRegistrationRequest) error {
	requiredFields := []string{}

	if request.Username == "" {
		requiredFields = append(requiredFields, "username")
	}

	if request.Email == "" {
		requiredFields = append(requiredFields, "email")
	}

	if request.Password == "" {
		requiredFields = append(requiredFields, "password")
	}

	if request.PhoneNumber == "" {
		requiredFields = append(requiredFields, "phoneNumber")
	}

	if request.Role == "" {
		requiredFields = append(requiredFields, "role")
	}

	if len(requiredFields) > 0 {
		return errs.RequiredFieldError{
			RequiredFields: requiredFields,
		}
	}

	if !e.AppContext.UtilInterface.IsPhoneNumberValid(request.PhoneNumber) {
		return errs.InvalidFieldError{
			InvalidFields: []string{"phoneNumber"},
		}
	}

	return nil
}

func (e *UserEntity) ValidateAndGetRoleID(ctx context.Context, roleName string) (int, error) {
	invalidRoleErr := errs.InvalidFieldError{
		InvalidFields: []string{"role"},
	}

	roleResult, err := e.RoleService.GetRoleByName(ctx, roleName)
	if err != nil {
		return 0, errs.InternalError{
			Code:    errs.GenerateErrorCode(errs.ServiceError, errs.ErrUnknown),
			Message: err.Error(),
		}
	}

	if roleResult.ID > 0 {
		return int(roleResult.ID), nil
	}

	return 0, invalidRoleErr
}

func (e *UserEntity) CreateUser(ctx context.Context, request dto.UserRegistrationRequest) error {
	hashedPassword := crypter.GetDefaultCrypter().GenerateHash(request.Password)
	roleID, err := e.ValidateAndGetRoleID(ctx, request.Role)
	if err != nil {
		return err
	}

	user := userModel.User{
		Username:    request.Username,
		Email:       request.Email,
		Password:    hashedPassword,
		PhoneNumber: request.PhoneNumber,
		RoleID:      roleID,
	}
	return e.UserService.CreateUser(ctx, &user)
}

func (e *UserEntity) ValidateLoginUser(ctx context.Context, request dto.UserLoginRequest) error {
	if request.Username == "" {
		return errs.RequiredFieldError{
			RequiredFields: []string{"username"},
		}
	}

	if request.Password == "" {
		return errs.RequiredFieldError{
			RequiredFields: []string{"password"},
		}
	}

	return nil
}

func (e *UserEntity) LoginUser(ctx context.Context, request dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	loginResponse := dto.UserLoginResponse{}
	user, err := e.UserService.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return loginResponse, err
	}

	isPasswordValid := crypter.GetDefaultCrypter().CompareHash(user.Password, request.Password)
	if !isPasswordValid {
		e.AppContext.Logger.Warning("invalid password", zap.Any("request", request))
		return loginResponse, errs.InvalidLoginErr{
			Message: "invalid login",
		}
	}

	generatedToken, err := e.AppContext.UtilInterface.GenerateNewToken(xjwt.CustomClaims{
		Username:    user.Username,
		Email:       user.Email,
		Role:        "",
		PhoneNumber: user.PhoneNumber,
	})
	if err != nil {
		e.AppContext.Logger.Error("failed to generate token", err, zap.Any("request", request))
		return loginResponse, err
	}

	loginResponse = dto.UserLoginResponse{
		Token: generatedToken,
		User: dto.UserLoginInfo{
			Username:    user.Username,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Role:        "",
		},
	}

	return loginResponse, nil
}
