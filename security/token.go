package security

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"
)

type UserAuth struct {
	AccessToken  string
	RefreshToken string
}

// This holds the information required to create a personalised token.
type claims struct {
	Email  string
	UserID int
	jwt.StandardClaims
}

// Create jwtKey to be used in creating a specific signature.
var (
	jwtAccessKey  = []byte("client")
	jwtRefreshKey = []byte("refresh")
)

// Create Struct to attach methods.
type auth struct{}

func AuthenticateUser(email string, userID, accTknExpTime,
	RefTknExpTime int) (*UserAuth, error) {
	accTkn, err := generateAccessToken(email, userID, accTknExpTime)
	if err != nil {
		return nil, err
	}
	refTkn, err := generateRefreshToken(email, userID, RefTknExpTime)
	if err != nil {
		return nil, err
	}
	auth := UserAuth{
		AccessToken:  accTkn,
		RefreshToken: refTkn,
	}
	return &auth, nil
}

func AuthorizeRequest(req *http.Request) (*claims, error) {
	refreshToken, err := extractRefreshToken(req)
	if err != nil {
		return nil, err
	}
	if err := authenticateRefreshToken(refreshToken); err != nil {
		return nil, err
	}

	accessToken, err := extractAccessToken(req)
	if err != nil {
		return nil, err
	}
	return authenticateAccessToken(accessToken)
}

func RefreshUserTokens(req *http.Request, accTknExpTime,
	refTknExpTime int) (*UserAuth, error) {
	claims, err := AuthorizeRequest(req)
	if err != nil {
		return nil, err
	}
	return AuthenticateUser(claims.Email, claims.UserID, accTknExpTime, refTknExpTime)
}

func (a *auth) DeleteAccess(res http.ResponseWriter, req *http.Request) (int, error) {
	type response struct {
		Message     string
		AccessToken string
	}
	cookie, err := req.Cookie("refreshToken")
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return 0, errors.New("")
	}
	cookie.Value = ""
	http.SetCookie(res, cookie)

	// Reset the access/bearer token.
	resp, _ := json.Marshal(&response{
		Message:     "Log Out Successful",
		AccessToken: "",
	})
	res.Header().Add("Content-Type", "application/json")
	return res.Write(resp)
}

// Helper Functions.
func generateAccessToken(email string, userID, expTime int) (string, error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(expTime))
	claims := &claims{
		Email:  email,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtAccessKey)
}

func authenticateAccessToken(token string) (*claims, error) {
	claims := &claims{}
	tkn, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtAccessKey, nil
		})
	if err != nil {
		return nil, errors.New("can not parse the token")
	}
	if !tkn.Valid {
		return nil, errors.New("token is invalid or expired")
	}
	return claims, nil
}

func authenticateRefreshToken(token string) error {
	claims := &claims{}
	tkn, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtRefreshKey, nil
		})
	if err != nil {
		return errors.New("")
	}
	if !tkn.Valid {
		return errors.New("")
	}
	return nil
}

func generateRefreshToken(email string, userID, expTime int) (string, error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(expTime))
	claims := &claims{
		Email:  email,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtRefreshKey)
}

func extractAccessToken(req *http.Request) (string, error) {
	tokenStr := req.Header.Get("Authorization")
	if tokenStr == "" {
		return "", errors.New("Invalid header")
	}

	splitToken := strings.Split(tokenStr, "Bearer")
	if len(splitToken) != 2 {
		return "", errors.New("Invalid header")
	}

	tokenStr = strings.TrimSpace(splitToken[1])
	return tokenStr, nil
}

func extractRefreshToken(req *http.Request) (string, error) {
	cookie, err := req.Cookie("refreshToken")
	if err != nil {
		return "", errors.New("")
	}
	return cookie.Value, nil
}
