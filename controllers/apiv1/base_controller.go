package apiv1

import "github.com/astaxie/beego"

// BaseAPIController ...
type BaseAPIController struct {
	beego.Controller
}

const InvalidInput = "400"
const NotFound = "404"
const InternalServerError = "500"

type success struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func (c *BaseAPIController) getStringParam(key string) string {
	value := c.GetString(key)
	if value == "" {
		c.Ctx.Input.SetData("error", "missing parameter "+key)
		c.Abort(InvalidInput)
	}
	return value
}
