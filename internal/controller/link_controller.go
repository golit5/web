package controller

import (
	"myapp/internal/model"
	"myapp/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LinkController struct {
	service service.LinkService
}

func NewLinkController(service service.LinkService) *LinkController {
	return &LinkController{service}
}

func (c *LinkController) CreateLink(ctx *gin.Context) {
	var link model.Link
	if err := ctx.ShouldBindJSON(&link); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateLink(&link); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create link"})
		return
	}
	ctx.JSON(http.StatusCreated, link)
}

func (c *LinkController) GetAllLinks(ctx *gin.Context) {
	links, err := c.service.GetAllLinks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch links"})
		return
	}
	ctx.JSON(http.StatusOK, links)
}

func (c *LinkController) GetLinkByID(ctx *gin.Context) {
	id_int, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}
	id := uint(id_int)
	link, err := c.service.GetLinkByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}
	ctx.JSON(http.StatusOK, link)
}

func (c *LinkController) UpdateLink(ctx *gin.Context) {
	var link model.Link
	if err := ctx.ShouldBindJSON(&link); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.UpdateLink(&link); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update link"})
		return
	}
	ctx.JSON(http.StatusOK, link)
}

func (c *LinkController) DeleteLink(ctx *gin.Context) {
	id_int, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}
	id := uint(id_int)
	if err := c.service.DeleteLink(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete link"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Link deleted successfully"})
}
