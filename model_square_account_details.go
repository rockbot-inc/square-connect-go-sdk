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

// Additional details about Square Account payments.
type SquareAccountDetails struct {
	// Unique identifier for the payment source used for this payment.
	PaymentSourceToken string `json:"payment_source_token,omitempty"`
	// Information about errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
}