package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

//	"./pkg"
)

type transferencia struct {
	ID		 string `json:"id"`
	Valor	 string `json:"valor"`
	IdPagante string `json:"IdPagante"`
	IdRecebedor	 float64 `json:"IdRecebedor"`
}

/*var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
  {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
  {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}*/

func postTransferencia(c *gin.Context) {
	fmt.Println("AAAAAAAAAAAAAAAAA!")

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
	// função de verificação dos dados de sessao

	//return true;
}

func main() {
	router := gin.Default()
	// router.GET("/albums", getAlbums)
	router.POST("/transferencia", postTransferencia)

	router.Run("localhost:8080")
}
