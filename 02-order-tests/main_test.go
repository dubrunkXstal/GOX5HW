package main

import (
	"testing"
)

func TestZero(t *testing.T) {
	order, err := newOrder(48)
	if err != nil {
		t.Errorf("Ошибка: %d", err)
		return
	}

	result := total(order)
	expect := int64(0)

	if result != expect {
		t.Errorf("Ошибка: валидный заказ без позиций должен иметь сумму 0")
	}
}

func TestSum(t *testing.T) {
	order, err := newOrder(48)
	if err != nil {
		t.Errorf("Ошибка: %d", err)
		return
	}

	_ = addItem(&order, Item{Name: "Кофе", Price: 19_990})
	_ = addItem(&order, Item{Name: "Фильтр", Price: 2_500})

	result := total(order)
	expect := int64(19_990 + 2_500)

	if result != expect {
		t.Errorf("Ошибка: после добавления позиций их сумма не равна сумме позиций")
	}
}

func TestAdd(t *testing.T) {
	order, err := newOrder(48)
	if err != nil {
		t.Errorf("Ошибка: %d", err)
		return
	}

	result := addItem(&order, Item{Name: "Кофе", Price: 19_990})
	expect := error(nil)

	if result != expect {
		t.Errorf("Ошибка: непустая ошибка при корректном заказе")
	}

	result = addItem(nil, Item{Name: "Кофе", Price: 19_990})
	expect = ErrNoOrder

	if result != expect {
		t.Errorf("Ошибка: несоответствие ошибки при некорректном заказе")
	}

	result = addItem(&order, Item{Name: "", Price: 19_990})
	expect = ErrEmptyName

	if result != expect {
		t.Errorf("Ошибка: несоответствие ошибки при некорректном имени заказа")
	}

	result = addItem(&order, Item{Name: "Кофе", Price: -1})
	expect = ErrBadPrice

	if result != expect {
		t.Errorf("Ошибка: несоответствие ошибки при некорректной цене заказа")
	}
}

func TestCopy(t *testing.T) {
	order, err := newOrder(48)
	if err != nil {
		t.Errorf("Ошибка: %d", err)
		return
	}

	order_copy := order

	_ = addItem(&order_copy, Item{Name: "Кофе", Price: 19_990})

	result := len(order.Items)
	expect := 0

	if result != expect {
		t.Errorf("Ошибка: изменение копии Order меняет исходную структуру")
	}
}

func TestMap(t *testing.T) {
	if Statuses == nil {
		t.Error("Ошибка: Statuses не создан")
		return
	}

	result := Statuses["new"]
	expect := "новый"

	if result != expect {
		t.Errorf("Ошибка: ключ new не совпадает со значением")
	}

	result = Statuses["processing"]
	expect = "в обработке"

	if result != expect {
		t.Errorf("Ошибка: ключ processing не совпадает со значением")
	}

	result = Statuses["paid"]
	expect = "оплачен"

	if result != expect {
		t.Errorf("Ошибка: ключ paid не совпадает со значением")
	}

	result = Statuses["cancelled"]
	expect = "отменён"

	if result != expect {
		t.Errorf("Ошибка: ключ cancelled не совпадает со значением")
	}
}
