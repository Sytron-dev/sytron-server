package routes

import (
	"github.com/gofiber/fiber/v2"

	assetController "sytron-server/api/handlers/assets"
	controller "sytron-server/api/handlers/service"
	"sytron-server/constants"
)

func initServicesRoutes(router *fiber.App) {
	// bed & breakfast
	router.Get("/services/bnb", controller.GetBNBs())
	router.Get("/services/bnb/:id", controller.GetSingleBNB()) // also gets assets and prices
}

func initProtectedServicesRoutes(router *fiber.App) {
	router.Post("/services/bnb", controller.CreateBNB())
	router.Post("/services/bnb/:id/image", controller.UploadBNBImage())
	router.Put("/services/bnb/:id", controller.UpdateBNB())
	router.Delete("/services/bnb/:id", controller.DeleteBNB())

	// asset management

	router.Post(
		"/services/bnb/:id/assets",
		assetController.CreateAsset(constants.BNB_COLLECTION_NAME, constants.BNB_STORAGE_DIR),
	)
	router.Put(
		"/services/bnb/:ref_id/assets/:id",
		assetController.UpdateAsset(constants.BNB_COLLECTION_NAME, constants.BNB_STORAGE_DIR),
	)
}
