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

// Defines the body parameters that can be included in requests to the [BulkDeleteCustomers]($e/Customers/BulkDeleteCustomers) endpoint.
type BulkDeleteCustomersRequest struct {
	// The IDs of the [customer profiles](entity:Customer) to delete.
	CustomerIds []string `json:"customer_ids"`
}
