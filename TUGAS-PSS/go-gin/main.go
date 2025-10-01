package main

import (
  "net/http"
  "log"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "github.com/gin-gonic/gin"
)

type Product struct {
  ID uint `gorm:"primaryKey" json:"id"`
  Name string `json:"name"`
  Category string `json:"category"`
  Price float64 `json:"price"`
  Stock int `json:"stock"`
  Description string `json:"description"`
  ImageURL string `json:"image_url"`
  CreatedAt  gorm.DeletedAt `gorm:"autoCreateTime" json:"created_at"`
  UpdatedAt  gorm.DeletedAt `gorm:"autoUpdateTime" json:"updated_at"`
}

func main() {
  dsn := "root:@tcp(127.0.0.1:3306)/cute_store?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal(err)
  }

  db.AutoMigrate(&Product{})

  r := gin.Default()

  r.GET("/products", func(c *gin.Context) {
    var products []Product
    db.Find(&products)
    c.JSON(http.StatusOK, products)
  })

  r.GET("/products/:id", func(c *gin.Context) {
    var p Product
    if err := db.First(&p, c.Param("id")).Error; err != nil {
      c.JSON(http.StatusNotFound, gin.H{"error":"not found"})
      return
    }
    c.JSON(http.StatusOK, p)
  })

  r.POST("/products", func(c *gin.Context) {
    var p Product
    if err := c.ShouldBindJSON(&p); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()}); return }
    db.Create(&p)
    c.JSON(http.StatusCreated, p)
  })

  r.PUT("/products/:id", func(c *gin.Context) {
    var p Product
    if err := db.First(&p, c.Param("id")).Error; err != nil { c.JSON(http.StatusNotFound, gin.H{"error":"not found"}); return }
    var input Product
    if err := c.ShouldBindJSON(&input); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()}); return }
    db.Model(&p).Updates(input)
    c.JSON(http.StatusOK, p)
  })

  r.DELETE("/products/:id", func(c *gin.Context) {
    if err := db.Delete(&Product{}, c.Param("id")).Error; err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, gin.H{"message":"deleted"})
  })

  r.Run(":8080")
}
