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
// SubscriptionPricingType : Determines the pricing of a [Subscription]($m/Subscription)
type SubscriptionPricingType string

// List of SubscriptionPricingType
const (
	STATIC_SubscriptionPricingType SubscriptionPricingType = "STATIC"
	RELATIVE_SubscriptionPricingType SubscriptionPricingType = "RELATIVE"
)
