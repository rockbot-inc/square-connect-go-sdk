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

// Represents a [RetrieveLocationCustomAttribute]($e/LocationCustomAttributes/RetrieveLocationCustomAttribute) response. Either `custom_attribute_definition` or `errors` is present in the response.
type RetrieveLocationCustomAttributeResponse struct {
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
