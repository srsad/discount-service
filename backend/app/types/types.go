package types

type MessageType string

const (
	SUCCESS MessageType = "success"
	ERROR   MessageType = "error"
	WARNING MessageType = "warning"
)

type StoreType string

const (
	STORE        StoreType = "store"
	PICKUP_POINT StoreType = "pickup-point"
)
