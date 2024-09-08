package main

import (
	"mime"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Method int

const (
	MList Method = iota
	MFind
	MInsert
	MUpdate
	MDelete
	MLAST
)

type Route func(gin.IRoutes, string, ...gin.HandlerFunc) gin.IRoutes

var Routes [MLAST]Route = [MLAST]Route{
	gin.IRoutes.GET,    //MList
	gin.IRoutes.GET,    //MFind
	gin.IRoutes.POST,   //MInsert
	gin.IRoutes.PUT,    //MUpdate
	gin.IRoutes.DELETE, //MDelete
}

func SetupRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetTrustedProxies(nil)
	r.LoadHTMLGlob("templates/*")
	mime.AddExtensionType(".css", "text/css")
	r.Static("/s/", "./assets/")

	for _, x := range BookMap {
		Routes[x.Method](r, x.Path, Handler[Book]((Endpoint[Book])(x)))
	}
	for _, x := range MovieMap {
		Routes[x.Method](r, x.Path, Handler[Movie]((Endpoint[Movie])(x)))
	}
	return r
}

func Handler[T Collectible](x Endpoint[T]) func(c *gin.Context) {
	return func(c *gin.Context) {
		var id int
		if x.Method == MFind || x.Method == MUpdate || x.Method == MDelete {
			p := c.Param("id")
			id, _ = strconv.Atoi(p)
		}
		var y any
		switch x.Method {
		case MList:
			rows := x.Item.DbList()
			y = gin.H{x.Kword: rows}
		case MFind:
			y = x.Item.DbFind(id)
		case MInsert:
			z := any(x.Item).(*T)
			c.BindJSON(z)
			x.Item.DbInsert(z)
			y = z
		case MUpdate:
			z := any(x.Item).(*T)
			c.BindJSON(z)
			x.Item.DbUpdate(z, id)
			y = z
		case MDelete:
			x.Item.DbDelete(id)
		}
		if x.Method != MDelete {
			c.HTML(http.StatusOK, x.Temp, y)
		}
	}
}

type Endpoint[T Collectible] struct {
	Path   string
	Item   Storable[T]
	Temp   string
	Kword  string
	Method Method
}

type EndpointBook Endpoint[Book]

var B *Book = &Book{}

var BookMap []EndpointBook = []EndpointBook{
	EndpointBook{Path: "/books", Item: B, Temp: "listbooks.html", Kword: "books", Method: MList},
	EndpointBook{Path: "/books/:id", Item: B, Temp: "viewbook.html", Method: MFind},
	EndpointBook{Path: "/books/:id/edit", Item: B, Temp: "editbook.html", Method: MFind},
	EndpointBook{Path: "/books", Item: B, Temp: "viewbook.html", Method: MInsert},
	EndpointBook{Path: "/books/:id", Item: B, Temp: "viewbook.html", Method: MUpdate},
	EndpointBook{Path: "/books/:id", Item: B, Method: MDelete},
}

type EndpointMovie Endpoint[Movie]

var M *Movie = &Movie{}

var MovieMap []EndpointMovie = []EndpointMovie{
	EndpointMovie{Path: "/movies", Item: M, Temp: "listmovies.html", Kword: "movies", Method: MList},
	EndpointMovie{Path: "/movies/:id", Item: M, Temp: "viewmovie.html", Method: MFind},
	EndpointMovie{Path: "/movies/:id/edit", Item: M, Temp: "editmovie.html", Method: MFind},
	EndpointMovie{Path: "/movies", Item: M, Temp: "viewmovie.html", Method: MInsert},
	EndpointMovie{Path: "/movies/:id", Item: M, Temp: "viewmovie.html", Method: MUpdate},
	EndpointMovie{Path: "/movies/:id", Item: M, Method: MDelete},
}
