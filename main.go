package main

func main() {
	SetUpDB()
	SetUpRouter()
	Router.Run(":3000")
}
