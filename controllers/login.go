package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("my_secret_key")

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type Credentials struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var users = map[string]string{
	"username": "password",
	"user2":    "password2",
}

type LoginController struct {
	beego.Controller
}

func (lc *LoginController) Login() {
	var creds Credentials
	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &creds)
	fmt.Println(creds, 1, time.Now())
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		lc.Data["status"] = 400
		lc.Data["json"] = "Bad Request"
		lc.ServeJSON()
	}
	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		lc.Data["status"] = 401
		lc.Data["json"] = "Unauthorized"
		lc.ServeJSON()
	}
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(60 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	fmt.Println(tokenString, 1, time.Now())
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		lc.Data["status"] = 500
		lc.Data["json"] = "Internal server error"
		lc.ServeJSON()
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	lc.Data["status"] = 200
	lc.Data["json"] = tokenString
	lc.ServeJSON()
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXJuYW1lIiwiZXhwIjoxNjAxNjIzMjMyfQ.3hl8Os70Vvv0yjMIWLD3Kwp-wH2kWCKEaIJmavg03pA

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/astaxie/beego"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/go-redis/redis/v7"
//	"net/http"
//	"os"
//	"strings"
//	"time"
//)
//
//type LoginController struct {
//	beego.Controller
//}
//
//type User struct {
//	ID       uint64 `json:"id"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//	Phone    string `json:"phone"`
//}
//
//var user = User{
//	ID:       1,
//	Username: "username",
//	Password: "password",
//	Phone:    "49123454322", //this is a random number
//}
//
//type Claims struct {
//	Username string `json:"username"`
//	jwt.StandardClaims
//}
//
//// Login
//func (lc *LoginController) Login() {
//	var u User
//	fmt.Println("********************")
//	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &u)
//	fmt.Println(u)
//
//	if err != nil {
//		lc.Data["json"] = "Invalid json provided"
//		lc.Data["status"] = 404
//		lc.ServeJSON()
//	}
//	//compare the user from the request, with the one we defined:
//	if user.Username != u.Username || user.Password != u.Password {
//		lc.Data["json"] = "Please provide valid login details"
//		lc.Data["status"] = 404
//		lc.ServeJSON()
//	}
//	token, err := CreateToken(user.ID)
//	fmt.Println("********************", token)
//	if err != nil {
//		lc.Data["json"] = "unprocessable entity"
//		lc.Data["status"] = 404
//		lc.ServeJSON()
//	}
//	lc.Data["json"] = token
//	lc.Data["status"] = 200
//	lc.ServeJSON()
//}
//
//func CreateToken(userId uint64) (string, error) {
//	var err error
//	//Creating Access Token
//	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
//	atClaims := jwt.MapClaims{}
//	atClaims["authorized"] = true
//	atClaims["user_id"] = userId
//	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
//	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
//	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
//	if err != nil {
//		return "", err
//	}
//	fmt.Println(atClaims)
//	return token, nil
//}
//
//var client *redis.Client
//
//func init() {
//	//Initializing redis
//	dsn := os.Getenv("REDIS_DSN")
//	if len(dsn) == 0 {
//		dsn = "localhost:6379"
//	}
//	client = redis.NewClient(&redis.Options{
//		Addr: dsn, //redis port
//	})
//	_, err := client.Ping().Result()
//	fmt.Println(err)
//	//if err != nil {
//	//	panic(err)
//	//}
//}
//
//func ExtractToken(r *http.Request) string {
//	bearToken := r.Header.Get("Authorization")
//	//normally Authorization the_token_xxx
//	strArr := strings.Split(bearToken, " ")
//	if len(strArr) == 2 {
//		return strArr[1]
//	}
//	return "Authorized"
//}
//
//func (lc *LoginController) VerifyToken() (*jwt.Token, error) {
//	var r http.Request
//	tokenString := ExtractToken(&r)
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		//Make sure that the token method conform to "SigningMethodHMAC"
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(os.Getenv("ACCESS_SECRET")), nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	return token, nil
//}
//
//type AccessDetails struct {
//	AccessUuid string
//	UserId     uint64
//}
//
////
////func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
////	token, err := VerifyToken(r)
////	if err != nil {
////		return nil, err
////	}
////	claims, ok := token.Claims.(jwt.MapClaims)
////	if ok && token.Valid {
////		accessUuid, ok := claims["access_uuid"].(string)
////		if !ok {
////			return nil, err
////		}
////		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
////		if err != nil {
////			return nil, err
////		}
////		return &AccessDetails{
////			AccessUuid: accessUuid,
////			UserId:     userId,
////		}, nil
////	}
////	return nil, err
////}
//
//func DeleteAuth(givenUuid string) (int64, error) {
//	deleted, err := client.Del(givenUuid).Result()
//	if err != nil {
//		return 0, err
//	}
//	return deleted, nil
//}
//
////func Logout(c *gin.Context) {
////	au, err := ExtractTokenMetadata(c.Request)
////	if err != nil {
////		c.JSON(http.StatusUnauthorized, "unauthorized")
////		return
////	}
////	deleted, delErr := DeleteAuth(au.AccessUuid)
////	if delErr != nil || deleted == 0 { //if any goes wrong
////		c.JSON(http.StatusUnauthorized, "unauthorized")
////		return
////	}
////	c.JSON(http.StatusOK, "Successfully logged out")
////}
