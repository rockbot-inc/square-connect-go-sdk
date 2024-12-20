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

// Describes query filters to apply.
type InvoiceFilter struct {
	// Limits the search to the specified locations. A location is required.  In the current implementation, only one location can be specified.
	LocationIds []string `json:"location_ids"`
	// Limits the search to the specified customers, within the specified locations.  Specifying a customer is optional. In the current implementation,  a maximum of one customer can be specified.
	CustomerIds []string `json:"customer_ids,omitempty"`
}
