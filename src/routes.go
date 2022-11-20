package src

import (
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/labstack/echo/v4"
)

// API Routes
func Routes(e *echo.Echo, handler *handler.Handler) {
	e.GET("/", handler.RootHandler)

	e.GET("/login", handler.LoginHandler)
	e.GET("/login/url", handler.LoginURLHandler)

	e.GET("/logout", handler.LogoutHandler)

	// Login Page
	e.GET("/user/me", handler.MeHandler)

	// update my profile
	e.PUT("/user", handler.UpdateUserHandler)

	// Get all users
	e.GET("/users", handler.AllUsersHandler)
	// Set select user
	e.POST("/user/select", handler.ChangeSelect)

	// bio
	e.GET("/user/bio", handler.BioHandler)
	e.POST("/user/bio", handler.CreateBioHandler)
	e.PUT("/user/bio", handler.UpdateBioHandler)
	e.DELETE("/user/bio", handler.DeleteBioHandler)

	// location
	e.GET("/user/location", handler.LocationHandler)
	e.POST("/user/location", handler.CreateLocationHandler)
	e.PUT("/user/location", handler.UpdateLocationHandler)
	e.DELETE("/user/location", handler.DeleteLocationHandler)

	// product
	e.GET("/user/product", handler.ProductHandler)
	e.POST("/user/product", handler.CreateProductHandler)
	e.PUT("/user/product", handler.UpdateProductHandler)
	e.DELETE("/user/product", handler.DeleteProductHandler)

	// link
	e.GET("/user/link", handler.LinkHandler)
	e.POST("/user/link", handler.CreateLinkHandler)
	e.PUT("/user/link", handler.UpdateLinkHandler)
	e.DELETE("/user/link", handler.DeleteLinkHandler)

	// Category
	e.GET("/user/category", handler.CategoryHandler)
	e.POST("/user/category", handler.CreateCategoryHandler)
	e.PUT("/user/category", handler.UpdateCategoryHandler)
	e.DELETE("/user/category", handler.DeleteCategoryHandler)

	// Notice
	e.GET("/user/notice", handler.NoticeHandler)
	e.PUT("/user/notice", handler.UpdateNoticeHandler)

	// Contacts
	e.GET("/user/contact", handler.ContactGetHandler)
	e.DELETE("/user/contact", handler.ContactDeleteHandler)

	// Public endpoints
	e.GET("/public/profile", handler.PublicProfileHandler)
	e.GET("/public/product", handler.PublicProductsHandler)

	e.POST("/contact", handler.ContactHandler)
}
