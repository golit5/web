package controller

import (
	"myapp/internal/model"
	"myapp/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// LinkController контроллер для работы с ссылками
// @Description Контроллер для создания, получения, обновления и удаления ссылок
type LinkController struct {
	service service.LinkService
}

// NewLinkController создает новый экземпляр LinkController
func NewLinkController(service service.LinkService) *LinkController {
	return &LinkController{service}
}

// CreateLink создает новую ссылку
// @Summary Создание новой ссылки
// @Description Создает ссылку в базе данных
// @Tags links
// @Accept json
// @Produce json
// @Param link body model.Link true "Link object"
// @Success 201 {object} model.Link
// @Failure 400 {object} string "Invalid request"
// @Router /links [post]
func (c *LinkController) CreateLink(ctx *gin.Context) {
	var link model.Link
	if err := ctx.ShouldBindJSON(&link); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	if err := c.service.CreateLink(&link); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to create link"})
		return
	}
	ctx.JSON(http.StatusCreated, link)
}

// GetAllLinks возвращает все ссылки
// @Summary Получить все ссылки
// @Description Получить все ссылки из базы данных
// @Tags links
// @Produce json
// @Success 200 {array} model.Link
// @Failure 500 {object} string "Internal server error"
// @Router /links [get]
func (c *LinkController) GetAllLinks(ctx *gin.Context) {
	links, err := c.service.GetAllLinks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to fetch links"})
		return
	}
	ctx.JSON(http.StatusOK, links)
}

// GetLinkByID возвращает ссылку по ID
// @Summary Получить ссылку по ID
// @Description Получить ссылку по ее уникальному идентификатору
// @Tags links
// @Produce json
// @Param id path int true "Link ID"
// @Success 200 {object} model.Link
// @Failure 404 {object} string "Link not found"
// @Router /links/{id} [get]
func (c *LinkController) GetLinkByID(ctx *gin.Context) {
	id_int, err := strconv.Atoi(ctx.Param("id"))
	id := uint(id_int)
	link, err := c.service.GetLinkByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": "Link not found"})
		return
	}
	ctx.JSON(http.StatusOK, link)
}

// UpdateLink обновляет информацию о ссылке
// @Summary Обновить ссылку
// @Description Обновить ссылку в базе данных
// @Tags links
// @Accept json
// @Produce json
// @Param id path int true "Link ID"
// @Param link body model.Link true "Updated link object"
// @Success 200 {object} model.Link
// @Failure 400 {object} string "Invalid request"
// @Failure 404 {object} string "Link not found"
// @Router /links/{id} [put]
func (c *LinkController) UpdateLink(ctx *gin.Context) {
	var link model.Link
	if err := ctx.ShouldBindJSON(&link); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	if err := c.service.UpdateLink(&link); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to update link"})
		return
	}
	ctx.JSON(http.StatusOK, link)
}

// DeleteLink удаляет ссылку по ID
// @Summary Удалить ссылку
// @Description Удалить ссылку из базы данных по ID
// @Tags links
// @Produce json
// @Param id path int true "Link ID"
// @Success 200 {object} string "Link deleted"
// @Failure 404 {object} string "Link not found"
// @Router /links/{id} [delete]
func (c *LinkController) DeleteLink(ctx *gin.Context) {
	id_int, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	id := uint(id_int)
	if err := c.service.DeleteLink(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to delete link"})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Link deleted successfully"})
}
