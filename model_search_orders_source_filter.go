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

// A filter based on order `source` information.
type SearchOrdersSourceFilter struct {
	// Filters by the [Source](entity:OrderSource) `name`. The filter returns any orders with a `source.name` that matches any of the listed source names.  Max: 10 source names.
	SourceNames []string `json:"source_names,omitempty"`
}
