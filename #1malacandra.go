package main
//目的地距離56000000 km ,到達目的地需要28天,每小時為多少?
import (
	"fmt"
)

func main() {
	const distance = 56000000
	var hour = 28 * 24
	fmt.Printf("距離:%v km / 時間:%v h = 時速:%v km/h", distance, hour, distance/hour)
}
