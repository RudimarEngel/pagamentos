package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"database/sql"
  _ "github.com/go-sql-driver/mysql"

	// "reflect"
	"strconv"
	"io/ioutil"
	"encoding/json"

//	"./pkg"
)

type transferencia struct {
	ID		 string `json:"id"`
	Valor	 float64 `json:"valor"`
	IdPagante int `json:"IdPagante"`
	IdRecebedor	 int `json:"IdRecebedor"`
}

type usuario struct {
	UsuarioId 		int `json:"id"`
  UsuarioTipoId int `json:"UsuarioTipoId"`
  Nome 					string `json:"Nome"`
  CpfCnpj 			string `json:"CpfCnpj"`
  Email 				string `json:"Email"`
}

type autorizacao struct {
	Authorization bool
}

func dbConnection() *sql.DB {
	
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/pagamentos")
	if err != nil {
		panic(err.Error())
	}

	return db
}

func consultaAutorizacao() bool {
	resp, err := http.Get("https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31")
	if err != nil {
		fmt.Println("err: ", err)
	}
   defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var verify autorizacao
	_ = json.Unmarshal([]byte(string(body)), &verify)

	fmt.Println(" -> verify.Authorization: ", verify.Authorization)

	return verify.Authorization
}

func postTransferencia(c *gin.Context) {
	var novaTransferencia transferencia

	if err := c.BindJSON(&novaTransferencia); err != nil {
		fmt.Println("err: ", err)
		return
	}

	db := dbConnection();
	// Verifica o Id do recebedor, considerando que este seja o dono da máquina
	// Mas, acredito que o login deve ser da maquininha, ou de alguma outra ferramenta de pagamento.
	if varificarLogin(db, novaTransferencia.IdRecebedor) {

		// Verifica os saldos
		// var saldoPagador := consultaSaldo(Id)
		defer db.Close()

		// Consulta o servidor autorizador
		if !consultaAutorizacao() {
			c.IndentedJSON( http.StatusUnauthorized, "Acesso negado!" );
		}

		// Realiza a tarnsferência

		// Verifica os novos saldos
		
		// c.IndentedJSON(http.StatusCreated, novaTransferencia)
		c.IndentedJSON(http.StatusOK , "Pagamento efetuado!");
	} else {
		c.IndentedJSON( http.StatusUnauthorized, "Acesso negado!" );
	} 
}


func varificarLogin (db *sql.DB, UsuarioId int) bool {

	query := "SELECT Nome FROM Usuario WHERE UsuarioId = " + strconv.Itoa(UsuarioId) + ";"
	
	results, err := db.Query(query)
	if err !=nil {
			panic(err.Error())
	}
	if results.Next() {
		return true;
	} else {
		return false;
	}
	
}

func main() {
	router := gin.Default()
	// router.GET("/albums", getAlbums)
	router.POST("/transferencia", postTransferencia)

	router.Run("localhost:8080")
}
