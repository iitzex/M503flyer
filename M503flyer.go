package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/parnurzeal/gorequest"
)

type wayPoint struct {
	name string
	lat  float64
	long float64
}

//M503 waypoints
var M503 = [...]wayPoint{
	{"BEGMO", 27.59, 121.50},
	{"OKATO", 27.35, 121.34},
	{"NUDPO", 26.46, 121.04},
	{"DUMAS", 26.07, 120.41},
	{"PONEN", 25.38, 120.23},
	{"OBKEL", 24.59, 119.52},
	{"APAKA", 23.51, 118.26}}

// {"TOLAK", 23.06, 117.29},
// {"LAPUG", 22.59, 117.22}}

//M503check for waypoints
func M503check(callsign string, lat float64, long float64) {
	if lat > 26 || lat < 21.5 {
		return
	}
	if long > 123 || long < 116 {
		return
	}

	var dist float64
	for _, wp := range M503 {
		dist = (wp.lat-lat)*(wp.lat-lat) + (wp.long-long)*(wp.long-long)
		if dist < 0.1 {
			fmt.Printf("%s %s %f %f %.3f\n", callsign, wp.name, wp.lat, wp.long, dist)
		}
	}
}

func main() {
	for {
		check()
		fmt.Println("--")
		time.Sleep(3 * time.Second)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func check() {
	request := gorequest.New()
	addr := "https://opensky-network.org/api/states/all"
	_, body, _ := request.Get(addr).End()

	js, err := simplejson.NewJson([]byte(body))
	if err != nil {
		panic("json format error")
	}
	t, _ := js.Get("time").Int64()
	// fmt.Println("time", t)
	T := time.Unix(t, 0)
	fmt.Println(T)

	states, err := js.Get("states").Array()

	for i := range states {
		p := js.Get("states").GetIndex(i)
		callsign, _ := p.GetIndex(1).String()
		longitude, _ := p.GetIndex(5).Float64()
		latitude, _ := p.GetIndex(6).Float64()
		// fmt.Println(callsign, latitude, longitude)
		M503check(callsign, latitude, longitude)
	}
}
