package main

func main() {

	// NIL CHANNEL
	/* sending or receving on a nil channel,
	   blocks forever */
	{
		var ch chan error
		go func() {
			ch <- nil
			<-ch
			println("nil never ok")
		}()
	}

	// UNBUFFERED CHANNEL
	/* sending or receving on an unbuffered channel,
	   blocks until the data is received/sent */
	{
		ch := make(chan error)
		go func() {
			ch <- nil
			println("unbuffered ok")
		}()
		<-ch
	}

	// BUFFERED CHANNEL
	/* sending or receiving on a buffered channel,
	   doesn't block if the channel is not full/empty */
	{
		ch := make(chan error, 2)
		ch <- nil
		ch <- nil
		go func() {
			ch <- nil
			ch <- nil
			ch <- nil
			println("buffered ok")
		}()

		for i := 0; i < 5; i++ {
			<-ch
		}
	}

	// RANGE CHANNEL
	/* using range for receiving from a channel,
	   keeps draining data until the channel is closed */
	{
		ch := make(chan error, 10)
		for i := 0; i < 10; i++ {
			ch <- nil
		}

		// only the sender should always close the channel!
		close(ch)

		for _ = range ch {
			// ...
		}
		println("range ok")
	}

	// CLOSED CHANNEL
	/* receiving from a closed channel,
	   returns zero value of the data and
	   false as second parameter */
	{
		ch := make(chan error)
		close(ch)
		v, ok := <-ch

		if v == nil && ok {
			println("closed ok")
		}

		// send on a closed channel will panic!
		// ch <- nil

		// closing already closed channel, will panic!
		// close(ch)
	}

	// SELECT CHANNELS
	/* calling select for sending or receiving on channels,
	   blocks until some case is ready or
	   runs default case if present and none is ready */
	{
		in := make(chan error, 1)
		in <- nil
		out := make(chan error, 1)
		ch := make(chan error)
		close(ch)

		select {
		case in <- nil:
			// blocks; channel is full
		case <-out:
			// blocks; channel is empty
		case <-ch:
			// doesn't block; channel is closed
			println("doesn't block; channel is closed")
		default:
			// runs if all other cases block
			println("runs if all other cases block")
		}
		println("default done")
	}
}
