package types

type Direction string

const (
	Forward Direction = "forward"
	Up      Direction = "up"
	Down    Direction = "down"
)

type Command struct {
	Direction Direction
	Units     int
}
