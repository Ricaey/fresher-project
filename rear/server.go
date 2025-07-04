package main

import (
    "encoding/json"
    "io"
    "log"
    "net/http"
    "strconv"
    "sync"
    "fresher-project/rear/db"
)

// 统一响应结构体
type Resp struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

// 评论结构体
type Comment struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Content string `json:"content"`
}

// 评论列表和互斥锁
var (
    comments   = make([]Comment, 0)
    commentsMu sync.Mutex
    nextID     = 1
)

// 获取评论处理函数
func GetComments(w http.ResponseWriter, r *http.Request) {
    // 解析分页参数
    pageStr := r.URL.Query().Get("page")
    sizeStr := r.URL.Query().Get("size")
    page, _ := strconv.Atoi(pageStr)
    size, _ := strconv.Atoi(sizeStr)
    if page < 1 {
        page = 1
    }

    commentsMu.Lock()
    defer commentsMu.Unlock()

    total := len(comments)
    var result []Comment

    // size=-1 返回所有评论
    if size == -1 {
        result = comments
    } else {
        start := (page - 1) * size
        end := start + size
        if start > total {
            result = []Comment{}
        } else {
            if end > total {
                end = total
            }
            result = comments[start:end]
        }
    }

    resp := Resp{
        Code: 0,
        Msg:  "success",
        Data: map[string]interface{}{
            "total":    total,
            "comments": result,
        },
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

// 添加评论处理函数
func AddComment(w http.ResponseWriter, r *http.Request) {
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

    commentsMu.Lock()
    comment := Comment{
        ID:      nextID,
        Name:    req.Name,
        Content: req.Content,
    }
    nextID++
    comments = append(comments, comment)
    commentsMu.Unlock()

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
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        resp := Resp{Code: 1, Msg: "参数错误", Data: nil}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return
    }

    commentsMu.Lock()
    defer commentsMu.Unlock()
    idx := -1
    for i, c := range comments {
        if c.ID == id {
            idx = i
            break
        }
    }
    if idx == -1 {
        resp := Resp{Code: 2, Msg: "评论不存在", Data: nil}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return
    }
    // 删除评论
    comments = append(comments[:idx], comments[idx+1:]...)

    resp := Resp{Code: 0, Msg: "success", Data: nil}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func main() {
    if err := db.InitDB(); err != nil {
        log.Fatal(err)
    }
    http.HandleFunc("/comment/get", GetComments)
    http.HandleFunc("/comment/add", AddComment)
    http.HandleFunc("/comment/delete", DeleteComment)
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "pong~")
    })
    log.Println("Server running at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}

