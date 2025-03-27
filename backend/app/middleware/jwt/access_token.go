package jwt

import (
	"backend/app/config"
	"backend/app/config/constant"
	"backend/app/db/postgre"
	"backend/app/server"
	"backend/pkg/api/v1/auth/models"
	"backend/pkg/api/v1/auth/repository"
	"backend/pkg/api/v1/auth/repository/impl"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Claims struct {
	Username string `json:"username"`
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

type Connection struct {
	read  *gorm.DB
	write *gorm.DB
}

type Repository struct {
	auth repository.AuthRepo
}

type JwtAccess struct {
	conn       Connection
	repository Repository
}

var (
	start                       = time.Now().UTC()
	LOGIN_EXPIRATION_DURATION   = time.Duration(24) * time.Hour
	REFRESH_EXPIRATION_DURATION = start.AddDate(0, 0, 3)
	JWT_SIGNING_METHOD          = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY           = []byte(config.GetConfig().SecretKey)
)

func GenerateAccessToken(username string, userId int64, name string) (string, int64, error) {
	expirationTime := time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix()

	claims := &Claims{
		Username: username,
		UserId:   int(userId),
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	accessToken, err := token.SignedString([]byte(JWT_SIGNATURE_KEY))

	if err != nil {
		return "", 0, err
	}

	return accessToken, expirationTime, nil
}

func ValidateAccessToken(accessToken string) (*Claims, string, bool, error) {
	claims := &Claims{}
	dbConn, err := postgre.ConnectDB()
	if err != nil {
		return nil, "", false, err
	}
	Jwtconn := &JwtAccess{
		conn: Connection{
			write: dbConn,
			read:  dbConn,
		},
		repository: Repository{
			auth: impl.NewAuth(dbConn),
		},
	}

	tokenStrings := strings.Replace(accessToken, "Bearer ", "", -1)
	token, err := jwt.ParseWithClaims(tokenStrings, claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_SIGNATURE_KEY, nil
	})

	v, _ := err.(*jwt.ValidationError)
	if err != nil {
		if v.Errors == jwt.ValidationErrorExpired {
			return claims, "", true, err
		} else {
			if !token.Valid {
				return claims, "", true, constant.ErrUnauthorized
			}
			return nil, "", true, err
		}
	}

	req := models.RequestCredentialValidate{
		Username: claims.Username,
		UserId:   claims.UserId,
		Name:     claims.Name,
	}
	resultSQL, err := Jwtconn.repository.auth.GetCredential(req)
	if err != nil {
		return nil, "", true, err
	}

	return claims, resultSQL, false, nil
}

func GenerateRefreshToken(username string, userId int64, name string) (string, int64, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour).Unix()

	claims := &Claims{
		Username: username,
		UserId:   int(userId),
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	refreshToken, err := token.SignedString([]byte(JWT_SIGNATURE_KEY))
	if err != nil {
		return "", 0, err
	}

	return refreshToken, expirationTime, nil
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return server.ResponseStatusUnauthorized(c, "Missing Authorization header", nil, nil, nil)
		}

		claims, _, isInvalid, err := ValidateAccessToken(authHeader)
		if err != nil || isInvalid {
			return server.ResponseStatusUnauthorized(c, "Unauthorized", nil, nil, nil)
		}
		c.Set("claims", claims)
		return next(c)
	}
}
