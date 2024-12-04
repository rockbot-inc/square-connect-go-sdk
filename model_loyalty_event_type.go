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
// LoyaltyEventType : The type of the loyalty event.
type LoyaltyEventType string

// List of LoyaltyEventType
const (
	ACCUMULATE_POINTS_LoyaltyEventType LoyaltyEventType = "ACCUMULATE_POINTS"
	CREATE_REWARD_LoyaltyEventType LoyaltyEventType = "CREATE_REWARD"
	REDEEM_REWARD_LoyaltyEventType LoyaltyEventType = "REDEEM_REWARD"
	DELETE_REWARD_LoyaltyEventType LoyaltyEventType = "DELETE_REWARD"
	ADJUST_POINTS_LoyaltyEventType LoyaltyEventType = "ADJUST_POINTS"
	EXPIRE_POINTS_LoyaltyEventType LoyaltyEventType = "EXPIRE_POINTS"
	OTHER_LoyaltyEventType LoyaltyEventType = "OTHER"
	ACCUMULATE_PROMOTION_POINTS_LoyaltyEventType LoyaltyEventType = "ACCUMULATE_PROMOTION_POINTS"
)
