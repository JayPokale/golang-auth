package controllers

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jaypokale/golang-auth/config"
	"github.com/jaypokale/golang-auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func includeKey(user *models.User) (*models.ResponceUser, error) {
	responceUser := &models.ResponceUser{}
	record, _ := json.Marshal(user)
	json.Unmarshal(record, &responceUser)

	claims := jwt.MapClaims{
		"_id":      user.ID,
		"password": user.Password,
	}

	jwtKey := os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return nil, err
	}

	responceUser.Key = tokenString

	return responceUser, nil
}

func CreateUser(ctx *fiber.Ctx) error {
	userCollection := config.GetCollection("users")

	var user *models.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	user.IsAdmin = false

	existingUser := &models.User{}
	err := userCollection.FindOne(ctx.Context(), bson.M{"email": user.Email}).Decode(existingUser)
	if err == nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User already exists"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error hashing password"})
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	insertedResult, err := userCollection.InsertOne(ctx.Context(), user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating user"})
	}

	user.ID = insertedResult.InsertedID.(primitive.ObjectID)

	responceUser, err := includeKey(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User created but response generation failed"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(responceUser)
}

func LoginUser(ctx *fiber.Ctx) error {
	var user *models.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	userCollection := config.GetCollection("users")

	existingUser := &models.User{}
	err := userCollection.FindOne(ctx.Context(), bson.M{"email": user.Email}).Decode(existingUser)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error finding user"})
	}

	if existingUser == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Credentials"})
	}

	responceUser, err := includeKey(existingUser)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Oops, a server error occurred"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(responceUser)
}
