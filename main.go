package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"database/sql"
  _ "github.com/go-sql-driver/mysql"

//	"./pkg"
)

type transferencia struct {
	ID		 string `json:"id"`
	Valor	 string `json:"valor"`
	IdPagante string `json:"IdPagante"`
	IdRecebedor	 float64 `json:"IdRecebedor"`
}

type usuario struct {
	UsuarioId 		int `json:"id"`
  UsuarioTipoId int `json:"UsuarioTipoId"`
  Nome 					string `json:"Nome"`
  CpfCnpj 			string `json:"CpfCnpj"`
  Email 				string `json:"Email"`
}

func dbConnection() {
	fmt.Println("pppppppp")
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/pagamentos")
	if err != nil {
		fmt.Println("PPPPPPPP")
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("db: ", db)
	results, err := db.Query("SELECT Nome FROM Usuario where UsuarioId = 1;")
	fmt.Println("results: ", results)
	if err !=nil {
			panic(err.Error())
	}
	for results.Next() {
		var usuario usuario
		err = results.Scan(&usuario.Nome)
		if err !=nil {
				panic(err.Error())
		}
		fmt.Println("Nome: " ,usuario.Nome)
	}

	fmt.Println("Success!")
}

func postTransferencia(c *gin.Context) {
	fmt.Println("AAAAAAAAAAAAAAAAA!")

	dbConnection();

	varificarLogin();

	var novaTransferencia transferencia

	if err := c.BindJSON(&novaTransferencia); err != nil {
		return
	}

//	L.Demo()
	
	c.IndentedJSON(http.StatusCreated, novaTransferencia)
}


func varificarLogin() {
	fmt.Println("BBBBBBBBBBBBBBBB!")
	// função de verificação dos dados de sessao, usando o token

	//return true;
}

func main() {
	router := gin.Default()
	// router.GET("/albums", getAlbums)
	router.POST("/transferencia", postTransferencia)

	router.Run("localhost:8080")
}
