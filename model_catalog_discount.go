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

// A discount applicable to items.
type CatalogDiscount struct {
	// The discount name. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points.
	Name string `json:"name,omitempty"`
	DiscountType *CatalogDiscountType `json:"discount_type,omitempty"`
	// The percentage of the discount as a string representation of a decimal number, using a `.` as the decimal separator and without a `%` sign. A value of `7.5` corresponds to `7.5%`. Specify a percentage of `0` if `discount_type` is `VARIABLE_PERCENTAGE`.  Do not use this field for amount-based or variable discounts.
	Percentage string `json:"percentage,omitempty"`
	AmountMoney *Money `json:"amount_money,omitempty"`
	// Indicates whether a mobile staff member needs to enter their PIN to apply the discount to a payment in the Square Point of Sale app.
	PinRequired bool `json:"pin_required,omitempty"`
	// The color of the discount display label in the Square Point of Sale app. This must be a valid hex color code.
	LabelColor string `json:"label_color,omitempty"`
	ModifyTaxBasis *CatalogDiscountModifyTaxBasis `json:"modify_tax_basis,omitempty"`
	MaximumAmountMoney *Money `json:"maximum_amount_money,omitempty"`
}
