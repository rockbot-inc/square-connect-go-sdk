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

// Describes the pricing for the subscription.
type SubscriptionPricing struct {
	Type_ *SubscriptionPricingType `json:"type,omitempty"`
	// The ids of the discount catalog objects
	DiscountIds []string `json:"discount_ids,omitempty"`
	PriceMoney *Money `json:"price_money,omitempty"`
}
