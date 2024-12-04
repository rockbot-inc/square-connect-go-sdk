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

// Represents a Square gift card.
type GiftCard struct {
	// The Square-assigned ID of the gift card.
	Id string `json:"id,omitempty"`
	Type_ *GiftCardType `json:"type"`
	GanSource *GiftCardGanSource `json:"gan_source,omitempty"`
	State *GiftCardStatus `json:"state,omitempty"`
	BalanceMoney *Money `json:"balance_money,omitempty"`
	// The gift card account number (GAN). Buyers can use the GAN to make purchases or check  the gift card balance.
	Gan string `json:"gan,omitempty"`
	// The timestamp when the gift card was created, in RFC 3339 format.  In the case of a digital gift card, it is the time when you create a card  (using the Square Point of Sale application, Seller Dashboard, or Gift Cards API).   In the case of a plastic gift card, it is the time when Square associates the card with the  seller at the time of activation.
	CreatedAt string `json:"created_at,omitempty"`
	// The IDs of the [customer profiles](entity:Customer) to whom this gift card is linked.
	CustomerIds []string `json:"customer_ids,omitempty"`
}
