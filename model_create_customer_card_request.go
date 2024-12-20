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

// Defines the fields that are included in the request body of a request to the `CreateCustomerCard` endpoint.
type CreateCustomerCardRequest struct {
	// A card nonce representing the credit card to link to the customer.  Card nonces are generated by the Square payment form when customers enter their card information. For more information, see [Walkthrough: Integrate Square Payments in a Website](https://developer.squareup.com/docs/web-payments/take-card-payment).  __NOTE:__ Card nonces generated by digital wallets (such as Apple Pay) cannot be used to create a customer card.
	CardNonce string `json:"card_nonce"`
	BillingAddress *Address `json:"billing_address,omitempty"`
	// The full name printed on the credit card.
	CardholderName string `json:"cardholder_name,omitempty"`
	// An identifying token generated by [Payments.verifyBuyer()](https://developer.squareup.com/reference/sdks/web/payments/objects/Payments#Payments.verifyBuyer). Verification tokens encapsulate customer device information and 3-D Secure challenge results to indicate that Square has verified the buyer identity.
	VerificationToken string `json:"verification_token,omitempty"`
}
