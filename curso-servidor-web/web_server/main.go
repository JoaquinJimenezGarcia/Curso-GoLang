package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/api", server.AddMidleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
