package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		log.Println(fmt.Sprintf("Header[%q] = %q", k, v))
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		log.Println(fmt.Sprintf("qs[%q] = %q", k, v))
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, r.Form)
}

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var width, height = 600, 320
var zscale = float64(height) * 0.4
var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var xyscale float64 = float64(width) / 2 / xyrange

func surface(w http.ResponseWriter, form url.Values) {
	var err error
	width_str = form["width"]
	height_str = form["width"]
	width, err = strconv.Atoi(form["width"][0])
	height, err = strconv.Atoi(form["height"][0])
	if err != nil {
		log.Print(err)
	}

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			valid := true
			ax, ay, ok := corner(i+1, j)
			valid = valid && ok
			bx, by, ok := corner(i, j)
			valid = valid && ok
			cx, cy, ok := corner(i, j+1)
			valid = valid && ok
			dx, dy, ok := corner(i+1, j+1)
			valid = valid && ok
			if valid {
				fmt.Fprintf(w, "<polygon fill='#FF0000' points='%g,%g %g,%g %g,%g %g,%g' />\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	if math.IsInf(sx, 1) || math.IsInf(sx, -1) || math.IsNaN(sx) || math.IsInf(sy, 1) || math.IsInf(sy, -1) || math.IsNaN(sy) {
		return 0, 0, false
	}
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func g(x, y float64) float64 {
	return 0
}

func h(x, y float64) float64 {
	return math.Sin(x * y)
}
