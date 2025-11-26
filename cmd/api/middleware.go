package main

import (
	"event-api-app/internal/database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware 的職責：
// 1. 驗證 JWT token
// 2. 從資料庫載入用戶資料
// 3. 將用戶資料傳遞給後續處理器
func (app *application) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(app.jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		userId := claims["user_id"].(float64)

		user, err := app.models.Users.Get(int(userId))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		c.Set("user", user) // 將用戶物件存入 context
		c.Next()            // 呼叫下一個處理器
	}
}

// CORSMiddleware 處理跨域請求
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// 檢查用戶使用是否已經通過 email 驗證
func RequireVerifiedUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user") // 從 context 獲取用戶物件

		if !exists {
			// 如果 Context 中沒有用戶，返回 401 Unauthorized
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 確保用戶物件類型正確
		u, ok := user.(*database.User)
		if !ok || !u.Verified {
			// 如果用戶未驗證，返回 403 Forbidden
			c.JSON(http.StatusForbidden, gin.H{"error": "Email not verified"})
			c.Abort()
			return
		}

		c.Next()
	}
}
