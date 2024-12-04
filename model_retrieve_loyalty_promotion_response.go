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

// Represents a [RetrieveLoyaltyPromotionPromotions]($e/Loyalty/RetrieveLoyaltyPromotion) response.
type RetrieveLoyaltyPromotionResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	LoyaltyPromotion *LoyaltyPromotion `json:"loyalty_promotion,omitempty"`
}
