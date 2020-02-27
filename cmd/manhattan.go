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
	if value < 0 {
		value = -value
	}
	return value
}

func manhattan(state2D [][]int) int {
	cost := 0
	for i := 0; i < env.size; i++ {
		for j := 0; j < env.size; j++ {
			if state2D[i][j] != 0 {
				cost += abs((env.finalState.coord[state2D[i][j]].y - i)) + abs((env.finalState.coord[state2D[i][j]].x - j))
			}
		}
	}
	return (cost)
}
