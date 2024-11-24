package routes

import (
	"database/sql"
	"digital-wallet-api/internal/controllers"
	"digital-wallet-api/internal/middleware"
	"digital-wallet-api/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {

	// Initialize Services
	userService := service.NewUserService(db)
	walletService := service.NewWalletService(db)
	transactionService := service.NewTransactionService(db)
	groupWalletService := service.NewGroupWalletService(db)
	groupWalletMemberService := service.NewGroupWalletMemberService(db)

	// Authentication
	authController := controllers.NewAuthController(db)
	authRoutes := router.Group("/api/users")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.CreateUser)
	}

	// User
	userController := controllers.NewUserController(userService)
	userRoutes := router.Group("/api/users", middleware.JWTAuth())
	{
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	// Wallet
	walletController := controllers.NewWalletController(walletService)
	walletRoutes := router.Group("/api/wallet", middleware.JWTAuth())
	{
		walletRoutes.POST("/:id", walletController.CreateWallet)
		walletRoutes.GET("/:wallet_id/balance", walletController.CheckBalance)
		walletRoutes.GET("/:wallet_id/info", walletController.GetInfoByWalletId)
		walletRoutes.PUT("/:wallet_id/topup", walletController.TopUp)
		walletRoutes.PUT("/:wallet_id/freeze", walletController.FreezeWallet)
		walletRoutes.PUT("/:wallet_id/unfreeze", walletController.UnfreezeWallet)
	}

	// Transaction
	transactionController := controllers.NewTransactionController(transactionService)
	transactionRoutes := router.Group("/api/transactions", middleware.JWTAuth())
	{
		transactionRoutes.POST("/transfer", transactionController.Transfer)
		transactionRoutes.GET("/:wallet_id/income", transactionController.IncomeHistory)
		transactionRoutes.GET("/:wallet_id/expense", transactionController.ExpenseHistory)
		transactionRoutes.GET("/:wallet_id", transactionController.TransactionsHistory)
	}

	// GroupWallet
	groupWalletController := controllers.NewGroupWalletController(groupWalletService)
	groupWalletRoutes := router.Group("/api/gw", middleware.JWTAuth())
	{
		groupWalletRoutes.POST("/:wallet_id", groupWalletController.CreateGroupWallet)    // Create a new group wallet
		groupWalletRoutes.PUT("/:wallet_id", groupWalletController.UpdateGroupWalletGoal) // Update the goal for a group wallet
		groupWalletRoutes.GET("/:wallet_id", groupWalletController.GetGroupWalletByID)    // Get group wallet details by ID
		groupWalletRoutes.DELETE("/:wallet_id", groupWalletController.DeleteGroupWallet)  // Delete a group wallet
	}

	// GroupWalletMember
	groupWalletMemberController := controllers.NewGroupWalletMemberController(groupWalletMemberService)
	groupWalletMemberRoutes := router.Group("/api/gwm", middleware.JWTAuth())
	{
		groupWalletMemberRoutes.POST("/:wallet_id", groupWalletMemberController.AddMember)                          // Add a member to a group wallet
		groupWalletMemberRoutes.PUT("/:wallet_id/:member_id", groupWalletMemberController.UpdateMemberContribution) // Update a member's contribution
		groupWalletMemberRoutes.GET("/:wallet_id", groupWalletMemberController.GetAllMembers)                       // Get all members of a group wallet
		groupWalletMemberRoutes.DELETE("/:wallet_id/:member_id", groupWalletMemberController.RemoveMember)          // Remove a member from a group wallet
	}

}
