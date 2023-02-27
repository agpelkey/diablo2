package main

func main() {

	server := NewAPIServer(":8080", Routes)

	server.Run()
}
