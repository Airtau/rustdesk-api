package admin

type ActiveConnectionsReq struct {
    Page     int `form:"page"`
    PageSize int `form:"page_size"`
}
