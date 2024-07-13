package handler

import (
	"TODOapp/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body model.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.ToDoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

type getAllListsResponse struct {
	Data []model.TodoList `json:"data"`
}

// @Summary		Get All Lists
// @Tags			lists
// @Description	get all todo lists for the authenticated user
// @ID				get-all-lists
// @Accept			json
// @Produce		json
// @Success		200		{object}	getAllListsResponse
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.services.ToDoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary		Get List By ID
// @Tags			lists
// @Description	get a todo list by its ID
// @ID				get-list-by-id
// @Accept			json
// @Produce		json
// @Param			id		path		int		true	"List ID"
// @Success		200		{object}	model.TodoList
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.services.ToDoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary		Update List
// @Tags			lists
// @Description	update a todo list
// @ID				update-list
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"List ID"
// @Param			input	body		model.UpadteListInput	true	"Updated list info"
// @Success		200
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var input model.UpadteListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.ToDoList.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
}

// @Summary		Delete List
// @Tags			lists
// @Description	delete a todo list
// @ID				delete-list
// @Accept			json
// @Produce		json
// @Param			id		path		int		true	"List ID"
// @Success		200		{object}	statusResponse
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.ToDoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})

}
