package entities

type Platillo struct {
	ID                int     `json:"id"`
	Nombre            string  `json:"nombre"`
	Precio            float64 `json:"precio"`
	TiempoPreparacion int     `json:"tiempo_preparacion"`
}
