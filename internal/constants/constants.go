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

const (
	ApplicationJsonMime = "application/json"
	TextPlainMime       = "text/plain"
)

var AvailableMimeType = struct {
	ApplicationJson string
	TextPlain       string
}{
	ApplicationJson: ApplicationJsonMime,
	TextPlain:       TextPlainMime,
}
