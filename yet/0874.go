type Robot struct {
	x int
	y int
	// NWSE => 0123
	dir uint8
}

func New() *Robot {
	return &Robot{x: 0, y: 0, dir: 0}
}

func euclidDis(r *Robot) int {
	return (r.x * r.x) + (r.y * r.y)
}

func robotSim(commands []int, obstacles [][]int) int {
	r := New()

	// -30000 => 0
	// 0 => 30000
	// 30000 => 60000
	obs_arr := make([][]bool, 60001)
	for i := 0; i < 60001; i++ {
		obs_arr[i] = make([]bool, 60001)
	}
	for _, obs := range obstacles {
		obs_arr[obs[0]+30000][obs[1]+30000] = true
	}

	for _, c := range commands {
		if c == -2 {
			r.dir = (r.dir + 1) % 4
		} else if c == -1 {
			r.dir = (r.dir + 3) % 4
		} else {

			switch r.dir {
			case 0:
				if -30000 <= r.x && r.x <= 30000 && -30000 <= r.y+1 && r.y+1 <= 30000 {
					for i := 0; i < c; i++ {
						if !obs_arr[r.x+30000][r.y+30001] {
							r.y++
						} else {
							break
						}
					}
				} else {
					r.y += c
				}
			case 1:
				if -30000 <= r.x-1 && r.x-1 <= 30000 && -30000 <= r.y && r.y <= 30000 {
					for i := 0; i < c; i++ {
						if !obs_arr[r.x+29999][r.y+30000] {
							r.x++
						} else {
							break
						}
					}
				} else {
					r.x -= c
				}
				r.x -= c
			case 2:
				if -30000 <= r.x && r.x <= 30000 && -30000 <= r.y-1 && r.y-1 <= 30000 {
					for i := 0; i < c; i++ {
						if !obs_arr[r.x+30000][r.y+29999] {
							r.y--
						} else {
							break
						}
					}
				} else {
					r.y -= c
				}
				r.y -= c
			case 3:
				if -30000 <= r.x+1 && r.x+1 <= 30000 && -30000 <= r.y && r.y <= 30000 {
					for i := 0; i < c; i++ {
						if !obs_arr[r.x+30001][r.y+30000] {
							r.x++
						} else {
							break
						}
					}
				} else {
					r.x += c
				}
				r.x += c
			}
		}
	}

	return euclidDis(r)
}