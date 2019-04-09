package comment

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type Api struct{}

func New() *Api { return &Api{} }

func (api *Api) Routes(g *echo.Group) {
	log.Info("Registering comment routes")
	u := g.Group("comment/")
	u.GET("get", api.getOne)
	u.GET("getAll", api.getAll)
	u.PUT("update/:id", api.update)
	u.POST("create", api.create)
	u.DELETE("delete/:id", api.delete)
}

func (api *Api) getOne(ctx echo.Context) error { return nil }
func (api *Api) getAll(ctx echo.Context) error { return nil }
func (api *Api) update(ctx echo.Context) error { return nil }
func (api *Api) create(ctx echo.Context) error { return nil }
func (api *Api) delete(ctx echo.Context) error { return nil }
