package main

func main() {
	addr := ":4000"
	s := NewAPIServer(addr)
	s.Run()
}
