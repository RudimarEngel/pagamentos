package main

type album struct {
	ID		 string `json;"id"`
	Title	 string `json;"title"`
	Artist string `json;"artist"`
	Price	 string `json;"price"`
}

var albuns []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
  {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
  {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums( c *gin.Context ) {
	c.IntentedJSON(http.StatusOK, albums)
}