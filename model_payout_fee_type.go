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
// PayoutFeeType : Represents the type of payout fee that can incur as part of a payout.
type PayoutFeeType string

// List of PayoutFeeType
const (
	TRANSFER_FEE_PayoutFeeType PayoutFeeType = "TRANSFER_FEE"
	TAX_ON_TRANSFER_FEE_PayoutFeeType PayoutFeeType = "TAX_ON_TRANSFER_FEE"
)
