package admin

import (
    "strconv"
    
    "github.com/gin-gonic/gin"
    "github.com/lejianwen/rustdesk-api/v2/http/response"
    "github.com/lejianwen/rustdesk-api/v2/service"
)

type ActiveConnections struct {
}

// ListActiveConnections 获取活动连接列表
// @Tags 活动连接
// @Summary 活动连接列表
// @Description 获取当前活跃的连接列表
// @Accept  json
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "页大小"
// @Success 200 {object} response.Response{data=response.PageData}
// @Failure 500 {object} response.Response
// @Router /admin/active_connections/list [get]
// @Security token
func (ac *ActiveConnections) ListActiveConnections(c *gin.Context) {
    page := 1
    pageSize := 10
    
    if p := c.Query("page"); p != "" {
	if val, err := strconv.Atoi(p); err == nil && val > 0 {
	    page = val
	}
    }
    if ps := c.Query("page_size"); ps != "" {
	if val, err := strconv.Atoi(ps); err == nil && val > 0 {
	    pageSize = val
	}
    }
    if pageSize > 100 {
	pageSize = 100
    }
    
    activeService := &service.ActiveConnectionsService{}
    list, total, err := activeService.ListActive(uint(page), uint(pageSize))
    if err != nil {
	response.Fail(c, 500, err.Error())
	return
    }
    
    response.Success(c, response.PageData{
	Total: int(total),
	Page:  page,
	List:  list,
    })
}
