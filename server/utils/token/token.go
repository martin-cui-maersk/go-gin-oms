package token

import (
	"fmt"
	"github.com/martin-cui-maersk/go-gin-oms/global"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken 生成token
func GenerateToken(userId uint) (string, error) {
	// ttl, err := strconv.Atoi(os.Getenv("JWT_TTL"))
	// if err != nil {
	//	return "", err
	// }
	ttl := global.Server.JWT.ExpiresTime
	iss := global.Server.JWT.Iss
	sub := global.Server.JWT.Sub

	claims := jwt.MapClaims{}
	claims["iss"] = iss
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Second * time.Duration(ttl)).Unix()
	claims["nbf"] = time.Now().Unix()
	claims["sub"] = sub
	claims["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(global.Server.JWT.Secret))
}

// CheckTokenValid 检查token是否有效
func CheckTokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.Server.JWT.Secret), nil
	})
	if err != nil {
		return err
	}

	return nil
}

// ExtractToken 从请求头中获取token
func ExtractToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// ExtractTokenID 从jwt中解析出user_id
func ExtractTokenID(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.Server.JWT.Secret), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	// 如果jwt有效，将user_id转换为浮点数字符串，然后再转换为 uint32
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}

	return 0, nil
}
