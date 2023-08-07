package domain

type ResponseType string
const (
	ResponseTypeCode ResponseType = "code"
)

type CodeChallengeMethodType string
const (
	CodeChallengeMethodS256 CodeChallengeMethodType = "S256"
)

type AuthorizeRequest struct {
	ResponseType        ResponseType            `form:"response_type" binding:"required"`
	ClientId            string                  `form:"client_id" binding:"required"`
	RedirectUri         string                  `form:"redirect_uri" binding:"required"`
	Scope               string                  `form:"scope" binding:"required"`
	State               string                  `form:"state" binding:"required"`
	CodeChallenge       string                  `form:"code_challenge" binding:"required"`
	CodeChallengeMethod CodeChallengeMethodType `form:"code_challenge_method" binding:"required"`
}

type Authorize struct {}

func (r *AuthorizeRequest) ResponseTypeCode() bool {
	return r.ResponseType == ResponseTypeCode
}

func (r *AuthorizeRequest) CodeChallengeMethodS256() bool {
	return r.CodeChallengeMethod == CodeChallengeMethodS256
}
