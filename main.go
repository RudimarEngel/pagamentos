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
	Valor	 			float64 `json:"valor"`
	IdPagante 	int `json:"IdPagante"`
	IdRecebedor	int `json:"IdRecebedor"`
	MaquinaId		int `json:"MaquinaId"`
}

type usuario struct {
	UsuarioId 		int 	 `json:"id"`
	UsuarioTipoId int 	 `json:"UsuarioTipoId"`
	Nome 					string `json:"Nome"`
	CpfCnpj 			string `json:"CpfCnpj"`
	Email 				string `json:"Email"`
	Ativo					bool	 `json:"Ativo"`
	DeletedAt			string `json:"DeletedAt"`
	Tipo					string `json:"Tipo"`
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

	query := "UPDATE Conta SET Saldo = " + fmt.Sprintf("%f", novoSaldo) + 
						 " WHERE UsuarioId = " + strconv.Itoa(UsuarioId) + ";"
					
	_, err := db.Query(query)
	if err !=nil {
		return false
	}

	return true
}

func registroBilhetes(db *sql.DB, pagante int, recebedor int, acao int, maquina int, valor float64) bool {

	hora := time.Now()
	tabela := "Bilhetes_" + hora.Format("20060102")
	
	query := "CREATE TABLE IF NOT EXISTS " + tabela + "  (" +
							tabela + "Id bigint not null AUTO_INCREMENT," +
							"IdUsuarioPag bigint," +
							"IdUsuarioRec bigint," +
							"AcaoId bigint not null," +
							"Valor decimal(9,2)," +
							"MaquinaId bigint," +
							"CreatedAt timestamp default current_timestamp()," +
							"PRIMARY KEY (" + tabela + "Id)"+
						") AUTO_INCREMENT=1;"
	
	_, err := db.Query(query)
	if err !=nil {
		return false
	}

	// realiza o insert de dados
	query = "INSERT INTO " + tabela + "(IdUsuarioPag, IdUsuarioRec, AcaoId, Valor, MaquinaId, CreatedAt)" +
					"VALUES ("+ strconv.Itoa(pagante) + "," + strconv.Itoa(recebedor) + "," +
					strconv.Itoa(acao) + "," + fmt.Sprintf("%f", valor) +","+ strconv.Itoa(maquina) +",'" + 
					hora.Format("2006-01-02 15:04:05") +
					"');"
	
	_, err = db.Query(query)
	if err !=nil {
		return false
	}
	return true
}

func verificarUsuarios(db *sql.DB, recebedor int, pagante int) bool {

	if pagante == recebedor {
		return false;
	}

	// Verifica se o usuário pagante é do tipo comum
	query := "  select Usuario.Nome, UsuarioTipo.UsuarioTipoId, UsuarioTipo.Tipo"+
						" from Usuario "+
						" inner join UsuarioTipo on UsuarioTipo.UsuarioTipoId = Usuario.UsuarioTipoId " +
						" inner join Conta on Conta.UsuarioId = Usuario.UsuarioId " +
						" where Usuario.UsuarioId = "+ strconv.Itoa(pagante) + 
						"	 and UsuarioTipo.Tipo = 'comum' "+
						"	 and Usuario.Ativo = 1 " +
						"	 and Usuario.DeletedAt = '0000-00-00 00:00:00';" 
	
	results, err := db.Query(query)
	if err !=nil {
		panic(err.Error())
	}
	if !results.Next() {
		return false
	}

	// Verifica se o usuário recebedor é do tipo comum ou lojista
	query = " select Usuario.Nome, UsuarioTipo.UsuarioTipoId, UsuarioTipo.Tipo " +
					" from Usuario " +
					" inner join UsuarioTipo on UsuarioTipo.UsuarioTipoId = Usuario.UsuarioTipoId " +
					" inner join Conta on Conta.UsuarioId = Usuario.UsuarioId " +
					" where Usuario.UsuarioId = " + strconv.Itoa(recebedor) + 
					"	 and ( UsuarioTipo.Tipo = 'lojista' OR UsuarioTipo.Tipo = 'comum' ) " +
					"	 and Usuario.Ativo = 1 " +
					"	 and Usuario.DeletedAt = '0000-00-00 00:00:00';"

	results, err = db.Query(query)
	if err !=nil {
		panic(err.Error())
	}
	if !results.Next() {
		return false
	}

	return true;
} 

func postTransferencia(c *gin.Context) {
	var novaTransferencia transferencia
	if err := c.BindJSON(&novaTransferencia); err != nil {
		fmt.Println("err: ", err)
		return
	}

	db := dbConnection();

	if verificarUsuarios(db, novaTransferencia.IdRecebedor, novaTransferencia.IdPagante) {
		// Verifica o Id do recebedor, considerando que este seja o dono da máquina, ou o dono do app do smartphone
		if verificarLogin(db, novaTransferencia.IdRecebedor) {

			// Verifica o saldo do pagador
			saldoPagador := consultaSaldo(db, novaTransferencia.IdPagante);

			if saldoPagador.Saldo >= novaTransferencia.Valor {
				// consulta o saldo do recebedor
				saldoRecebedor := consultaSaldo(db, novaTransferencia.IdRecebedor);

				// Consulta o servidor autorizador
				if !consultaAutorizacao() {
					c.IndentedJSON( http.StatusUnauthorized, "Pagamento não autorizado!" ); // 401
					defer db.Close();
				}

				// Realiza a tarnsferência - alteração de saldos
				novoSaldoPagador   := saldoPagador.Saldo - novaTransferencia.Valor;
				novoSaldoRecebedor := saldoRecebedor.Saldo + novaTransferencia.Valor;
				operacaoExecutada := atualizarSaldo(db, novaTransferencia.IdPagante, novoSaldoPagador);

				if operacaoExecutada {
					operacaoExecutada = atualizarSaldo(db, novaTransferencia.IdRecebedor, novoSaldoRecebedor);
				}

				// Verifica os novos saldos?

				// Salva o registro na tabelha bilhetes para possíveis cancelamentos.
				if operacaoExecutada {
					registroBilhetes(db, novaTransferencia.IdPagante, novaTransferencia.IdRecebedor, 1, novaTransferencia.MaquinaId, novaTransferencia.Valor);
				} 

				// fecha a conexão com o banco
				defer db.Close();
				
				// c.IndentedJSON(http.StatusCreated, novaTransferencia)
				if operacaoExecutada {
					c.IndentedJSON(http.StatusOK , "Pagamento efetuado!");
				} else {
					registroBilhetes(db, novaTransferencia.IdPagante, novaTransferencia.IdRecebedor, 2, novaTransferencia.MaquinaId,novaTransferencia.Valor);
					defer db.Close();
					c.IndentedJSON(http.StatusInternalServerError , "Erro ao efetuar o pagamento!");
				}

			} else {
				// fecha a conexão com o banco
				registroBilhetes(db, novaTransferencia.IdPagante, novaTransferencia.IdRecebedor, 2, novaTransferencia.MaquinaId,novaTransferencia.Valor);
				defer db.Close();
				c.IndentedJSON(http.StatusUnauthorized , "Saldo Insuficiente!");
			}

		} else {
			// fecha a conexão com o banco
			registroBilhetes(db, novaTransferencia.IdPagante, novaTransferencia.IdRecebedor, 2, novaTransferencia.MaquinaId,novaTransferencia.Valor);
			defer db.Close();
			c.IndentedJSON( http.StatusUnauthorized, "Acesso negado!" );
		}
	} else {
		registroBilhetes(db, novaTransferencia.IdPagante, novaTransferencia.IdRecebedor, 2, novaTransferencia.MaquinaId,novaTransferencia.Valor);
		c.IndentedJSON(http.StatusUnauthorized , "Erro nos dados inseridos ou de cadastro!");
	} 
}

func verificarLogin (db *sql.DB, UsuarioId int) bool {

	// Esta consulta deve ser alterada para buscar dados de sessao e verifcação de token
	query := "SELECT Nome FROM Usuario WHERE UsuarioId = " + strconv.Itoa(UsuarioId) + ";";
	
	results, err := db.Query(query);
	if err !=nil {
			panic(err.Error());
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
	router.POST("/transferencia", postTransferencia);

	router.Run("localhost:8080");
}
