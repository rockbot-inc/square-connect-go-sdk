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

// Defines the fields that are included in the response body of a request to the [RetrieveTransaction](api-endpoint:Transactions-RetrieveTransaction) endpoint.  One of `errors` or `transaction` is present in a given response (never both).
type RetrieveTransactionResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Transaction *Transaction `json:"transaction,omitempty"`
}
