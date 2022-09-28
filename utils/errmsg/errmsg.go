package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//code = 1000... 用户模块错误
	ErrorEmailUsed        = 1001
	ErrorPasswordWrong    = 1002
	ErrorEmailNotExist    = 1003
	ErrorTokenNotExist    = 1004
	ErrorTokenRuntime     = 1005
	ErrorTokenWrong       = 1006
	ErrorTokenFormatWrong = 1007
	ErrorUserNoRight      = 1008
	//code = 2000...
	ErrorCatenameUsed = 2001
	ErrorCateNotExist = 2002

	//code = 3000
	ErrorArticleNotExist = 3001

	//code = 3000...
)

var codemsg = map[int]string{
	SUCCESS:               "SUCCESS",
	ERROR:                 "FAIL",
	ErrorEmailUsed:        "ERROR_EMAIL_USED",
	ErrorPasswordWrong:    "ERROR_PASSWORD_WRONG",
	ErrorEmailNotExist:    "ERROR_EMAIL_NOT_EXIST",
	ErrorTokenNotExist:    "ERROR_TOkEN_NOT_EXIST",
	ErrorTokenRuntime:     "ERROR_TOKEN_RUNTIME",
	ErrorTokenWrong:       "ERROR_TOKEN_WRONG",
	ErrorTokenFormatWrong: "ERROR_TOKEN_FORMAT_WRONG",
	ErrorCatenameUsed:     "ERROR_CATENAME_USED",
	ErrorArticleNotExist:  "ERROR_ARTICLE_NOT_EXIST",
	ErrorCateNotExist:     "ERROR_CATE_NOT_EXIST",
	ErrorUserNoRight:      "ERROR_USER_NO_RIGHT",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
