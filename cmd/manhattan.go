package main

import "math"

type tab struct {
	size int;
	configuration [3][3]int;
	cost int;
	heuristicValue int;
	expansionOrder int;
}

func manhattan() {
	i := 0;
	j := 0;
	x := 0;
	y := 0;
	var actual tab;
	var goal tab;
	actual.size = 3;
	actual.configuration[0][0] = 1;
	actual.configuration[0][1] = 2;
	actual.configuration[0][2] = 3;
	actual.configuration[1][0] = 0;
	actual.configuration[1][1] = 4;
	actual.configuration[1][2] = 5;
	actual.configuration[2][0] = 8;
	actual.configuration[2][1] = 7;
	actual.configuration[2][2] = 6;


	goal.configuration[0][0] = 1;
	goal.configuration[0][1] = 2;
	goal.configuration[0][2] = 3;
	goal.configuration[1][0] = 8;
	goal.configuration[1][1] = 0;
	goal.configuration[1][2] = 4;
	goal.configuration[2][0] = 7;
	goal.configuration[2][1] = 6;
	goal.configuration[2][2] = 5;

	for (i < actual.size) {
		j = 0;
		for (j < actual.size){
			if (actual.configuration[i][j] != goal.configuration[i][j] && actual.configuration[i][j] !=0) {
				x = 0;
				for(x < actual.size) {
					y = 0;
					for (y < actual.size){
						if(actual.configuration[i][j] == goal.configuration[x][y]){
							actual.heuristicValue = actual.heuristicValue + int(math.Abs(float64(i-x)) + math.Abs(float64(j-y)));
						}
						y++;
					}
					x++;
				}
			}
			j++;
		}
		i++;
	}
	println(actual.heuristicValue);

}