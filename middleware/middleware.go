package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jaypokale/golang-auth/config"
	"github.com/jaypokale/golang-auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getUserFromToken(ctx *fiber.Ctx) (*models.User, error) {
	tokenString := ctx.Get("Authorization")
	if tokenString == "" {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "No token provided")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	ID := claims["_id"].(string)
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Not a valid request")
	}

	user := &models.User{}
	users := config.GetCollection("users")
	err = users.FindOne(ctx.Context(), bson.M{"_id": objectID}).Decode(user)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Not a valid request")
	}

	if user == nil || user.Password != claims["password"].(string) {
		return nil, fiber.NewError(fiber.StatusForbidden, "Not a valid user")
	}

	return user, nil
}

func VerifyUser(ctx *fiber.Ctx) error {
	user, err := getUserFromToken(ctx)
	if err != nil {
		return err
	}

	ctx.Locals("id", user.ID)
	return ctx.Next()
}

func VerifyAdmin(ctx *fiber.Ctx) error {
	user, err := getUserFromToken(ctx)
	if err != nil {
		return err
	}

	if user.IsAdmin {
		ctx.Locals("id", ctx.Params("id"))
		return ctx.Next()
	}

	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Admin access required"})
}
