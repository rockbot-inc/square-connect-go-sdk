/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Represents a [CreateGiftCard]($e/GiftCards/CreateGiftCard) request.
type CreateGiftCardRequest struct {
	// A unique identifier for this request, used to ensure idempotency. For more information,  see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey string `json:"idempotency_key"`
	// The ID of the [location](entity:Location) where the gift card should be registered for  reporting purposes. Gift cards can be redeemed at any of the seller's locations.
	LocationId string `json:"location_id"`
	GiftCard *GiftCard `json:"gift_card"`
}
