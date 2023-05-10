package client

type ClientInterface interface {
	CloseConnection() error
	SendMessage([]byte) (int, error)
	ReceiveMessage([]byte) (int, error)
}
