package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Uwasaru/Sayka/presentation/json"
	"github.com/Uwasaru/Sayka/usecase"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

const oneWeek = 60 * 60 * 24 * 7

type AuthHandler struct {
	authUC usecase.IAuthUsecase
	userUC usecase.IUserUsecase
}

func NewAuthHandler(auc usecase.IAuthUsecase, uuc usecase.IUserUsecase) AuthHandler {
	return AuthHandler{
		authUC: auc,
		userUC: uuc,
	}
}

func (u *AuthHandler) Login(ctx *gin.Context) {
	redirectURL := ctx.Query("redirect_url")
	url, _, err := u.authUC.GetAuthURL(ctx, redirectURL)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	ctx.Redirect(http.StatusFound, url)
}

func (u *AuthHandler) Callback(ctx *gin.Context) {
	if errFormValue := ctx.Query("error"); errFormValue != "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": errFormValue},
		)
		return
	}
	state := ctx.Query("state")
	if state == "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": fmt.Errorf("state empty")},
		)
		return
	}

	code := ctx.Query("code")
	if code == "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": fmt.Errorf("code empty")},
		)
		return
	}
	// 認証
	redirectURL, sessionID, err := u.authUC.Authorization(ctx, state, code)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	if redirectURL == "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": fmt.Errorf("redirect url empty")},
		)
		return
	}

	// jwtの生成
	claims := jwt.MapClaims{
		"session": sessionID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // 72時間が有効期限
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, _ := token.SignedString([]byte(os.Getenv("rawPrivKey")))
	ctx.SetSameSite(http.SameSiteLaxMode)
	if os.Getenv("ENVIRONMENT") == "dev" {
		ctx.SetCookie("token", accessToken, oneWeek, "", "sayka.jp", true, true)
	} else {
		// フロントデプロイ次第、domainを登録
		ctx.SetCookie("token", accessToken, oneWeek, "", "sayka.jp", true, true)
	}
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Redirect(http.StatusFound, redirectURL)
}

// func (u *AuthHandler) User(ctx *gin.Context) {
// 	fmt.Println("User Usecase")
// 	// contextからuserIdを取得
// 	fmt.Println("ctx.Request.Context()", ctx.Request.Context())
// 	userId, ok := mycontext.GetUser(ctx.Request.Context())
// 	if !ok {
// 		fmt.Println("userId not found in context")
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
// 		return
// 	}
// 	fmt.Println("userId", userId)

// 	user, err := u.userUC.GetUser(ctx, userId)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	fmt.Println("user", user)

// 	ctx.JSON(http.StatusOK, gin.H{"user": user})
// }

func (u *AuthHandler) User(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("jwt")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("rawPrivKey")), nil
	})
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	str := ""
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		str, ok = claims["session"].(string)
		if !ok {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{"error": fmt.Errorf("kata asa-sion")},
			)
			return
		}
	} else {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": fmt.Errorf("token empty")},
		)
		return
	}

	// sessionの有効期限を確認
	ok, err := u.authUC.CheckSessionExpiry(ctx, str)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	if !ok {
		err := u.authUC.DeleteSession(ctx, str)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)
		}
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "session expiry err"},
		)
		return
	}
	// sessionからuserIdを取得
	userId, err := u.authUC.GetUserIdFromSession(ctx, str)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	user, err := u.userUC.GetUser(ctx, userId)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	userjson := json.UserEntityToJson(user)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": userjson},
	)
}

func (u *AuthHandler) SetCookieFromUserId(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	session, err := u.authUC.GetSessionFromUserId(ctx, userId)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	// jwtの生成
	claims := jwt.MapClaims{
		"session": session,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // 72時間が有効期限
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, _ := token.SignedString([]byte(os.Getenv("rawPrivKey")))
	ctx.SetSameSite(http.SameSiteLaxMode)
	if os.Getenv("ENVIRONMENT") == "dev" {
		ctx.SetCookie("token", accessToken, oneWeek, "", "", true, true)
	} else {
		// フロントデプロイ次第、domainを登録
		ctx.SetCookie("token", accessToken, oneWeek, "", "", true, true)
	}
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": "ok"},
	)
}
