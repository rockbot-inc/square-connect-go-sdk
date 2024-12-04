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

// Defines the fields included in the response body from the [BulkCreateCustomers]($e/Customers/BulkCreateCustomers) endpoint.
type BulkCreateCustomersResponse struct {
	// A map of responses that correspond to individual create requests, represented by key-value pairs.  Each key is the idempotency key that was provided for a create request and each value is the corresponding response. If the request succeeds, the value is the new customer profile. If the request fails, the value contains any errors that occurred during the request.
	Responses map[string]CreateCustomerResponse `json:"responses,omitempty"`
	// Any top-level errors that prevented the bulk operation from running.
	Errors []ModelError `json:"errors,omitempty"`
}
