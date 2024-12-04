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

// Contains query criteria for the search.
type SearchOrdersQuery struct {
	Filter *SearchOrdersFilter `json:"filter,omitempty"`
	Sort *SearchOrdersSort `json:"sort,omitempty"`
}
