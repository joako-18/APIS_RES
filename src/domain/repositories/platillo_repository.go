package repositories

import (
	"api/src/core"
	"api/src/domain/entities"
)

type PlatilloRepository struct{}

func (r PlatilloRepository) GetAll() ([]entities.Platillo, error) {
	rows, err := core.DB.Query("SELECT id, nombre, precio, tiempo_preparacion FROM platillos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var platillos []entities.Platillo
	for rows.Next() {
		var p entities.Platillo
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Precio, &p.TiempoPreparacion); err != nil {
			return nil, err
		}
		platillos = append(platillos, p)
	}

	return platillos, nil
}

func (r PlatilloRepository) Create(p entities.Platillo) error {
	_, err := core.DB.Exec("INSERT INTO platillos (nombre, precio, tiempo_preparacion) VALUES (?, ?, ?)",
		p.Nombre, p.Precio, p.TiempoPreparacion)
	return err
}

func (r PlatilloRepository) Update(p entities.Platillo) error {
	_, err := core.DB.Exec("UPDATE platillos SET nombre=?, precio=?, tiempo_preparacion=? WHERE id=?",
		p.Nombre, p.Precio, p.TiempoPreparacion, p.ID)
	return err
}

func (r PlatilloRepository) Delete(id int) error {
	_, err := core.DB.Exec("DELETE FROM platillos WHERE id=?", id)
	return err
}
