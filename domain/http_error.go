package domain

type ErrorCode string

const (
	InvalidRequest       ErrorCode = "invalid_request"
	InvalidClient        ErrorCode = "invalid_client"
	InvalidGrant         ErrorCode = "invalid_grant"
	UnauthorizedClient   ErrorCode = "unauthorized_client"
	UnsupportedGrantType ErrorCode = "unsupported_grant_type"
	InvalidScope         ErrorCode = "invalid_scope"
)

type HTTPError struct {
	Error            ErrorCode `json:"error"`
	ErrorDescription string    `json:"error_description"`
	ErrorUri         string    `json:"error_uri"`
}
