package controllers

import (
	"fmt"
	"models"

	_ "github.com/astaxie/beego/logs"
)

func (c *MainController) GetStockdata() {
	var symbol, stock_exchange string
	c.Ctx.Input.Bind(&symbol, "symbol")
	c.Ctx.Input.Bind(&stock_exchange, "stock_exchange")
	lp, _ := models.GetStockData(symbol, stock_exchange)
	fmt.Println(lp)
	c.Data["json"] = lp
	c.ServeJSON()
}
