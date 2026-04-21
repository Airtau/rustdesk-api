package admin

import (
    "github.com/gin-gonic/gin"
    "github.com/lejianwen/rustdesk-api/v2/http/controller"
    "github.com/lejianwen/rustdesk-api/v2/http/request/admin"
    "github.com/lejianwen/rustdesk-api/v2/http/response"
    "github.com/lejianwen/rustdesk-api/v2/service"
)

type ActiveConnectionsController struct {
    controller.Controller
}

// ListActiveConnections 获取活动连接列表
func (acc *ActiveConnectionsController) ListActiveConnections(ctx *gin.Context) {
    var req admin.ActiveConnectionsReq
    if err := ctx.ShouldBindQuery(&req); err != nil {
        response.Fail(ctx, 400, err.Error())
        return
    }
    
    // Получаем параметры пагинации
    page := req.Page
    if page <= 0 {
        page = 1
    }
    pageSize := req.PageSize
    if pageSize <= 0 {
        pageSize = 10
    }
    
    activeService := &service.ActiveConnectionsService{}
    list, total, err := activeService.ListActive(uint(page), uint(pageSize))
    if err != nil {
        response.Fail(ctx, 500, err.Error())
        return
    }
    
    response.Success(ctx, response.PageData{
        Total: int(total), // преобразуем int64 в int
        Page:  page,
        List:  list,
    })
}

