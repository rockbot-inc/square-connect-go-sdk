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

// Defines the response returned by [CreatePayment]($e/Payments/CreatePayment).  If there are errors processing the request, the `payment` field might not be present, or it might be present with a status of `FAILED`.
type CreatePaymentResponse struct {
	// Information about errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Payment *Payment `json:"payment,omitempty"`
}
