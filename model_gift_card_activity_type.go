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
// GiftCardActivityType : Indicates the type of [gift card activity]($m/GiftCardActivity).
type GiftCardActivityType string

// List of GiftCardActivityType
const (
	ACTIVATE_GiftCardActivityType GiftCardActivityType = "ACTIVATE"
	LOAD_GiftCardActivityType GiftCardActivityType = "LOAD"
	REDEEM_GiftCardActivityType GiftCardActivityType = "REDEEM"
	CLEAR_BALANCE_GiftCardActivityType GiftCardActivityType = "CLEAR_BALANCE"
	DEACTIVATE_GiftCardActivityType GiftCardActivityType = "DEACTIVATE"
	ADJUST_INCREMENT_GiftCardActivityType GiftCardActivityType = "ADJUST_INCREMENT"
	ADJUST_DECREMENT_GiftCardActivityType GiftCardActivityType = "ADJUST_DECREMENT"
	REFUND_GiftCardActivityType GiftCardActivityType = "REFUND"
	UNLINKED_ACTIVITY_REFUND_GiftCardActivityType GiftCardActivityType = "UNLINKED_ACTIVITY_REFUND"
	IMPORT__GiftCardActivityType GiftCardActivityType = "IMPORT"
	BLOCK_GiftCardActivityType GiftCardActivityType = "BLOCK"
	UNBLOCK_GiftCardActivityType GiftCardActivityType = "UNBLOCK"
	IMPORT_REVERSAL_GiftCardActivityType GiftCardActivityType = "IMPORT_REVERSAL"
	TRANSFER_BALANCE_FROM_GiftCardActivityType GiftCardActivityType = "TRANSFER_BALANCE_FROM"
	TRANSFER_BALANCE_TO_GiftCardActivityType GiftCardActivityType = "TRANSFER_BALANCE_TO"
)
