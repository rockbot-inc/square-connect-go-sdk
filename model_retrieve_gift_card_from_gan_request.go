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

// A request to retrieve gift cards by their GANs.
type RetrieveGiftCardFromGanRequest struct {
	// The gift card account number (GAN) of the gift card to retrieve. The maximum length of a GAN is 255 digits to account for third-party GANs that have been imported. Square-issued gift cards have 16-digit GANs.
	Gan string `json:"gan"`
}