package main

func main() {
	// sanity check
	ch := make(chan string, 1)
	<-ch
}
