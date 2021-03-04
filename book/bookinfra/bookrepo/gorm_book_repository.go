package bookrepo

// CreateProduct create a new product in the database
// func CreateProduct(db string) {
// 	return func(c *fiber.Ctx) error {
// 	product := new(BookRepo)
// 	db := db.Get()

// 	if err := c.BodyParser(product); err != nil {
// 		return fiber.ErrBadRequest
// 	}

// 	db.Create(product)

// 	return c.JSON(product)
// }
// }

// UpdateProduct is a handler for update product by id
// func UpdateProduct(c *fiber.Ctx) error {
// 	var body map[string]interface{}
// 	var product BookRepo
// 	db := db.Get()

// 	result := db.Find(&product, c.Params("id"))
// 	if result.RowsAffected == 0 {
// 		return fiber.ErrNotFound
// 	}

// 	if err := c.BodyParser(&body); err != nil {
// 		return fiber.ErrBadRequest
// 	}

// 	db.Model(&product).Updates(body)

// 	return c.JSON(product)
// }

// // DeleteProduct is a handler for delete product by id
// func DeleteProduct(c *fiber.Ctx) error {

// 	db := db.Get()
// 	db.Delete(&BookRepo{}, c.Params("id"))

// 	return c.JSON(&fiber.Map{
// 		"message": "Product Deleted!",
// 	})
// }
