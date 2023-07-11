package channels

import (
	"fmt"
	"math/rand"
)

func LoadLeader(c chan<- string) {
	leaders := []string{"Harry", "Max", "Tom"}
	c <- leaders[rand.Intn(len(leaders)-1)]
}

func LoadTeamMember(c chan<- string) {
	members := []string{"Mari", "Jorge", "Raul", "Thomas"}
	c <- members[rand.Intn(len(members)-1)]
}

func ShowUsers(c <-chan string) {
	msg := <-c
	fmt.Println(msg)
}

func TestChannelDirection() {
	var c_receiver chan<- string = make(chan<- string)
	var c_sender <-chan string = make(<-chan string)

	go LoadLeader(c_receiver)
	go LoadTeamMember(c_receiver)
	go ShowUsers(c_sender)
}
