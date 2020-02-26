/* **************************************************************************** */
/*                                                                              */
/*                                                         :::      ::::::::    */
/*    manhattan.go                                       :+:      :+:    :+:    */
/*                                                     +:+ +:+         +:+      */
/*    By: jdarko <marvin@42.fr>                      +#+  +:+       +#+         */
/*                                                 +#+#+#+#+#+   +#+            */
/*    Created: 2020/02/26 13:29:16 by jdarko            #+#    #+#              */
/*    Updated: 2020/02/26 13:29:19 by jdarko           ###   ########lyon.fr    */
/*                                                                              */
/* **************************************************************************** */

package main

func abs(value int) int {
	if (value < 0){
		value = -value;
	}
	return value;
}

func manhattan(state2D[][] int) int {
	i := 0;
	j := 0;
	cost := 0;

	//env.size = 3;
	//env.finalCoord[0].x = 1;
	//env.finalCoord[0].y = 1;
	//env.finalCoord[1].x = 0;
	//env.finalCoord[1].y = 0;
	//env.finalCoord[2].x = 1;
	//env.finalCoord[2].y = 0;
	//env.finalCoord[3].x = 2;
	//env.finalCoord[3].y = 0;
	//env.finalCoord[4].x = 2;
	//env.finalCoord[4].y = 1;
	//env.finalCoord[5].x = 2;
	//env.finalCoord[5].y = 2;
	//env.finalCoord[6].x = 1;
	//env.finalCoord[6].y = 2;
	//env.finalCoord[7].x = 0;
	//env.finalCoord[7].y = 2;
	//env.finalCoord[8].x = 0;
	//env.finalCoord[8].y = 1;

	for (i < env.size) {
		j = 0;
		for (j < env.size){
			if (state2D[i][j] != 0){
				cost += abs((env.finalCoord[state2D[i][j]].y - i)) + abs((env.finalCoord[state2D[i][j]].x - j))
			}
			j++;
		}
		i++;
	}
	print(cost);
	return (cost);
}
