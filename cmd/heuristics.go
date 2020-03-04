/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   heuristics.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: gaennuye <gaennuye@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/27 14:12:47 by gaennuye          #+#    #+#             */
/*   Updated: 2020/03/04 16:48:13 by gaennuye         ###   ########lyon.fr   */
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
	current.totalCost = int(sum) //+ current.initialCost
}

/* Tiles out of place heuristic */

func tilesOutOfPlace(current *state) {

	counter := 0

	for i := 0; i < env.size; i++ {
		for j := 0; j < env.size; j++ {
			if current.state2D[i][j] != env.finalState.state2D[i][j] {
				counter++
			}
		}
	}
	current.heuristicCost += counter
	current.totalCost = counter //+ current.initialCost
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
	actualState.totalCost = cost //+ actualState.initialCost
}
