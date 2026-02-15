package database

import (
	"log"
<<<<<<< HEAD

=======
	"os"
>>>>>>> 3ec0fcc45b1353a0c45a67e4249a857d83244eb4
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
<<<<<<< HEAD
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
=======

	stringDeConexao := "host=" + os.Getenv("HOST") +
		" user=" + os.Getenv("USER") +
		" password=" + os.Getenv("PASSWORD") +
		" dbname=" + os.Getenv("DBNAME") +
		" port=" + os.Getenv("PORT") +
		" sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
>>>>>>> 3ec0fcc45b1353a0c45a67e4249a857d83244eb4
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
<<<<<<< HEAD
}
=======
}
>>>>>>> 3ec0fcc45b1353a0c45a67e4249a857d83244eb4
