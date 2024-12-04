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

type PaymentBalanceActivityFeeDetail struct {
	// The ID of the payment associated with this activity This will only be populated when a principal LedgerEntryToken is also populated. If the fee is independent (there is no principal LedgerEntryToken) then this will likely not be populated.
	PaymentId string `json:"payment_id,omitempty"`
}
