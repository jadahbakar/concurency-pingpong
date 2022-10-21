package main

import (
	"log"
	"math/rand"
	"time"
)

type ball struct {
	hits              int
	lastPlayerHitBall string
}

func main() {
	// instantiate the channel
	table := make(chan *ball)
	done := make(chan *ball)
	// each player with difference go routine
	go player("dedy", table, done)
	go player("adit", table, done)

	go player("diah", table, done)
	go player("anin", table, done)

	referree(table, done)
}

func referree(table chan *ball, done chan *ball) {
	// pass new ball to table(channel)
	table <- new(ball)
	for {
		select {
		case ball := <-done:
			log.Println("winner is: ", ball.lastPlayerHitBall)
			return

		}

	}
}

// table is a channel
// imagine when player A hit the ball,
// the ball will hit the table (which is channel)
// and after that from table will bounching to player B

// done to return ball to referree
func player(name string, table chan *ball, done chan *ball) {

	for {
		// create random conditon when player cannot return the ball to another player
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		// create 2 condition
		select {
		// get ball, is the condition that we get the ball
		case ball := <-table:
			v := r.Intn(1000)
			if v%11 == 0 {
				// when stop
				// return the ball to channel ball
				// instead return the ball to table (channel) -> so the ball can get to referree
				log.Println(name, "drop the ball")
				done <- ball
				// for the loop break, and make sure the GoRoutine finish, to avoid memory leaks
				return
			}
			// count how many ball beeing hits
			ball.hits++
			ball.lastPlayerHitBall = name
			log.Println(name, "hits the ball", ball.hits)
			time.Sleep(50 * time.Millisecond)
			// after player A finish
			// give the ball to player B
			// so assign the ball to channel / return the ball to table (channel)
			table <- ball
		case <-time.After(2 * time.Second):
			return

		}

	}

}
