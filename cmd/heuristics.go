/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   heuristics.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/27 14:12:47 by gaennuye          #+#    #+#             */
/*   Updated: 2020/02/28 20:04:16 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"math"
)

/* Euclidian heuristic */

func euclidian(current *state) {

	sum := 0.0

	for i := 0; i < env.size; i++ {
		for j := 0; j < env.size; j++ {
			if current.state2D[i][j] != env.finalState.state2D[i][j] {

				goalC := env.finalState.coord[current.state2D[i][j]]
				currentErr := math.Pow(float64(i-goalC.x), 2) + math.Pow(float64(j-goalC.y), 2)
				sum += currentErr
			}
		}
	}
	current.heuristicCost += int(sum)
}

/* Tiles out of place heuristic */

func tilesOutOfPlace(current *state) {

	counter := 0
	sqSize := env.size * env.size

	for i := 0; i < sqSize; i++ {
		if current.state1D[i] != env.finalState.state1D[i] {
			counter++
		}
	}
	current.heuristicCost += int(counter)
}

func manhattan(actualState *state) {
	state2d := actualState.state2D
	coord := env.finalState.coord
	cost := 0
	for i := 0; i < env.size; i++ {
		for j := 0; j < env.size; j++ {
			if state2d[i][j] != 0 {
				cost += abs((coord[state2d[i][j]].y - i)) + abs((coord[state2d[i][j]].x - j))
			}
		}
	}
	actualState.heuristicCost += cost
}
