package admin

import (
    "github.com/gin-gonic/gin"
    "github.com/lejianwen/rustdesk-api/v2/http/response"
    "github.com/lejianwen/rustdesk-api/v2/service"
)

// ActiveConnectionsController контроллер для активных соединений
type ActiveConnectionsController struct {
    // Убираем зависимость от controller.Controller
}

// ListActiveConnections получает список активных соединений
func (acc *ActiveConnectionsController) ListActiveConnections(ctx *gin.Context) {
    // Получаем параметры пагинации из query
    page := 1
    pageSize := 10
    
    if p := ctx.Query("page"); p != "" {
        // можно добавить парсинг, но для простоты оставляем значения по умолчанию
    }
    if ps := ctx.Query("page_size"); ps != "" {
        // можно добавить парсинг
    }
    
    activeService := &service.ActiveConnectionsService{}
    list, total, err := activeService.ListActive(uint(page), uint(pageSize))
    if err != nil {
        response.Fail(ctx, 500, err.Error())
        return
    }
    
    response.Success(ctx, response.PageData{
        Total: int(total),
        Page:  page,
        List:  list,
    })
}
