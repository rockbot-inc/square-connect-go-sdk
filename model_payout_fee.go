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

// Represents a payout fee that can incur as part of a payout.
type PayoutFee struct {
	AmountMoney *Money `json:"amount_money,omitempty"`
	// The timestamp of when the fee takes effect, in RFC 3339 format.
	EffectiveAt string `json:"effective_at,omitempty"`
	Type_ *PayoutFeeType `json:"type,omitempty"`
}
