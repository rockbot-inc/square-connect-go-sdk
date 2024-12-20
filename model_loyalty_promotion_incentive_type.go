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
// LoyaltyPromotionIncentiveType : Indicates the type of points incentive for a [loyalty promotion]($m/LoyaltyPromotion), which is used to determine how buyers can earn points from the promotion.
type LoyaltyPromotionIncentiveType string

// List of LoyaltyPromotionIncentiveType
const (
	MULTIPLIER_LoyaltyPromotionIncentiveType LoyaltyPromotionIncentiveType = "POINTS_MULTIPLIER"
	ADDITION_LoyaltyPromotionIncentiveType LoyaltyPromotionIncentiveType = "POINTS_ADDITION"
)
