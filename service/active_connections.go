package service

import (
    "github.com/lejianwen/rustdesk-api/v2/model"
    "github.com/lejianwen/rustdesk-api/v2/model/custom_types"
    "time"
)

type ActiveConnectionsService struct{}

// ActiveConnection представляет активное соединение с данными из peers
type ActiveConnection struct {
    ID        uint   `json:"id"`
    PeerID    string `json:"peer_id"`
    Hostname  string `json:"hostname"`
    TargetIP  string `json:"target_ip"`
    FromPeer  string `json:"from_peer"`
    FromName  string `json:"from_name"`
    FromIP    string `json:"from_ip"`
    UUID      string `json:"uuid"`
    CreatedAt string `json:"created_at"`
    Action    string `json:"action"`
    CloseTime int64  `json:"close_time"`
}

// ListActive возвращает все активные соединения с данными из peers
func (as *ActiveConnectionsService) ListActive(page, pageSize uint) (list []*ActiveConnection, total int64, err error) {
    var auditConns []*model.AuditConn
    
    // Запрос к audit_conns с условием action='new' и close_time=0
    tx := DB.Model(&model.AuditConn{}).
        Where("action = ? AND close_time = ?", "new", 0).
        Order("id DESC")
    
    tx.Count(&total)
    tx.Scopes(Paginate(page, pageSize)).Find(&auditConns)
    
    // Собираем peer_id для запроса к peers
    peerIds := make([]string, 0)
    for _, conn := range auditConns {
        if conn.PeerId != "" {
            peerIds = append(peerIds, conn.PeerId)
        }
    }
    
    // Загружаем peers
    var peers []*model.Peer
    if len(peerIds) > 0 {
        DB.Where("id IN ?", peerIds).Find(&peers)
    }
    
    // Создаём карту peer_id -> peer
    peerMap := make(map[string]*model.Peer)
    for _, peer := range peers {
        peerMap[peer.Id] = peer
    }
    
    // Формируем результат
    list = make([]*ActiveConnection, 0)
    for _, conn := range auditConns {
        // Преобразуем AutoTime в строку, используя приведение к time.Time
        var createdAtStr string
        t := time.Time(conn.CreatedAt)
        if !t.IsZero() {
            createdAtStr = t.Format("2006-01-02 15:04:05")
        } else {
            createdAtStr = ""
        }
        
        activeConn := &ActiveConnection{
            ID:        conn.Id,
            PeerID:    conn.PeerId,
            FromPeer:  conn.FromPeer,
            FromName:  conn.FromName,
            FromIP:    conn.Ip,
            UUID:      conn.Uuid,
            CreatedAt: createdAtStr,
            Action:    conn.Action,
            CloseTime: conn.CloseTime,
        }
        
        // Добавляем данные из peers
        if peer, ok := peerMap[conn.PeerId]; ok {
            activeConn.Hostname = peer.Hostname
            activeConn.TargetIP = peer.LastOnlineIp
            if activeConn.TargetIP == "" {
                activeConn.TargetIP = peer.IP
            }
        } else {
            activeConn.Hostname = "-"
            activeConn.TargetIP = "-"
        }
        
        list = append(list, activeConn)
    }
    
    return list, total, nil
}
