package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,alphanum"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Password      string             `json:"password,omitempty" bson:"password,omitempty" validate:"required"`
	IsAdmin       bool               `json:"isAdmin,omitempty" bson:"isAdmin,omitempty"`
	Phone         string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"omitempty,numeric"`
	Gender        string             `json:"gender,omitempty" bson:"gender,omitempty"`
	HowDidYouHear string             `json:"howDidYouHear,omitempty" bson:"howDidYouHear,omitempty"`
	City          string             `json:"city,omitempty" bson:"city,omitempty"`
	State         string             `json:"state,omitempty" bson:"state,omitempty"`
	CreatedAt     time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt     time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type ResponceUser struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,alphanum"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Key       string             `json:"key,omitempty" bson:"key,omitempty" validate:"required"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"omitempty,numeric"`
	Gender    string             `json:"gender,omitempty" bson:"gender,omitempty"`
	City      string             `json:"city,omitempty" bson:"city,omitempty"`
	State     string             `json:"state,omitempty" bson:"state,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
