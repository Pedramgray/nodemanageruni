package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Node struct {
    ID       int    `json:"id" gorm:"primaryKey"`
    Name     string `json:"name"`
    Priority int    `json:"priority"`
}

var db *gorm.DB

func main() {
    var err error
    db, err = gorm.Open(sqlite.Open("nodes.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("خطا در اتصال به دیتابیس:", err)
    }
    db.AutoMigrate(&Node{})

    r := gin.Default()
    r.Static("/", "./static")

    r.POST("/add", func(c *gin.Context) {
        var newNode Node
        if err := c.ShouldBindJSON(&newNode); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر"})
            return
        }
        result := db.Create(&newNode)
        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "نود اضافه شد", "node": newNode})
    })

    r.GET("/find/:id", func(c *gin.Context) {
        var node Node
        id := c.Param("id")
        result := db.First(&node, id)
        if result.Error != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "یافت نشد"})
            return
        }
        c.JSON(http.StatusOK, node)
    })

    r.DELETE("/delete/:id", func(c *gin.Context) {
        var node Node
        id := c.Param("id")
        result := db.Delete(&node, id)
        if result.RowsAffected == 0 {
            c.JSON(http.StatusNotFound, gin.H{"error": "یافت نشد"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "نود حذف شد"})
    })

    r.POST("/boost", func(c *gin.Context) {
        type BoostInput struct {
            ID          int `json:"id"`
            NewPriority int `json:"new_priority"`
        }
        var input BoostInput
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر"})
            return
        }
        var node Node
        result := db.First(&node, input.ID)
        if result.Error != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "یافت نشد"})
            return
        }
        if input.NewPriority > node.Priority {
            node.Priority = input.NewPriority
            db.Save(&node)
            c.JSON(http.StatusOK, gin.H{"message": "اولویت افزایش یافت", "node": node})
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": "اولویت جدید باید بزرگ‌تر باشد"})
        }
    })

    r.POST("/handle", func(c *gin.Context) {
        var top Node
        result := db.Order("priority DESC").First(&top)
        if result.Error != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "صف خالی است"})
            return
        }
        db.Delete(&top)
        c.JSON(http.StatusOK, top)
    })

    r.GET("/queue", func(c *gin.Context) {
        var nodes []Node
        db.Order("priority DESC").Find(&nodes)
        c.JSON(http.StatusOK, nodes)
    })

    r.GET("/tree", func(c *gin.Context) {
        var nodes []Node
        db.Order("id ASC").Find(&nodes)
        c.JSON(http.StatusOK, nodes)
    })

    r.Run(":8080")
}
