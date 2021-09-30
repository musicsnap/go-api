package middlewares

type ApiError struct {
	Level  int // -1 部分失败 fatal  0 成功 1 部分失败，无伤大雅
	Errors []string
}

var globalError *ApiError

func (e *ApiError) Error() string {
	return "error"
}

const (
	FATAL   = -1
	SUCESSS = 0
	WARNING = 1
)

func PushError(message string) {
	if globalError == nil {
		globalError = new(ApiError)
	}
	globalError.Errors = append(globalError.Errors, message)
}

func ResetError() {
	if globalError == nil {
		globalError = new(ApiError)
	}
	globalError.Errors = []string{}
}

func SetLevel(level int) {
	globalError.Level = level
}

func GetLevel() int {
	if globalError == nil {
		return SUCESSS
	}
	return globalError.Level
}

func GetErrors() []string {
	if globalError == nil {
		return make([]string, 0)
	}
	return globalError.Errors
}
