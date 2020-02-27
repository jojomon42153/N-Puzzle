/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 13:48:27 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/27 10:34:43 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

func createDataSet(n int) {
	var initial *state = &state{
		parent:        nil,
		state2D:       make([][]int, n),
		state1D:       make([]int, n*n),
		coord:         nil,
		initialCost:   0,
		heuristicCost: 0,
		totalCost:     0,
		isOpen:        true,
	}
	// init state 2D
	// 35 34 33 32 31 30
	// 29 28 27 26 25 24
	// 23 22 21 20 19 18
	// 17 16 15 14 13 12
	// 11 10 9  8  7  6
	// 5  4  3  2  1  0
	for i := 0; i < n; i++ {
		initial.state2D[i] = make([]int, n)
	}
	var i int = n*n - 1
	for y, line := range initial.state2D {
		for x := range line {
			initial.state2D[y][x] = i
			i--
		}
	}
	initial.zeroCoord = searchZeroCoord(initial.state2D)
	// init state1D
	for _, line := range initial.state2D {
		initial.state1D = append(initial.state1D, line...)
	}
	env.openedSet.tab[0] = initial
}

func main() {
	n := parse()
	initEnv(n)
	createDataSet(n)
	//aStar()
}
