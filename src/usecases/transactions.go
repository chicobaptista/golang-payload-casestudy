package usecases

type Transaction interface {
	Execute() (success bool, err error)
}
