package logger

type Service interface {
	Msg(string)
	Error(string)
}
