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

type DestinationDetailsCardRefundDetails struct {
	Card *Card `json:"card,omitempty"`
	// The method used to enter the card's details for the refund. The method can be `KEYED`, `SWIPED`, `EMV`, `ON_FILE`, or `CONTACTLESS`.
	EntryMethod string `json:"entry_method,omitempty"`
	// The authorization code provided by the issuer when a refund is approved.
	AuthResultCode string `json:"auth_result_code,omitempty"`
}