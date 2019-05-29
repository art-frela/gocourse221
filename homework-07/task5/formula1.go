//
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	distance float64 = 100
)

type bolide struct {
	number     string
	driver     string
	delayStart time.Duration
	maxSpeed   int
}

func (c *bolide) sitDriver(driver string) {
	c.driver = driver
}

func (c *bolide) newBolide(number string) {
	c.number = number
	c.delayStart = time.Duration(rand.Intn(10))
	c.maxSpeed = rand.Intn(200) + 20
}

func (c *bolide) toRun(distance float64, run <-chan struct{}, finished, heardbeat chan<- string, wg *sync.WaitGroup) {
	// go to start
	//
	time.Sleep(time.Second * c.delayStart)
	wg.Done() // report for ready to go
	heardbeat <- fmt.Sprintf("%s-%s ready to start", c.number, c.driver)
	// fromStart := float64(0.0)
	// go func() {
	// 	for {
	// 		time.Sleep(5 * time.Second)
	// 		delta := c.maxSpeed * 5 / 18 * 5   // convert km/h to m/s *5/18
	// 		fromStart += float64(delta) / 1000 // conver meters to kilometers
	// 		remain := distance - fromStart
	// 		heardbeat <- fmt.Sprintf("%s got %.2fkms from start, remain %.2fkms", c.number, fromStart, remain)
	// 	}
	// }()

	select {
	case <-run:
		// run
		durationRaceSec := distance / float64((c.maxSpeed * 5 / 18)) // duration in seconds
		time.Sleep(time.Duration(durationRaceSec) * time.Second)
		finished <- fmt.Sprintf("driver [%s], bolide #%s with result\t[%.1fsec]", c.driver, c.number, durationRaceSec)
		return
	}

}

func main() {

	// make drivers
	drivers := [][]string{{"Speedracer-01", "111"}, {"Speedracer-02", "222"}, {"Speedracer-03", "333"}}
	var wg sync.WaitGroup
	var wgE sync.WaitGroup
	start := make(chan struct{})
	finish := make(chan string, len(drivers))
	heardbeat := make(chan string)
	// make cars-bolides
	rand.Seed(time.Now().Unix())
	t := time.Now()
	for _, driver := range drivers {
		car := bolide{}
		car.newBolide(driver[1])
		car.sitDriver(driver[0])
		wg.Add(1)
		wgE.Add(1)
		// print car properties
		fmt.Printf("New bolide #%s with driver: %s. MaxSpeed=%d; to start delay [%dsec]\n", car.number, car.driver, car.maxSpeed, car.delayStart)
		// engine car
		go car.toRun(distance, start, finish, heardbeat, &wg)
	}
	wg.Wait() // waiting for arrive to start line
	fmt.Printf("Distance - %.f meters\nLet's GOOOOoooooo!\n", distance)
	close(start) // eq command START!!!

	place := 1
	go func() { //start finish print
		for finisher := range finish {
			fmt.Printf("Place-%d for %s\n", place, finisher)
			place++
			wgE.Done()
		}
		fmt.Println("Race over!")
	}()

	go func() { //start heardbeat print
		for hb := range heardbeat {
			fmt.Printf("\t%s from start %q\n", hb, time.Since(t))
		}

	}()
	wgE.Wait()

}
