package mod

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i

	}
}
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {

			dst <- i
		}

	}
}
func sieve() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Start generate() as a subprocess.
	for {
		prime := <-ch

		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
