package presenters

import "jocer/pkg/cfg"

const (
	CfgKeySecuredKeywords cfg.Key = "VIEWS_OPTIONS_SECURED_KEYWORDS"
	CfgKeyMaxStringLength cfg.Key = "VIEWS_OPTIONS_MAX_STRING_LENGTH"

	CfgDefaultKeySecuredKeyword = ""
	CfgDefaultMaxStringLength   = 5
)

type (
	Config struct {
		SecuredKeywords []string
		MaxStringLength int
	}
)
