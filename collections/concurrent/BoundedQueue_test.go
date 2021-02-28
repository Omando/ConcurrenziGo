package concurrent

import (
	"container/list"
	"fmt"
	"github.com/cucumber/godog"
	"sync"
	testing "testing"
)

type BoundedQueueTest struct {
	consumerCount int
	producerCount int
	capacity      int
	buffer        *BoundedBuffer
}

// You can run and debug this test
func Test_Main(m *testing.T) {
	opts := godog.Options{
		Format: "progress",
		Paths:  []string{"features"},
		//Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	}

	godog.TestSuite{
		//TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options: &opts,
	}.Run()
}

var boundedQueueTest BoundedQueueTest

func thereAreProducers(producerCount int) error {
	boundedQueueTest.producerCount = producerCount
	return nil
}

func thereAreConsumers(consumerCount int) error {
	boundedQueueTest.consumerCount = consumerCount
	return nil
}

func boundedQueueCapacity(capacity int) error {
	boundedQueueTest.capacity = capacity
	return nil
}

func whenProducersAndConsumersRun() error {
	boundedQueueTest.buffer = &BoundedBuffer{
		Data:      list.New(),
		Capacity:  boundedQueueTest.capacity,
		Condition: sync.NewCond(&sync.Mutex{}),
	}

	// Two waitgroups are used to synchronize the start and end of goroutines.
	// All goroutines wait on startGate to be opened by this function
	var startGate sync.WaitGroup
	startGate.Add(1)

	// This function waits on this endGate to be opened once the last goroutine completes
	// producers
	var endGate sync.WaitGroup
	endGate.Add(boundedQueueTest.consumerCount + boundedQueueTest.producerCount)

	for i := 0; i < boundedQueueTest.producerCount; i++ {
		go func(itemToAdd int) {
			startGate.Wait()
			boundedQueueTest.buffer.Enqueue(itemToAdd)
			endGate.Done()
			fmt.Println("Adding completed")
		}(i)
	}

	// consumers
	for i := 0; i < boundedQueueTest.consumerCount; i++ {
		go func() {
			startGate.Wait()
			item := boundedQueueTest.buffer.Dequeue()
			endGate.Done()
			fmt.Println("Removed ", item)
		}()
	}

	// Start all goroutines
	fmt.Println("Starting all goroutines")
	startGate.Done()

	// Wait until all go routines are done
	endGate.Wait()
	fmt.Println("All goroutines completed")
	return nil
}

func thereShouldBeItemsRemaining(remainingItemCount int) error {
	actualItemCount := boundedQueueTest.buffer.Data.Len()
	if remainingItemCount != actualItemCount {
		return fmt.Errorf("Expected %d items remaining, but found %d",
			remainingItemCount, actualItemCount)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^there are (\d+) producers$`, thereAreProducers)
	ctx.Step(`^there are (\d+) consumers$`, thereAreConsumers)
	ctx.Step(`^bounded queue capacity is (\d+)$`, boundedQueueCapacity)
	ctx.Step(`^when producers and consumers run$`, whenProducersAndConsumersRun)
	ctx.Step(`^there should be (\d+) items remaining$`, thereShouldBeItemsRemaining)
}
