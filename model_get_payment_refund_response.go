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

// Defines the response returned by [GetRefund]($e/Refunds/GetPaymentRefund).  Note: If there are errors processing the request, the refund field might not be present or it might be present in a FAILED state.
type GetPaymentRefundResponse struct {
	// Information about errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Refund *PaymentRefund `json:"refund,omitempty"`
}
