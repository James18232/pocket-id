package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stonith404/pocket-id/backend/internal/dto"
	"github.com/stonith404/pocket-id/backend/internal/middleware"
	"github.com/stonith404/pocket-id/backend/internal/service"
)

func NewUserGroupController(group *gin.RouterGroup, jwtAuthMiddleware *middleware.JwtAuthMiddleware, userGroupService *service.UserGroupService) {
	ugc := UserGroupController{
		UserGroupService: userGroupService,
	}

	group.GET("/user-groups", jwtAuthMiddleware.Add(true), ugc.list)
	group.GET("/user-groups/:id", jwtAuthMiddleware.Add(true), ugc.get)
	group.POST("/user-groups", jwtAuthMiddleware.Add(true), ugc.create)
	group.PUT("/user-groups/:id", jwtAuthMiddleware.Add(true), ugc.update)
	group.DELETE("/user-groups/:id", jwtAuthMiddleware.Add(true), ugc.delete)
	group.PUT("/user-groups/:id/users", jwtAuthMiddleware.Add(true), ugc.updateUsers)
}

type UserGroupController struct {
	UserGroupService *service.UserGroupService
}

func (ugc *UserGroupController) list(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	searchTerm := c.Query("search")

	groups, pagination, err := ugc.UserGroupService.List(searchTerm, page, pageSize)
	if err != nil {
		c.Error(err)
		return
	}

	var groupsDto = make([]dto.UserGroupDtoWithUserCount, len(groups))
	for i, group := range groups {
		var groupDto dto.UserGroupDtoWithUserCount
		if err := dto.MapStruct(group, &groupDto); err != nil {
			c.Error(err)
			return
		}
		groupDto.UserCount, err = ugc.UserGroupService.GetUserCountOfGroup(group.ID)
		if err != nil {
			c.Error(err)
			return
		}
		groupsDto[i] = groupDto
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       groupsDto,
		"pagination": pagination,
	})
}

func (ugc *UserGroupController) get(c *gin.Context) {
	group, err := ugc.UserGroupService.Get(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	var groupDto dto.UserGroupDtoWithUsers
	if err := dto.MapStruct(group, &groupDto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, groupDto)
}

func (ugc *UserGroupController) create(c *gin.Context) {
	var input dto.UserGroupCreateDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	group, err := ugc.UserGroupService.Create(input)
	if err != nil {
		c.Error(err)
		return
	}

	var groupDto dto.UserGroupDtoWithUsers
	if err := dto.MapStruct(group, &groupDto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, groupDto)
}

func (ugc *UserGroupController) update(c *gin.Context) {
	var input dto.UserGroupCreateDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	group, err := ugc.UserGroupService.Update(c.Param("id"), input)
	if err != nil {
		c.Error(err)
		return
	}

	var groupDto dto.UserGroupDtoWithUsers
	if err := dto.MapStruct(group, &groupDto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, groupDto)
}

func (ugc *UserGroupController) delete(c *gin.Context) {
	if err := ugc.UserGroupService.Delete(c.Param("id")); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (ugc *UserGroupController) updateUsers(c *gin.Context) {
	var input dto.UserGroupUpdateUsersDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	group, err := ugc.UserGroupService.UpdateUsers(c.Param("id"), input)
	if err != nil {
		c.Error(err)
		return
	}

	var groupDto dto.UserGroupDtoWithUsers
	if err := dto.MapStruct(group, &groupDto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, groupDto)
}
