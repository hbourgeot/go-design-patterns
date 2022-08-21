package main

import "fmt"

// Este patrón te permite suscribirte a nuevos eventos

type Observer interface {
	getId() string
	updateValue(string)
}

type Item struct {
	observers []Observer
	name      string
	avalaible bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i Item) UpdateAvalaible() {
	fmt.Printf("Item %s is avalaible\n", i.name)
	i.avalaible = true
	i.broadcast()
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

type EmailClient struct {
	id string
}

func (e EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s avalaible from client %s\n", value, e.id)
}

func (e *EmailClient) getId() string {
	return e.id
}

func main() {
	nvidiaItem := NewItem("RTX 3080")

	firstObserver := &EmailClient{
		id: "12ab",
	}

	secondObserver := &EmailClient{
		id: "34dc",
	}

	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvalaible()

}
