package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, unitExists := units[unit]
	if !unitExists {
		return false
	}
	_, itemExists := bill[item]
	if itemExists {
		bill[item] = bill[item] + units[unit]
		return true
	}
	bill[item] = units[unit]
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, unitExists := units[unit]
	_, itemExists := bill[item]
	if !unitExists || !itemExists {
		return false
	}
	removal := bill[item] - units[unit]
	if removal < 0 {
		return false
	}
	if removal == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = removal
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemName, itemExists := bill[item]
	return itemName, itemExists
}
