package main

func goSlice() {

	var colors = make([]string, 0)

	colors = append(colors, "yellow", "oragne", "pink")

	pl(colors)

	pl("array lenght is ", len(colors))

	var places = []string{"New York", "Huwaii", "Tokyo", "Melbourne"}

	for _, place := range places{

		pl(place)
	}

	var numbers = [][]int{{0,1,2,3},{4,5,6,7},{8,9,9,8}, {7,6,5,4},{3,2,1,0}}

	for _, row := range numbers{

		for i, col := range row {

			pf("%d ", col)

			if i == 3 {
				pl()
			}
		}
	}

	var games = make([]string, 3)

	pf("Games Lenght: %d \nGames Capacity: %d\n", len(games),cap(games))

	games[0] = "Spider Man 2"
	games[1] = "Cyberpunk"
	games[2] = "Howard's Legacy"
	games = append(games,  "DBZ Sparking Zero")

	pl("=========")

	pf("Games Lenght: %d \nGames Capacity: %d\n", len(games),cap(games))


}
