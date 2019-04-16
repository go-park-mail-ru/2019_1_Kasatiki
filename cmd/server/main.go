package server

func main() {
	advhater := App{}
	advhater.Initialize()
	advhater.Run(":8080")
}
