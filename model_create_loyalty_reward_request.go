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

// A request to create a loyalty reward.
type CreateLoyaltyRewardRequest struct {
	Reward *LoyaltyReward `json:"reward"`
	// A unique string that identifies this `CreateLoyaltyReward` request.  Keys can be any valid string, but must be unique for every request.
	IdempotencyKey string `json:"idempotency_key"`
}
