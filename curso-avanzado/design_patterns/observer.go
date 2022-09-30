package main

import "fmt"

type Topic interface {
	register(oberver Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

type EmailClient struct {
	id string
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is avaiblable\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (eC *EmailClient) getId() string {
	return eC.id
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending email... %s available from client %s\n", value, eC.getId())
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &EmailClient{
		id: "14ac",
	}

	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)

	nvidiaItem.UpdateAvailability()
}
