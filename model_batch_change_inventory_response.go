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

type BatchChangeInventoryResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The current counts for all objects referenced in the request.
	Counts []InventoryCount `json:"counts,omitempty"`
	// Changes created for the request.
	Changes []InventoryChange `json:"changes,omitempty"`
}
