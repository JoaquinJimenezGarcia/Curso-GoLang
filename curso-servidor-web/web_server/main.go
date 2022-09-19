package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.AddMidleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
