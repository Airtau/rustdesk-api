package admin

import (
    "github.com/gin-gonic/gin"
    "github.com/lejianwen/rustdesk-api/v2/service"
    "github.com/lejianwen/rustdesk-api/v2/http/response"
)

// ActiveConnectionsList возвращает список активных соединений с данными из peers
func ActiveConnectionsList(c *gin.Context) {
    page := GetPage(c)
    pageSize := GetPageSize(c)
    
    activeService := &service.ActiveConnectionsService{}
    list, total, err := activeService.ListActive(page, pageSize)
    
    if err != nil {
	response.Fail(c, err.Error())
	return
    }
    
    response.Success(c, response.PageData{
	List:     list,
	Total:    total,
	Page:     int64(page),
	PageSize: int64(pageSize),
    })
}