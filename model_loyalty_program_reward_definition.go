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

// Provides details about the reward tier discount. DEPRECATED at version 2020-12-16. Discount details are now defined using a catalog pricing rule and other catalog objects. For more information, see [Getting discount details for a reward tier](https://developer.squareup.com/docs/loyalty-api/loyalty-rewards#get-discount-details).
type LoyaltyProgramRewardDefinition struct {
	Scope *LoyaltyProgramRewardDefinitionScope `json:"scope"`
	DiscountType *LoyaltyProgramRewardDefinitionType `json:"discount_type"`
	// The fixed percentage of the discount. Present if `discount_type` is `FIXED_PERCENTAGE`. For example, a 7.25% off discount will be represented as \"7.25\". DEPRECATED at version 2020-12-16. You can find this information in the `discount_data.percentage` field of the `DISCOUNT` catalog object referenced by the pricing rule.
	PercentageDiscount string `json:"percentage_discount,omitempty"`
	// The list of catalog objects to which this reward can be applied. They are either all item-variation ids or category ids, depending on the `type` field. DEPRECATED at version 2020-12-16. You can find this information in the `product_set_data.product_ids_any` field of the `PRODUCT_SET` catalog object referenced by the pricing rule.
	CatalogObjectIds []string `json:"catalog_object_ids,omitempty"`
	FixedDiscountMoney *Money `json:"fixed_discount_money,omitempty"`
	MaxDiscountMoney *Money `json:"max_discount_money,omitempty"`
}
