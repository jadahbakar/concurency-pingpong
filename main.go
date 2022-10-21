package main

import (
	"log"
	"time"
)

type ball struct {
	hits int
}

func main() {
	// instantiate the channel
	table := make(chan *ball)
	// each player with difference go routine
	go player("dedy", table)
	go player("adit", table)

	// pass new ball to table(channel)
	table <- new(ball)
	time.Sleep(1 * time.Second)

	// to terminate
	<-table

}

// table is a channel
// imagine when player A hit the ball,
// the ball will hit the table (which is channel)
// and after that from table will bounching to player B
func player(name string, table chan *ball) {

	for {
		// get the ball, with assing variable with name ball
		ball := <-table
		// count how many ball beeing hits
		ball.hits++
		log.Println(name, "hits the ball", ball.hits)
		time.Sleep(50 * time.Millisecond)
		// after player A finish
		// give the ball to player B
		// so assign the ball to channel / return the ball to table (channel)
		table <- ball
	}

}
