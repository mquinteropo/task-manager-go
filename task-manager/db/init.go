package db

import (
	"fmt"
	"log"
	"task-manager/models"
)

func MigrateDB() {
	// Deshabilitar las claves foráneas temporalmente
	DB.Exec("SET FOREIGN_KEY_CHECKS=0;")

	// Eliminar las tablas en orden de dependencia para evitar conflictos
	DB.Migrator().DropTable(&models.TaskLog{})
	DB.Migrator().DropTable(&models.Task{})
	DB.Migrator().DropTable(&models.User{})
	DB.Migrator().DropTable(&models.Report{})
	DB.Migrator().DropTable(&models.Setting{})

	// Volver a habilitar las claves foráneas
	DB.Exec("SET FOREIGN_KEY_CHECKS=1;")

	// Migrar las tablas con la estructura corregida
	err := DB.AutoMigrate(&models.User{}, &models.Task{}, &models.TaskLog{}, &models.Report{}, &models.Setting{})
	if err != nil {
		log.Fatal(" Error en la migración de la base de datos:", err)
	}
	fmt.Println(" Migración de la base de datos completada")
}
