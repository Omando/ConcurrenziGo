package sequential

import (
	"fmt"
	"github.com/cucumber/godog"
	"strconv"
	"testing"
)

var list LinkedList // interface implemented by singlyLinkedList and doublyLinkedList
var isItemFound bool

/*  Given */
func linkedListImplementation(arg string) error {
	if arg == "singlyLinkedList" {
		list = NewSLL()
	} else if arg == "doublyLinkedList" {
		list = NewDLL()
	} else if arg == "CircularlyLinkedList" {
		list = NewCLL()
	}

	return nil
}

func iAppendItems(items *godog.Table) error {
	for i := 1; i < len(items.Rows); i++ {
		value, err := strconv.Atoi(items.Rows[i].Cells[0].Value)
		if err != nil {
		}
		list.Append(value)
	}
	return nil
}

/* When */
func iCreateANewList() error {
	list = NewSLL()
	return nil
}

func iPrependItems(items *godog.Table) error {
	for i := 1; i < len(items.Rows); i++ {
		value, err := strconv.Atoi(items.Rows[i].Cells[0].Value)
		if err != nil {
		}
		list.Prepend(value)
	}
	return nil
}

func iRemove(arg1 int) error {
	_, err := list.Remove(arg1)
	if err != nil {
		isItemFound = false
	}
	return nil
}

func iSearchFor(arg1 int) error {
	isItemFound = list.Contains(arg1)
	return nil
}

/* Then */
func headAndTailAreNilAndSizeIsZero() error {
	if actualValue, err := list.First(); err == nil {
		return fmt.Errorf("expected head to be nil, but actual is: %d", actualValue)
	}

	if actualValue, err := list.Last(); err == nil {
		return fmt.Errorf("expected tail to be nil, but actual is: %d", actualValue)
	}
	if list.Length() != 0 {
		return fmt.Errorf("expected length be 0, but actual is: %d", list.Length())
	}

	return nil
}

func headIs(expectedValue int) error {
	if actualValue, err := list.First(); err != nil && expectedValue != actualValue {
		return fmt.Errorf("expected head to be %d, but actual is: %d", expectedValue, actualValue)
	}
	return nil
}

func itemIsFound() error {
	if isItemFound != true {
		return fmt.Errorf("expected item to be found")
	}
	return nil
}

func itemIsNotFound() error {
	if isItemFound != false {
		return fmt.Errorf("expected item not to be found")
	}
	return nil
}

func tailIs(expectedValue int) error {
	if actualValue, err := list.Last(); err != nil && actualValue != expectedValue {
		return fmt.Errorf("expected tail to be: %d, but actual is: %d", expectedValue, actualValue)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^linked list implementation is "([^"]*)"$`, linkedListImplementation)
	ctx.Step(`^head and tail are nil and size is zero$`, headAndTailAreNilAndSizeIsZero)
	ctx.Step(`^Head is (\d+)$`, headIs)
	ctx.Step(`^I append items$`, iAppendItems)
	ctx.Step(`^I create a new list$`, iCreateANewList)
	ctx.Step(`^I prepend items$`, iPrependItems)
	ctx.Step(`^I remove (\d+)$`, iRemove)
	ctx.Step(`^I search for (\d+)$`, iSearchFor)
	ctx.Step(`^item is found$`, itemIsFound)
	ctx.Step(`^item is not found$`, itemIsNotFound)
	ctx.Step(`^Tail is (\d+)$`, tailIs)
}

func Test_Runner(t *testing.T) {
	opts := godog.Options{
		Format: "progress",
		Paths:  []string{"features"},
	}

	godog.TestSuite{
		Name:                "",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()
}
