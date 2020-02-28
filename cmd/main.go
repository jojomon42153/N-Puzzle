/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 13:48:27 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/28 18:03:20 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"log"
	"os"
)

func createFirstState(twoD [][]int) {
	initEnv(env.size)
	var initial *state = &state{
		parent:        nil,
		state2D:       twoD,
		state1D:       make([]int, 0),
		coord:         nil,
		initialCost:   0,
		heuristicCost: 0,
		totalCost:     0,
		isOpen:        true,
	}
	initial.zeroCoord = searchZeroCoord(initial.state2D)
	for _, line := range initial.state2D {
		initial.state1D = append(initial.state1D, line...)
	}
	initial.index = arrayToString(initial.state1D, ",")
	calcHeuristicCost(initial)

	if !checkSolvability(initial) {
		log.Fatal("Taquin is not resolvable")
		os.Exit(1)
	}

	env.openedSet.tab[0] = initial
	env.allSets[initial.index] = initial
}

func main() {
	calcHeuristicCost = manhattan
	fmt.Println("ici2")
	parse("ressources/correctInput/taquin_dim4_0.txt")

	ch = make(map[string]chan int, 2)
	ch["nbOpened"] = make(chan int)
	ch["nbClosed"] = make(chan int)
	go updateNbOpened()
	go updateNbClosed()

	aStar()

	for _, chanel := range ch {
		close(chanel)
	}
}
