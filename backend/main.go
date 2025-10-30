package main

import (
    "net/http"
    "os"
    "time" 

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
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

type Account struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Username  string    `json:"username" gorm:"unique;not null"`
    Password  string    `json:"-" gorm:"not null"` // JSONには含めない
    CreatedAt time.Time `json:"created_at"`
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
    if err := db.AutoMigrate(&Organization{}, &User{}, &Account{}); err != nil {
        panic(err)
    }

    e := echo.New()
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
        AllowHeaders: []string{echo.HeaderContentType},
    }))

    //認証API
    e.POST("/register", register)
    e.POST("/login", login)

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

func register(c echo.Context) error {
    type Req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    req := new(Req)
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
    }
    
    // パスワードをハッシュ化
    hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to hash password"})
    }
    
    account := Account{
        Username: req.Username,
        Password: string(hashed),
    }
    
    if err := db.Create(&account).Error; err != nil {
        return c.JSON(http.StatusConflict, map[string]string{"error": "username already exists"})
    }
    
    return c.JSON(http.StatusOK, map[string]interface{}{
        "id":       account.ID,
        "username": account.Username,
    })
}

// ログイン
func login(c echo.Context) error {
    type Req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    req := new(Req)
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
    }
    
    var account Account
    if err := db.Where("username = ?", req.Username).First(&account).Error; err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
    }
    
    // パスワード検証
    if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
    }
    
    return c.JSON(http.StatusOK, map[string]interface{}{
        "id":       account.ID,
        "username": account.Username,
        "message":  "login successful",
    })
}
