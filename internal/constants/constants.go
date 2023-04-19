package constants

type VerificationType int

const (
	APITokenValue    VerificationType = 1
	AccessTokenValue VerificationType = 2
)

var VerificationTypeConstants = struct {
	APIToken    VerificationType
	AccessToken VerificationType
}{
	APIToken:    APITokenValue,
	AccessToken: AccessTokenValue,
}
