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

// A response that contains a `GiftCard`. The response might contain a set of `Error` objects if the request resulted in errors.
type CreateGiftCardResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	GiftCard *GiftCard `json:"gift_card,omitempty"`
}
