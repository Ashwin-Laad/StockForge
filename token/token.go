package token

import (
	"context"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// SignedDetails represents the custom information stored inside the JWT.
type SignedDetails struct {
	Name  string `json:"name"`
	Email string `json:"email"`

	// StandardClaims contains the standard JWT fields such as
	// expiration time, issuer, etc.
	jwt.StandardClaims
}

// Secret key used to sign the JWT.
// It is loaded from the environment variable SECRET_KEY.
var secretKey = []byte(os.Getenv("SECRET_KEY"))

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

	// ---------------- Refresh Token ----------------

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

// ValidateToken verifies whether the JWT received from the client
// is genuine, has not expired, and was signed using our secret key.
//
// Parameters:
//
//	signedToken
//	    The JWT string received from the client.
//
// Returns:
//
//	claims
//	    The custom data stored inside the JWT (Name, Email, etc.)
//
//	message
//	    Error message if validation fails.
//
// Example:
//
//	Client sends:
//
//	Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5...
//
//	The middleware extracts:
//
//	eyJhbGciOiJIUzI1NiIsInR5...
//
//	and passes it into this function.
func ValidateToken(signedToken string) (*SignedDetails, string) {

	// Create an empty SignedDetails object.
	//
	// If the token is valid,
	// jwt.ParseWithClaims() will automatically fill this struct
	// with the information stored inside the JWT.
	claims := &SignedDetails{}

	// ParseWithClaims performs multiple operations:
	//
	// 1. Decodes the JWT.
	// 2. Verifies the digital signature.
	// 3. Checks whether the token format is valid.
	// 4. Stores the decoded payload inside "claims".
	//
	// Parameters:
	//
	// signedToken
	//     The JWT string received from the client.
	//
	// claims
	//     The struct where JWT data will be stored.
	//
	// callback function
	//     Returns the secret key that should be used
	//     to verify the token's signature.
	token, err := jwt.ParseWithClaims(

		signedToken,

		claims,

		func(token *jwt.Token) (interface{}, error) {

			// The JWT library calls this function internally
			// while verifying the token.
			//
			// It asks:
			//
			// "Which secret key should I use
			// to verify this token?"
			//
			// We return the same secret key that was used
			// while creating the JWT.
			return secretKey, nil
		},
	)

	// If parsing fails,
	// possible reasons are:
	//
	// • Token format is invalid.
	// • Signature has been modified.
	// • Wrong secret key.
	// • Token is corrupted.
	if err != nil {
		return nil, err.Error()
	}

	// Every JWT contains an expiration time.
	//
	// ExpiresAt is stored as a Unix timestamp.
	//
	// Example:
	//
	//	Token expires:
	//	26 July 2026
	//
	// Current time:
	//	27 July 2026
	//
	if !token.Valid {
		return nil, "invalid token"
	}
	// Even if parsing succeeds,
	// we must still verify that the JWT is considered valid.
	//
	// token.Valid checks:
	//
	// • Signature is correct.
	// • Expiration has not failed.
	// • Standard JWT validations passed.
	if !token.Valid {
		return nil, "invalid token"
	}

	// Everything passed.
	//
	// Return the decoded claims.
	//
	// Example:
	//
	// claims.Name
	// claims.Email
	//
	// can now be accessed by the middleware
	// or any protected route.
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
