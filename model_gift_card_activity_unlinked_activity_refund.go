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

// Represents details about an `UNLINKED_ACTIVITY_REFUND` [gift card activity type]($m/GiftCardActivityType).
type GiftCardActivityUnlinkedActivityRefund struct {
	AmountMoney *Money `json:"amount_money"`
	// A client-specified ID that associates the gift card activity with an entity in another system.
	ReferenceId string `json:"reference_id,omitempty"`
	// The ID of the refunded payment. This field is not used starting in Square version 2022-06-16.
	PaymentId string `json:"payment_id,omitempty"`
}
