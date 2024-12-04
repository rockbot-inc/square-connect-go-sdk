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

// Represents a [ListCustomerCustomAttributeDefinitions]($e/CustomerCustomAttributes/ListCustomerCustomAttributeDefinitions) response. Either `custom_attribute_definitions`, an empty object, or `errors` is present in the response. If additional results are available, the `cursor` field is also present along with `custom_attribute_definitions`.
type ListCustomerCustomAttributeDefinitionsResponse struct {
	// The retrieved custom attribute definitions. If no custom attribute definitions are found, Square returns an empty object (`{}`).
	CustomAttributeDefinitions []CustomAttributeDefinition `json:"custom_attribute_definitions,omitempty"`
	// The cursor to provide in your next call to this endpoint to retrieve the next page of results for your original request. This field is present only if the request succeeded and additional results are available. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
