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

// The creation source filter.  If one or more creation sources are set, customer profiles are included in, or excluded from, the result if they match at least one of the filter criteria.
type CustomerCreationSourceFilter struct {
	// The list of creation sources used as filtering criteria. See [CustomerCreationSource](#type-customercreationsource) for possible values
	Values []CustomerCreationSource `json:"values,omitempty"`
	Rule *CustomerInclusionExclusion `json:"rule,omitempty"`
}
