package transfer

type Job struct {
	ToCreate []string
	ToUpdate []string
	ToDelete []string
}

type OperationType uint8

const (
	OperationCreate OperationType = iota
	OperationUpdate               = iota
	OperationDelete               = iota
)
