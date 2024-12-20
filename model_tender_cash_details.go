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

// Represents the details of a tender with `type` `CASH`.
type TenderCashDetails struct {
	BuyerTenderedMoney *Money `json:"buyer_tendered_money,omitempty"`
	ChangeBackMoney *Money `json:"change_back_money,omitempty"`
}
