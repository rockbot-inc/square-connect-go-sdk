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

// Lists all event types that can be subscribed to.
type ListEventTypesRequest struct {
	// The API version for which to list event types. Setting this field overrides the default version used by the application.
	ApiVersion string `json:"api_version,omitempty"`
}
