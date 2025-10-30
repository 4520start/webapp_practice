package main

import (
    "net/http"
    "os"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type Organization struct {
    ID   uint   `json:"id" gorm:"primaryKey"`
    Name string `json:"name"`
}

type User struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    OrgID uint   `json:"org_id"`
}

var db *gorm.DB

func main() {
    dsn := os.Getenv("DATABASE_DSN")
    if dsn == "" {
        dsn = "host=db user=postgres password=pass dbname=myapp port=5432 sslmode=disable"
    }
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    db.AutoMigrate(&Organization{}, &User{})
    e := echo.New()
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
        AllowHeaders: []string{echo.HeaderContentType},
    }))
    e.GET("/users", getUsers)
    e.POST("/users", createUser)
    e.GET("/health", func(c echo.Context) error { return c.String(http.StatusOK, "ok") })

    e.Logger.Fatal(e.Start(":8080"))
}

func getUsers(c echo.Context) error {
    var users []User
    org := c.QueryParam("org_id")
    if org != "" {
        db.Where("org_id = ?", org).Find(&users)
    } else {
        db.Find(&users)
    }
    return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    db.Create(u)
    return c.JSON(http.StatusOK, u)
}
