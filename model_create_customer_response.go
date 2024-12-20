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

// Defines the fields that are included in the response body of a request to the [CreateCustomer]($e/Customers/CreateCustomer) or [BulkCreateCustomers]($e/Customers/BulkCreateCustomers) endpoint.  Either `errors` or `customer` is present in a given response (never both).
type CreateCustomerResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
}
