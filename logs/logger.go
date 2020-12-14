package logs

type ILogger interface {
	Info(str string)
	Debug(str string)
	Warning(str string)
	Error(str string)
}
