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

type CreateCatalogImageRequest struct {
	// A unique string that identifies this CreateCatalogImage request. Keys can be any valid string but must be unique for every CreateCatalogImage request.  See [Idempotency keys](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) for more information.
	IdempotencyKey string `json:"idempotency_key"`
	// Unique ID of the `CatalogObject` to attach this `CatalogImage` object to. Leave this field empty to create unattached images, for example if you are building an integration where an image can be attached to catalog items at a later time.
	ObjectId string `json:"object_id,omitempty"`
	Image *CatalogObject `json:"image"`
	// If this is set to `true`, the image created will be the primary, or first image of the object referenced by `object_id`. If the `CatalogObject` already has a primary `CatalogImage`, setting this field to `true` will replace the primary image. If this is set to `false` and you use the Square API version 2021-12-15 or later, the image id will be appended to the list of `image_ids` on the object.  With Square API version 2021-12-15 or later, the default value is `false`. Otherwise, the effective default value is `true`.
	IsPrimary bool `json:"is_primary,omitempty"`
}
