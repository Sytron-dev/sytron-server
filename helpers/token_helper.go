package helper

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "sytron-server/database"

    "github.com/golang-jwt/jwt/v5"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// SignedDetails
type SignedDetails struct {
    Email      string
    First_name string
    Last_name  string
    Uid        string
    jwt.RegisteredClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateAllTokens generates both teh detailed token and refresh token
func GenerateAllTokens(email string, firstName string, lastName string, uid string) (signedToken string, signedRefreshToken string, err error) {
    // Create claims for access token
    nowTime := time.Now()
    expireTime := nowTime.Add(time.Minute * 15)
    claims := SignedDetails{
        Email:      email,
        First_name: firstName,
        Last_name:  lastName,
        Uid:        uid,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expireTime),
        },
        //     ExpiresAt: &jwt.NumericDate{now.Add(time.Minute * 15)},
        // },
    }

    // Create claims for refresh token

    refreshClaims := &SignedDetails{
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expireTime),
        },
    }
    


    // Generate access token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err = token.SignedString([]byte(SECRET_KEY))
    if err != nil {
        return "", "", err
    }

    // Generate refresh token
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
    signedRefreshToken, err = refreshToken.SignedString([]byte(SECRET_KEY))
    if err != nil {
        return "", "", err
    }

    return signedToken, signedRefreshToken, nil
}

//ValidateToken validates the jwt token
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
    // nowTime := time.Now()
    
    token, err := jwt.ParseWithClaims(
        signedToken,
        &SignedDetails{},
        func(token *jwt.Token) (interface{}, error) {
            return []byte(SECRET_KEY), nil
        },
    )

    if err != nil {
        msg = err.Error()
        return
    }

    claims, ok := token.Claims.(*SignedDetails)
    if !ok {
        msg = fmt.Sprintf("the token is invalid")
        msg = err.Error()
        return
    }

    // if claims.ExpiresAt < time.Now().Local().Unix() {
    //     msg = fmt.Sprintf("token is expired")
    //     msg = err.Error()
    //     return
    // }
    // if claims.ExpiresAt < jwt.NewNumericDate(nowTime) {
    //     msg = fmt.Sprintf("token is expired")
    //     msg = err.Error()
    //     return
    // }

    return claims, msg
}

//UpdateAllTokens renews the user tokens when they login
func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
    var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

    var updateObj primitive.D

    updateObj = append(updateObj, bson.E{"token", signedToken})
    updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

    Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
    updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

    upsert := true
    filter := bson.M{"user_id": userId}
    opt := options.UpdateOptions{
        Upsert: &upsert,
    }

    _, err := userCollection.UpdateOne(
        ctx,
        filter,
        bson.D{
            {"$set", updateObj},
        },
        &opt,
    )
    defer cancel()

    if err != nil {
        log.Panic(err)
        return
    }

    return
}