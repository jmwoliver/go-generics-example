package nogenerics

import "fmt"

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

func Send(payload interface{}) {
	switch t := payload.(type) {
	case PassivePayload:
		passivePayload, ok := payload.(PassivePayload)
		if ok {
			passivePayload.send()
		}
	case DirectedPayload:
		directedPayload, ok := payload.(DirectedPayload)
		if ok {
			directedPayload.send()
		}
	default:
		fmt.Printf("Type unknown: %v\n", t)
	}
}

func Example() {
	passivePayload := PassivePayload{id: 13}
	directedPayload := DirectedPayload{id: 22}

	Send(passivePayload)
	Send(directedPayload)
}
