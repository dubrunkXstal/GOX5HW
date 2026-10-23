package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidID = errors.New("id должен быть положительным")
	ErrNoOrder   = errors.New("заказ не задан")
	ErrEmptyName = errors.New("имя не может быть пустым")
	ErrBadPrice  = errors.New("цена не может быть отрицательной")
)

type Item struct {
	Name  string
	Price int64
}

type Order struct {
	ID     int
	Items  []Item
	Status string
}

func newOrder(id int) (Order, error) {
	if id <= 0 {
		return Order{}, ErrInvalidID
	}

	return Order{ID: id, Status: "new"}, nil
}

func addItem(order *Order, item Item) error {
	if order == nil {
		return ErrNoOrder
	}
	if item.Name == "" {
		return ErrEmptyName
	}
	if item.Price < 0 {
		return ErrBadPrice
	}

	order.Items = append(order.Items, item)
	return nil
}

func total(order Order) int64 {
	var result int64

	for _, item := range order.Items {
		result += item.Price
	}

	return result
}

var Statuses = statusesInit()

func statusesInit() map[string]string {
	statuses := make(map[string]string)

	statuses["new"] = "новый"
	statuses["processing"] = "в обработке"
	statuses["paid"] = "оплачен"
	statuses["cancelled"] = "отменён"

	return statuses
}

func main() {
	order, err := newOrder(47)
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}

	_ = addItem(&order, Item{Name: "Кофе", Price: 19_990})
	_ = addItem(&order, Item{Name: "Фильтр", Price: 2_500})

	fmt.Printf("заказ %d, статус %s, сумма %d коп.\n",
		order.ID, Statuses[order.Status], total(order))
}
