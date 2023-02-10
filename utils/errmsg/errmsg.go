package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//code = 1000... 用户模块错误
	ErrorEmailExist              = 1001
	ErrorPasswordWrong           = 1002
	ErrorEmailNotExist           = 1003
	ErrorTokenNotExist           = 1004
	ErrorTokenExpired            = 1005
	ErrorTokenWrong              = 1006
	ErrorTokenFormatWrong        = 1007
	ErrorUserNoRight             = 1008
	ErrorVerificationCodeError   = 1009
	ErrorVerificationCodeExpired = 1010
	ErrorUserNotExist            = 1011
	ErrorPasswordLess8           = 1012
	ErrorEmailUsed               = 1013

	//code = 2000...
	ErrorInsufficientStock = 2001
	ErrorCateNotExist      = 2002

	//code = 3000
	ErrorOrderInvalid    = 3001
	ErrorOrderCantPay    = 3002
	ErrorOrderNotExist   = 3003
	ErrorInvalidQuantity = 3004
	ErrorInvalidEmail    = 3005

	//code = 4000
	ErrorCateNameExist = 4001

	//code = 5000
	ErrorNoticeNotExist = 5001
)

var codeMsg = map[int]string{
	SUCCESS: "SUCCESS",
	ERROR:   "FAIL",

	ErrorEmailExist:              "ERROR_EMAIL_EXIST",
	ErrorPasswordWrong:           "ERROR_PASSWORD_WRONG",
	ErrorEmailNotExist:           "ERROR_EMAIL_NOT_EXIST",
	ErrorTokenNotExist:           "ERROR_TOKEN_NOT_EXIST",
	ErrorTokenExpired:            "ERROR_TOKEN_EXPIRED",
	ErrorTokenWrong:              "ERROR_TOKEN_WRONG",
	ErrorTokenFormatWrong:        "ERROR_TOKEN_FORMAT_WRONG",
	ErrorUserNoRight:             "ERROR_USER_NO_RIGHT",
	ErrorVerificationCodeError:   "ERROR_VERIFICATION_CODE_ERROR",
	ErrorVerificationCodeExpired: "ERROR_VERIFICATION_CODE_Expired",
	ErrorUserNotExist:            "ERROR_USER_NOT_EXIST",
	ErrorPasswordLess8:           "ERROR_PASSWORD_LESS_8",
	ErrorEmailUsed:               "ERROR_EMAIL_USED",

	ErrorInsufficientStock: "ERROR_INSUFFICIENT_STOCK",

	ErrorOrderInvalid:    "ERROR_ORDER_INVALID",
	ErrorOrderCantPay:    "ERROR_ORDER_CANT_PAY",
	ErrorOrderNotExist:   "ERROR_ORDER_NOT_EXIST",
	ErrorInvalidQuantity: "ERROR_INVALID_QUANTITY",
	ErrorInvalidEmail:    "ERROR_INVALID_EMAIL",

	ErrorCateNameExist:  "ERROR_CATE_NAME_EXIST",
	ErrorNoticeNotExist: "ERROR_NOTICE_NOT_EXIST",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
