package partyrobot

import (
    "strconv"
    "fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	// panic("Please implement the Welcome v.function")
    return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	// panic("Please implement the HappyBirthday function")
    return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// panic("Please implement the AssignTable function")j
    return "Welcome to my party, " + name + "!\n" + "You have been assigned to table " + fmt.Sprintf("%03d", table) + ". Your table is " + direction + ", exactly " + strconv.FormatFloat(distance, 'f', 1, 64) + " meters from here.\n" + "You will be sitting next to " + neighbor + "."
    
}
