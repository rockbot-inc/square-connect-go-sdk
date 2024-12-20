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
// GiftCardActivityDeactivateReason : Indicates the reason for deactivating a [gift card]($m/GiftCard).
type GiftCardActivityDeactivateReason string

// List of GiftCardActivityDeactivateReason
const (
	SUSPICIOUS_ACTIVITY_GiftCardActivityDeactivateReason GiftCardActivityDeactivateReason = "SUSPICIOUS_ACTIVITY"
	UNKNOWN_REASON_GiftCardActivityDeactivateReason GiftCardActivityDeactivateReason = "UNKNOWN_REASON"
	CHARGEBACK_DEACTIVATE_GiftCardActivityDeactivateReason GiftCardActivityDeactivateReason = "CHARGEBACK_DEACTIVATE"
)
