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

// The query filter to return the search result by exact match of the specified attribute name and value.
type CatalogQueryExact struct {
	// The name of the attribute to be searched. Matching of the attribute name is exact.
	AttributeName string `json:"attribute_name"`
	// The desired value of the search attribute. Matching of the attribute value is case insensitive and can be partial. For example, if a specified value of \"sma\", objects with the named attribute value of \"Small\", \"small\" are both matched.
	AttributeValue string `json:"attribute_value"`
}