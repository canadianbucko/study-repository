package main

func main() {
	var nilChan chan int

	nilChan <- 0   // write to nil chan - block
	a := <-nilChan // read from nil chan - block
	close(nilChan) // close of nil chan - panic

	// read from closed chan - default value
	// write to closed chan - panic
	// close of closed chan - panic

}
