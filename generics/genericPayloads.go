package generics

import "fmt"

type Payload interface {
	PassivePayload | DirectedPayload
	send()
}

func Send[T Payload](payload T) {
	payload.send()
}

type PassivePayload struct {
	id int
}

func (pp PassivePayload) send() {
	fmt.Printf("PassivePayload.Send(%d)\n", pp.id)
}

type DirectedPayload struct {
	id int
}

func (dp DirectedPayload) send() {
	fmt.Printf("DirectedPayload.Send(%d)\n", dp.id)
}

func Example() {
	passivePayload := PassivePayload{id: 13}
	directedPayload := DirectedPayload{id: 22}

	Send(passivePayload)
	Send(directedPayload)
}
