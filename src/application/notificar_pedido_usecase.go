package application

type Notificador interface {
	NotificarPedido(pedidoID int, estado string) error
}

type NotificarPedidoCompletadoUseCase struct {
	Notificador Notificador
}

func (uc *NotificarPedidoCompletadoUseCase) Execute(pedidoID int, estado string) error {
	return uc.Notificador.NotificarPedido(pedidoID, estado)
}
