package models

import (
	"time"

	"github.com/yaien/clothes-store-api/pkg/api/helpers/epayco"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Ref       string             `json:"ref"`
	Cart      *Cart              `json:"cart"`
	Shipping  *Shipping          `json:"shipping"`
	Status    InvoiceStatus      `json:"status"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	Payment   *epayco.Payment    `json:"-"`
	GuestID   primitive.ObjectID `json:"guestId"`
}

type InvoiceStatus string

const (
	Created   InvoiceStatus = "created"
	Accepted  InvoiceStatus = "accepted"
	Rejected  InvoiceStatus = "rejected"
	Pending   InvoiceStatus = "pending"
	Completed InvoiceStatus = "completed"
)
