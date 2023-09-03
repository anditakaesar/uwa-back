package constants

type VerificationType int

const (
	APITokenValue VerificationType = 1
	JWTTokenValue VerificationType = 2
)

var VerificationTypeConstants = struct {
	APIToken VerificationType
	JWT      VerificationType
}{
	APIToken: APITokenValue,
	JWT:      JWTTokenValue,
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

const (
	PhoneNumberRegex string = "^628\\d{9,12}$"
)
