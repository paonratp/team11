package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/NoT-Ton/app/ent"
	"github.com/NoT-Ton/app/ent/user"
	"github.com/NoT-Ton/app/ent/location"
	"github.com/NoT-Ton/app/ent/bookborrow"
	"github.com/gin-gonic/gin"
)

// BookreturnController defines the struct for the bookreturn controller
type BookreturnController struct {
	client *ent.Client
	router gin.IRouter
}

// Bookreturn defines the struct for the bookreturn
type Bookreturn struct {
	UserID     int
	LocationID  int
	BookborrowID int
	Deadline    string
}

// CreateBookreturn handles POST requests for adding bookreturn entities
// @Summary Create bookreturn
// @Description Create bookreturn
// @ID create-bookreturn
// @Accept   json
// @Produce  json
// @Param bookreturn body bookreturn true "bookreturn entity"
// @Success 200 {object} ent.bookreturn
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookreturns [post]
func (ctl *BookreturnController) CreateBookreturn(c *gin.Context) {
	obj := Bookreturn{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bookreturn binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.UserID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	l, err := ctl.client.Location.
		Query().
		Where(location.IDEQ(int(obj.LocationID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "location not found",
		})
		return
	}

	b, err := ctl.client.Bookborrow.
		Query().
		Where(bookborrow.IDEQ(int(obj.BookborrowID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	now := time.Now()
	then := time.Date(2021, 1, 8, 15, 37, 0, time.UTC)
	after := time.Date(2021, 1, 15, 15, 37, 0, time.UTC)
	diff := after.sub(then)
	times, err := now.ADD(diff)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	br, err := ctl.client.Bookreturn.
		Create().
		SetUser(u).
		SetLocation(l).
		SetBookborrow(b).
		SetDeadline(times).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, b)
}

// ListBookreturn handles request to get a list of bookreturn entities
// @Summary List bookreturn entities
// @Description list bookreturn entities
// @ID list-bookreturn
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bookreturn
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookreturns [get]
func (ctl *BookreturnController) ListBookreturn(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	bookreturns, err := ctl.client.Bookreturn.
		Query().
		WithUser().
		WithLocation().
		WithBookborrow().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bookreturns)
}

// DeleteBookreturn handles DELETE requests to delete a bookreturn entity
// @Summary Delete a bookreturn entity by ID
// @Description get bookreturn by ID
// @ID delete-bookreturn
// @Produce  json
// @Param id path int true "Bookreturn ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookreturns/{id} [delete]
func (ctl *BookreturnController) DeleteBookreturn(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Bookreturn.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewBookreturnController creates and registers handles for the bookreturn controller
func NewBookreturnController(router gin.IRouter, client *ent.Client) *BookreturnController {
	bc := &BookreturnController{
		client: client,
		router: router,
	}
	bc.register()
	return bc
}

// InitBookreturnController registers routes to the main engine
func (ctl *BookreturnController) register() {
	bookreturns := ctl.router.Group("/bookreturns")

	bookreturns.GET("", ctl.ListBookreturn)
	bookreturns.POST("", ctl.CreateBookreturn)
	bookreturns.DELETE(":id", ctl.DeleteBookreturn)
}