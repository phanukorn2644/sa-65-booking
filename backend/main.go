package main

import (
	"github.com/phanukorn2644/sa-65-example/controller"
	"github.com/phanukorn2644/sa-65-example/middlewares"

	"github.com/phanukorn2644/sa-65-example/entity"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {

	entity.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{

			// User

			r.GET("/Bookings", controller.ListBookings)
			r.GET("/Booking/:id", controller.GetBooking)
			r.POST("/Bookings", controller.CreateBooking)
			r.PATCH("/Bookings", controller.UpdateBooking)
			r.DELETE("/Bookings/:id", controller.DeleteBooking)
			///
			r.GET("/Times", controller.ListTimes)
			r.GET("/Time/:id", controller.GetTime)
			r.POST("/Times", controller.CreateTime)
			r.PATCH("/Times", controller.UpdateTime)
			r.DELETE("/Times/:id", controller.DeleteTime)
			///
			r.GET("/Students", controller.ListStudent)
			r.GET("/Student/:id", controller.GetStudent)
			r.POST("/Students", controller.CreateStudent)
			// r.PATCH("/Students", controller.UpdateStudent)
			// r.DELETE("/Students/:id", controller.DeleteStudent)
			//user routes
			router.GET("/genders", controller.ListGenders)
			router.GET("/gender/:id", controller.GetGender)
			router.POST("/genders", controller.CreateGender)
			router.PATCH("/genders", controller.UpdateGender)
			router.DELETE("/genders/:id", controller.DeleteGender)

			//memberClass routes
			router.GET("/job_positions", controller.ListJob_Position)
			router.GET("/job_position/:id", controller.GetJob_Position)
			router.POST("/job_positions", controller.CreateJob_Position)
			router.PATCH("/job_positions", controller.UpdateJob_Position)
			router.DELETE("/job_positions/:id", controller.DeleteJob_Position)

			//province routes
			router.GET("/provinces", controller.ListProvince)
			router.GET("/province/:id", controller.GetProvince)
			router.POST("/provinces", controller.CreateProvince)
			router.PATCH("/provinces", controller.UpdateProvince)
			router.DELETE("/provinces/:id", controller.DeleteProvince)

			//role routes
			router.GET("/employees", controller.ListEmployee)
			router.GET("/employee/:id", controller.GetEmployee)
			router.POST("/employees", controller.CreateEmployee)
			router.PATCH("/employees", controller.UpdateEmployee)
			router.DELETE("/employees/:id", controller.DeleteEmployee)

			// Run the server
			r.GET("/Room_types", controller.ListRoom_types)
			r.GET("/Room_type/:id", controller.GetRoom_type)
			r.POST("/Room_types", controller.CreateRoom_type)
			r.PATCH("/Room_types", controller.UpdateRoom_type)
			r.DELETE("/Room_types/:id", controller.DeleteRoom_type)
			////
			r.GET("/Room_prices", controller.ListRoom_prices)
			r.GET("/Room_price/:id", controller.GetRoom_price)
			r.POST("/Room_prices", controller.CreateRoom_price)
			r.PATCH("/Room_prices", controller.UpdateRoom_price)
			r.DELETE("/Room_prices/:id", controller.DeleteRoom_price)
			////
			r.GET("/Set_of_furnitures", controller.ListSet_of_furnitures)
			r.GET("/Set_of_furniture/:id", controller.GetSet_of_furniture)
			r.POST("/Set_of_furnitures", controller.CreateSet_of_furniture)
			r.PATCH("/Set_of_furnitures", controller.UpdateSet_of_furniture)
			r.DELETE("/Set_of_furnitures/:id", controller.DeleteSet_of_furniture)
			////
			r.GET("/Rooms", controller.ListRoom)
			r.GET("/Room/:id", controller.GetRoom)
			r.POST("/Rooms", controller.CreateRoom)
			r.PATCH("/Rooms", controller.UpdateRoom)
			r.DELETE("/Rooms/:id", controller.DeleteRoom)

			r.GET("/Furnitures", controller.ListFurniture)
			r.GET("/Furniture/:id", controller.GetFurniture)
			r.POST("/Furnitures", controller.CreateFurniture)
			r.PATCH("/Furnitures", controller.UpdateFurniture)
			r.DELETE("/Furnitures/:id", controller.DeleteFurniture)

			r.GET("/Repairs", controller.ListRepairs)
			r.GET("/Repair/:id", controller.GetRepair)
			r.POST("/Repairs", controller.CreateRepair)

			router.POST("/semesters", controller.CreateSemester)
			router.GET("/semester/:id", controller.GetSemester)
			router.GET("/semesters", controller.ListSemesters)
			router.DELETE("/semesters/:id", controller.DeleteSemester)
			router.PATCH("/semesters", controller.UpdateSemester)

			// Payment_Bill Routes
			router.POST("/payment_bills", controller.CreatePayment_Bill)
			router.GET("/payment_bills", controller.ListPayment_Bills)
			router.GET("/payment_bill/:id", controller.GetPayment_Bill)
			router.PATCH("/payment_bills", controller.UpdatePayment_Bill)
			router.DELETE("/payment_bills/:id", controller.DeletePayment_Bill)

			// Role Routes
			router.GET("/roles", controller.ListRole)
			router.GET("/role/:id", controller.GetRole)
			router.POST("/roles", controller.CreateRole)
			router.PATCH("/roles", controller.UpdateRole)
			router.DELETE("/roles/:id", controller.DeleteRole)

			// program Routes
			router.GET("programs", controller.ListProgram)
			router.GET("/program/:id", controller.GetProgram)
			router.POST("/programs", controller.CreatePragram)
			router.PATCH("/programs", controller.UpdateProgram)
			router.DELETE("/programs/:id", controller.DeleteProgram)
		}
	}
	// Signup User Route
	r.POST("/signup", controller.CreateUser)
	// login User Route
	// r.POST("/login", controller.Login)
	r.POST("/signup_employee", controller.CreateEmployee)

	r.POST("/signup_student", controller.CreateLoginStudent)
	// login User Route
	r.POST("/login_employee", controller.LoginEmployee)

	r.POST("/login_student", controller.LoginStudent)

	// Run the server go run main.go
	r.Run("localhost: " + PORT)

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
