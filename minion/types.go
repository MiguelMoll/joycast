package minion

type RedisHandler interface {
	Handle(msg string)
}
