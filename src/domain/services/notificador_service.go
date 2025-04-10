package services

type Notificador interface {
	NotificarPedido(pedidoID int, estado string) error
}
