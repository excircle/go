package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	// WaitGroup for Synchronization
	var wg sync.WaitGroup

	// Channel to receive orders
	var receiveOrdersCh = make(chan order)

	// Channel to confirm (validate) non-negative ints are used
	var validOrderCh = make(chan order)

	// Channel to process erroneous orders 
	var invalidOrderCh = make(chan invalidOrder)

	// Step 1 - Routine to put JSON data into pipeline: Data doesn't move until this triggers
	go receiveOrders(receiveOrdersCh)

	// Step 2 - 
	go validateOrders(receiveOrdersCh, validOrderCh, invalidOrderCh)
	
	// WaitGroup Init
	wg.Add(1)

	// Print when valid order received
	go func() {
		order := <- validOrderCh
		fmt.Printf("Valid order received: %v\n", order)
		wg.Done()
	}()

	// Print when invalid order received
	go func() {
		order := <- invalidOrderCh
		fmt.Printf("Invalid order received: %v. Issue: %v\n", order.order, order.err)
		wg.Done()
	}()

	// Deincrement wait group - Terminate routines
	wg.Wait()
}

func validateOrders(in chan order, out chan order, errCh chan invalidOrder) {
	order := <- in //receiveOrdersCh
	if order.Quantity <= 0 {
		errCh <- invalidOrder{order: order, err: errors.New("quantity must be greater than zero")}
	} else {
		out <- order
	}
}

// receiveOrders is given "receiveOrdersCh"
func receiveOrders(out chan order) {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		out <- newOrder //receiveOrdersCh
	}
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}