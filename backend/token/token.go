package token

import (
	"context"
	
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// SignedDetails represents the custom information stored inside the JWT.
// set with the header (context) hence can directly access
type SignedDetails struct {
	Name  string `json:"name"`
	Email string `json:"email"`

	// StandardClaims contains the standard JWT fields such as
	// expiration time, issuer, etc.
	jwt.StandardClaims
}

// // Secret key used to sign the JWT.
// // It is loaded from the environment variable SECRET_KEY.
var secretKey = []byte("my-secret-key")

// GenerateToken creates both an access token and a refresh token.
func GenerateToken(name, email string) (string, string, error) {

	// ---------------- Access Token ----------------

	accessClaims := &SignedDetails{
		Name:  name,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	accessToken, err := accessJWT.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	

	refreshClaims := &SignedDetails{
		Name:  name,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	refreshToken, err := refreshJWT.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}


func ValidateToken(signedToken string) (*SignedDetails, string) {

	
	claims := &SignedDetails{}

	
	token, err := jwt.ParseWithClaims(

		signedToken,

		claims,

		func(token *jwt.Token) (interface{}, error) {

			// The JWT library calls this function internally
			// while verifying the token.
			//
			
			return secretKey, nil
		},
	)

	if err != nil {
		return nil, err.Error()
	}


	if !token.Valid {
		return nil, "invalid token"
	}
	// Even if parsing succeeds,
	// we must still verify that the JWT is considered valid.

	if !token.Valid {
		return nil, "invalid token"
	}

	//
	// claims.Name
	// claims.Email

	return claims, ""
}

func UpdateAllTokens(
	userCollection *mongo.Collection,
	ctx context.Context,
	token string,
	refreshToken string,
	userID string,
) error {

	//----------------------------------------------------
	// Find the user whose user_id matches the supplied ID.
	//----------------------------------------------------

	filter := bson.M{
		"user_id": userID,
	}

	//----------------------------------------------------
	// $set updates only the specified fields.
	//
	// The remaining fields inside the document
	// remain unchanged.
	//----------------------------------------------------

	update := bson.M{
		"$set": bson.M{
			"token":         token,
			"refresh_token": refreshToken,
			"updated_at":    time.Now(),
		},
	}

	//----------------------------------------------------
	// UpdateOne searches for one matching document
	// and applies the update.
	//----------------------------------------------------

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		update,
	)

	return err
}
