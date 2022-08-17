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
	"time"

)

type transferencia struct {
	// ID		 string `json:"id"`
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

type conta struct {
	ContaId 	int 		`json:"ContaId"`
	UsuarioId int 		`json:"UsuarioId"`
	Saldo	 		float64 `json:"Saldo"`
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

func consultaSaldo(db *sql.DB, UsuarioId int) conta {
	
	var resultadoConta conta

	query := "SELECT UsuarioId, Saldo FROM Conta WHERE UsuarioId = " + strconv.Itoa(UsuarioId) + ";"
	
	results, err := db.Query(query)
	if err !=nil {
		panic(err.Error())
	}
	for results.Next() {
		err = results.Scan(&resultadoConta.UsuarioId, &resultadoConta.Saldo)
		if err !=nil {
				panic(err.Error())
		}
	}
	
	return resultadoConta
}

func atualizarSaldo(db *sql.DB, UsuarioId int, novoSaldo float64) bool {

	fmt.Println("novoSaldo: ", novoSaldo)

	query := "UPDATE Conta SET Saldo = " + fmt.Sprintf("%f", novoSaldo) + 
						 " WHERE UsuarioId = " + strconv.Itoa(UsuarioId) + ";"
					
	_, err := db.Query(query)
	if err !=nil {
		return false
	}

	return true
}

func registroBilhetes(db *sql.DB) {

	// query := "SHOW TABLES LIKE 'pagamentos.Conta';"
	hora := time.Now()
	fmt.Println("AAAAAAAAAAAAAAAAA, hora: ", hora.Format("2006-01-02 15:04:05"))
	fmt.Println("AAAAAAAAAAAAAAAAA, hora: ", hora.Format("20060102"))
	tabela := "Bilhetes_" + hora.Format("20060102")
	query := "SELECT * FROM information_schema.tables "+
					 "WHERE table_schema = 'pagamentos'" +
						" AND table_name = '" + tabela + "' LIMIT 1;"

	fmt.Println("query: ", query)

	results, err := db.Query(query)
	if err !=nil {
			panic(err.Error())
	}
	if !results.Next() {
		fmt.Println(false);
		// Cria  a tabela do dia, caso ela não exista
		query = "CREATE TABLE IF NOT EXISTS " + tabela + " (" + 
	}

	// realiza o insert de dados
	fmt.Println("FAZ O INSERT")

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

		// Verifica o saldo do pagador
		saldoPagador := consultaSaldo(db, novaTransferencia.IdPagante)
		fmt.Println("saldoPagador.UsuarioId: ", saldoPagador.UsuarioId)
		fmt.Println("saldoPagador.Saldo: ", saldoPagador.Saldo)

		if saldoPagador.Saldo >= novaTransferencia.Valor {
			fmt.Println("Saldo Suficiente!")
			// consulta o saldo do recebedor
			saldoRecebedor := consultaSaldo(db, novaTransferencia.IdRecebedor)
			fmt.Println("saldoRecebedor.UsuarioId: ", saldoRecebedor.UsuarioId)
			fmt.Println("saldoRecebedor.Saldo: ", saldoRecebedor.Saldo)

			// Consulta o servidor autorizador
			if !consultaAutorizacao() {
				c.IndentedJSON( http.StatusUnauthorized, "Pagamento não autorizado!" ); // 401
				defer db.Close()
			}

			// Realiza a tarnsferência - alteração de saldos
			novoSaldoPagador   := saldoPagador.Saldo - novaTransferencia.Valor
			novoSaldoRecebedor := saldoRecebedor.Saldo + novaTransferencia.Valor
			operacaoExecutada := atualizarSaldo(db, novaTransferencia.IdPagante, novoSaldoPagador)

			if operacaoExecutada {
				operacaoExecutada = atualizarSaldo(db, novaTransferencia.IdRecebedor, novoSaldoRecebedor)
			}
			fmt.Println("operacaoExecutada: ", operacaoExecutada)

			// Verifica os novos saldos?

			// Salva o registro na tabelha bilhetes para possíveis cancelamentos.
			if operacaoExecutada {
				registroBilhetes(db);
			} /*else {
				// registra o erro no pagamento
			}*/


			// fecha a conexão com o banco
			defer db.Close()
			
			// c.IndentedJSON(http.StatusCreated, novaTransferencia)
			if operacaoExecutada {
				c.IndentedJSON(http.StatusOK , "Pagamento efetuado!");
			} else {
				c.IndentedJSON(http.StatusInternalServerError , "Erro ao efetuar o pagamento!");
			}
			

		} else {
			// fecha a conexão com o banco
			defer db.Close()
			fmt.Println("Saldo Insuficiente!")
			c.IndentedJSON(http.StatusOK , "Saldo Insuficiente!");
		}

		
	} else {
		// fecha a conexão com o banco
		defer db.Close()
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
