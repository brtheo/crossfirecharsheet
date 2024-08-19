package server

import (
	"crossfirecharsheet/cmd/web"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

// func (s *Server) RegisterRoutes() http.Handler {
// 	e := echo.New()
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())
// 	fileServer := http.FileServer(http.FS(web.Files))
// 	e.GET("/assets/*", echo.WrapHandler(fileServer))

// 	e.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))
// 	e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

// 	e.GET("/", s.HelloWorldHandler)

// 	return e
// }

// func (s *Server) HelloWorldHandler(c echo.Context) error {
// 	resp := map[string]string{
// 		"message": "Hello World",
// 	}

// 	return c.JSON(http.StatusOK, resp)
// }

func (app *PBApp) PocketBaseRoutes(e *core.ServeEvent) error {
	fileServer := http.FileServer(http.FS(web.Files))
	e.Router.GET("/assets/*", echo.WrapHandler(fileServer))
	e.Router.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))
	e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
	e.Router.POST("/hello", app.helloHandler)
	return nil
}

func (app *PBApp) helloHandler(c echo.Context) error {
	collection, err := app.App.Dao().FindCollectionByNameOrId("contacts")
	if err != nil {
		panic("rip")
	}

	record := models.NewRecord(collection)
	name := c.Request().FormValue("name")
	form := forms.NewRecordUpsert(app.App, record)
	// record.Set("firstName", name)
	form.LoadData(map[string]any{
		"firstName": name,
	})

	if err := form.Submit(); err != nil {

		panic(err)
	}

	return HTML(c, web.HelloPost(name))
}

func HTML(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
