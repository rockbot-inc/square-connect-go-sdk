/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package square

// Represents a [CreateLoyaltyPromotion]($e/Loyalty/CreateLoyaltyPromotion) request.
type CreateLoyaltyPromotionRequest struct {
	LoyaltyPromotion *LoyaltyPromotion `json:"loyalty_promotion"`
	// A unique identifier for this request, which is used to ensure idempotency. For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey string `json:"idempotency_key"`
}
