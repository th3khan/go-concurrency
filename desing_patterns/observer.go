package main

import "fmt"

type Topic interface {
	register(observe Observer)
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

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item: %s is available\n", i.name)
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

type EmailClient struct {
	id string
}

func (eC *EmailClient) getId() string {
	return eC.id
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email, Item: %s  Available, Client: %s\n", value, eC.id)
}

func main() {
	nvidiaItem := NewItem("RTX 3000")
	firtsObserver := &EmailClient{
		id: "1",
	}
	secondObserver := &EmailClient{
		id: "2",
	}
	nvidiaItem.register(firtsObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvailable()
}
