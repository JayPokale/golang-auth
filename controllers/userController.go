package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jaypokale/golang-auth/config"
	"github.com/jaypokale/golang-auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getUserID(ctx *fiber.Ctx) (primitive.ObjectID, error) {
	userID, ok := ctx.Locals("id").(string)
	if !ok {
		return primitive.NilObjectID, fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}
	return primitive.ObjectIDFromHex(userID)
}

func fetchUserByID(ctx *fiber.Ctx, userID primitive.ObjectID) (*models.ResponceUser, error) {
	users := config.GetCollection("users")

	user := models.ResponceUser{}
	err := users.FindOne(ctx.Context(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Error fetching user")
	}
	return &user, nil
}

func updateUserByID(ctx *fiber.Ctx, userID primitive.ObjectID, updateFields map[string]interface{}) (*models.ResponceUser, error) {
	users := config.GetCollection("users")

	update := bson.M{"$set": updateFields}

	result, err := users.UpdateOne(ctx.Context(), bson.M{"_id": userID}, update)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Error updating user")
	}

	if result.ModifiedCount == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	var updatedUser models.ResponceUser
	err = users.FindOne(ctx.Context(), bson.M{"_id": userID}).Decode(&updatedUser)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Error fetching updated user")
	}

	return &updatedUser, nil
}

func deleteUserByID(ctx *fiber.Ctx, userID primitive.ObjectID) error {
	users := config.GetCollection("users")

	result, err := users.DeleteOne(ctx.Context(), bson.M{"_id": userID})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error deleting user")
	}

	if result.DeletedCount == 0 {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return nil
}

// Handlers

func GetUsers(ctx *fiber.Ctx) error {
	users := config.GetCollection("users")
	cursor, err := users.Find(ctx.Context(), bson.D{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching users"})
	}

	var response []models.ResponceUser

	for cursor.Next(ctx.Context()) {
		user := models.ResponceUser{}
		if err := cursor.Decode(&user); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching users"})
		}
		response = append(response, user)
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func GetUserByID(ctx *fiber.Ctx) error {
	userID, err := getUserID(ctx)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := fetchUserByID(ctx, userID)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func UpdateUserByID(ctx *fiber.Ctx) error {
	userID, err := getUserID(ctx)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(fiber.Map{"error": err.Error()})
	}

	var requestBody map[string]interface{}
	if err := ctx.BodyParser(&requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	delete(requestBody, "password")
	delete(requestBody, "isAdmin")
	requestBody["UpdatedAt"] = time.Now()

	updatedUser, err := updateUserByID(ctx, userID, requestBody)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(updatedUser)
}

func DeleteUserByID(ctx *fiber.Ctx) error {
	userID, err := getUserID(ctx)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(fiber.Map{"error": err.Error()})
	}

	err = deleteUserByID(ctx, userID)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}
