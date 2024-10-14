package exception

func UnAuthorizedThunderException(message string) *ThunderException {
	return &ThunderException{
		Code:    401,
		Message: message,
	}
}
