// 286. Walls and Gates
// https://leetcode.com/problems/walls-and-gates/

package main

import "fmt"

var ()

type visits struct {
	Rooms map[int]map[int]bool
}

func (v *visits) AddRoom(r, c int) {
	v.Rooms[r][c] = true
}

func (v *visits) Visted(r, c int) bool {
	return v.Rooms[r][c]
}

type qInt struct {
	Items [][]int
}

func (q *qInt) Add(x, y int) {
	q.Items = append(q.Items, []int{x, y})
}

func (q *qInt) Len() int {
	return len(q.Items)
}

func (q *qInt) PopLeft() (int, int) {
	e := q.Items[0]
	q.Items = q.Items[1:]
	return e[0], e[1]
}

func (q *qInt) Print() {
	fmt.Printf("Current Queue: %v", q.Items)
}

func newVisits(r int) *visits {
	v := &visits{
		Rooms: make(map[int]map[int]bool),
	}
	for i := 0; i < r; i++ {
		v.Rooms[i] = make(map[int]bool)
	}
	return v
}

func queueRoom(r, c int, rooms [][]int, v *visits, q *qInt) {
	if r >= len(rooms) || r < 0 {
		return
	}
	if c >= len(rooms[0]) || c < 0 {
		return
	}
	if rooms[r][c] == -1 || v.Visted(r, c) {
		return
	}
	v.AddRoom(r, c)
	q.Add(r, c)
}

func wallsAndGates(rooms [][]int) {
	roomQ := qInt{}
	v := newVisits(len(rooms))
	dist := 0
	for r := 0; r < len(rooms); r++ {
		for c := 0; c < len(rooms[r]); c++ {
			if rooms[r][c] == 0 {
				v.AddRoom(r, c)
				roomQ.Add(r, c)
			}
		}
	}

	//for dist < len(rooms)*len(rooms[0]) {
	queueDepth := roomQ.Len()
	for roomQ.Len() > 0 {
		for i := 1; i <= queueDepth; i++ {
			//fmt.Printf("Parsing %v messages\n", queueDepth)
			r, c := roomQ.PopLeft()
			//fmt.Printf("r: %v, c: %v\n", r, c)
			//fmt.Printf("dist: %v\n", dist)
			rooms[r][c] = dist
			//fmt.Printf("Rooms[%v][%v] has a distance of %v\n", r, c, rooms[r][c])
			queueRoom(r+1, c, rooms, v, &roomQ)
			queueRoom(r-1, c, rooms, v, &roomQ)
			queueRoom(r, c+1, rooms, v, &roomQ)
			queueRoom(r, c-1, rooms, v, &roomQ)
		}
		dist += 1
		queueDepth = roomQ.Len()
	}
}

func main() {
	rooms := [][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}}
	wallsAndGates(rooms)
	/* test := make(map[int]map[int]bool)
	//test[0][1] = true
	if _, ok := test[1][0]; ok {
		fmt.Println("this feels dirty")
	}
	test[0] = make(map[int]bool)
	test[0][1] = true
	if test[1][0] == true {
		fmt.Println("this feels dirty")
	}
	fmt.Println(test[1][0]) */
}
