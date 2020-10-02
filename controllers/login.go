package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v7"
	"net/http"
	"os"
	"strings"
	"time"
)

type LoginController struct {
	beego.Controller
}

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
	Phone:    "49123454322", //this is a random number
}

// Login
func (lc *LoginController) Login() {
	var u User
	fmt.Println("********************")
	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &u)
	fmt.Println(u)

	if err != nil {
		lc.Data["json"] = "Invalid json provided"
		lc.Data["status"] = 404
		lc.ServeJSON()
	}
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		lc.Data["json"] = "Please provide valid login details"
		lc.Data["status"] = 404
		lc.ServeJSON()
	}
	token, err := CreateToken(user.ID)
	fmt.Println("********************", token)
	if err != nil {
		lc.Data["json"] = "unprocessable entity"
		lc.Data["status"] = 404
		lc.ServeJSON()
	}
	lc.Data["json"] = token
	lc.Data["status"] = 200
	lc.ServeJSON()
}

func CreateToken(userId uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	fmt.Println(atClaims)
	return token, nil
}

var client *redis.Client

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	fmt.Println(err)
	//if err != nil {
	//	panic(err)
	//}
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return "Authorized"
}

func (lc *LoginController) VerifyToken() (*jwt.Token, error) {
	var r http.Request
	tokenString := ExtractToken(&r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

//
//func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
//	token, err := VerifyToken(r)
//	if err != nil {
//		return nil, err
//	}
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if ok && token.Valid {
//		accessUuid, ok := claims["access_uuid"].(string)
//		if !ok {
//			return nil, err
//		}
//		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
//		if err != nil {
//			return nil, err
//		}
//		return &AccessDetails{
//			AccessUuid: accessUuid,
//			UserId:     userId,
//		}, nil
//	}
//	return nil, err
//}

func DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

//func Logout(c *gin.Context) {
//	au, err := ExtractTokenMetadata(c.Request)
//	if err != nil {
//		c.JSON(http.StatusUnauthorized, "unauthorized")
//		return
//	}
//	deleted, delErr := DeleteAuth(au.AccessUuid)
//	if delErr != nil || deleted == 0 { //if any goes wrong
//		c.JSON(http.StatusUnauthorized, "unauthorized")
//		return
//	}
//	c.JSON(http.StatusOK, "Successfully logged out")
//}
