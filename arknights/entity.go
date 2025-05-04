package arknights

type Resp[T any] struct {
	Code int
	Data T
	Msg  string
}
