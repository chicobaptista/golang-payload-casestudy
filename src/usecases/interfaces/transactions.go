package interfaces

type Transaction interface {
	Execute() (success bool, err error)
}
