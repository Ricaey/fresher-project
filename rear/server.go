package main

import (
    "encoding/json"
    "io"
    "net/http"
    "strconv"

    "fresher-project/rear/db"
)

// 统一响应结构体
type Resp struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

// 获取评论处理函数
func GetComments(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
       return
    }
    pageStr := r.URL.Query().Get("page")
    sizeStr := r.URL.Query().Get("size")
    page, _ := strconv.Atoi(pageStr)
    size, _ := strconv.Atoi(sizeStr)
    if page < 1 {
        page = 1
    }
    if size < 1 && size != -1 {
        size = 10
    }

    var comments []db.Comment
    var total int64
    db.DB.Model(&db.Comment{}).Count(&total)

    query := db.DB.Order("id asc")
    if size != -1 {
        offset := (page - 1) * size
        query = query.Offset(offset).Limit(size)
    }
    query.Find(&comments)

    resp := Resp{
        Code: 0,
        Msg:  "success",
        Data: map[string]interface{}{
            "total":    total,
            "comments": comments,
        },
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

// 添加评论处理函数
func AddComment(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
    var req struct {
        Name    string `json:"name"`
        Content string `json:"content"`
    }
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "invalid body", http.StatusBadRequest)
        return
    }
    err = json.Unmarshal(body, &req)
    if err != nil || req.Name == "" || req.Content == "" {
        resp := Resp{Code: 1, Msg: "参数错误", Data: nil}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return
    }

    comment := db.Comment{
        Name:    req.Name,
        Content: req.Content,
    }
    if err := db.DB.Create(&comment).Error; err != nil {
        resp := Resp{Code: 2, Msg: "数据库写入失败", Data: nil}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return
    }

    resp := Resp{
        Code: 0,
        Msg:  "success",
        Data: comment,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

// 删除评论处理函数
func DeleteComment(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        resp := Resp{Code: 1, Msg: "参数错误", Data: nil}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return
    }

    if err := db.DB.Delete(&db.Comment{}, id).Error; err != nil {
        resp := Resp{Code: 2, Msg: "数据库删除失败", Data: nil}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return
    }

    resp := Resp{Code: 0, Msg: "success", Data: nil}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}


