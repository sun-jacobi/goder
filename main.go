package main

func main() {
	SetUpDB()
	SetUpRouter()
	router.Run(":3000")
}
