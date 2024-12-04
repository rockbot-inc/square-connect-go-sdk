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

// Defines the fields that are included in the response body of a request to the [CreateCard]($e/Cards/CreateCard) endpoint.  Note: if there are errors processing the request, the card field will not be present.
type CreateCardResponse struct {
	// Errors resulting from the request.
	Errors []ModelError `json:"errors,omitempty"`
	Card *Card `json:"card,omitempty"`
}
