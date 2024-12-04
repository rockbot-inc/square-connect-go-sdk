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
// GiftCardActivityRedeemStatus : Indicates the status of a [gift card]($m/GiftCard) redemption. This status is relevant only for redemptions made from Square products (such as Square Point of Sale) because Square products use a  two-state process. Gift cards redeemed using the Gift Card Activities API always have a `COMPLETED` status.
type GiftCardActivityRedeemStatus string

// List of GiftCardActivityRedeemStatus
const (
	PENDING_GiftCardActivityRedeemStatus GiftCardActivityRedeemStatus = "PENDING"
	COMPLETED_GiftCardActivityRedeemStatus GiftCardActivityRedeemStatus = "COMPLETED"
	CANCELED_GiftCardActivityRedeemStatus GiftCardActivityRedeemStatus = "CANCELED"
)
