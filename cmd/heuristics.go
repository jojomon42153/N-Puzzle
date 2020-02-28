/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   heuristics.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: gaennuye <gaennuye@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/27 14:12:47 by gaennuye          #+#    #+#             */
/*   Updated: 2020/02/28 12:06:05 by gaennuye         ###   ########lyon.fr   */
/*                                                                            */
/* ************************************************************************** */

package main

import "math"
import "fmt"


/* Euclidian heuristic */

func euclidian(current *state) {

	sum := 0.0

	for i := 0; i < env.size; i++ {
		for j := 0; j < env.size; j++ {
			if (current.state2D[i][j] != env.finalState.state2D[i][j]) {

				goal_c := env.finalState.coord[current.state2D[i][j]]
				current_err := math.Pow(float64(i - goal_c.x), 2 ) + math.Pow(float64(j - goal_c.y), 2) 
				sum += current_err
			}
		}
	}
	fmt.Println("euclidian : ", sum)

	current.heuristicCost += int(sum)
}


/* Tiles out of place heuristic */

func tilesOutOfPlace(current *state) {

	counter := 0
	sq_size := env.size * env.size

	for i := 0; i < sq_size; i++ {
		if (current.state1D[i] != env.finalState.state1D[i]){
			counter++
		}
	}
	fmt.Println("toop : ", counter)

	current.heuristicCost += int(counter)
}

