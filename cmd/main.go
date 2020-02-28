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
	"flag"
	"log"
	"os"
)

func calculHeurisstic(state *state) {
	for _, function := range calcHeuristicCost {
		function(state)
	}
}

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

	calculHeurisstic(initial)

	if !checkSolvability(initial) {
		log.Fatal("Taquin is not resolvable")
		os.Exit(1)
	}

	env.openedSet.tab[0] = initial
	env.allSets[initial.index] = initial
}

func arg() {
	manhattanPtr := flag.Bool("hm", false, ": select manhattan as heuristic")
	toopPtr := flag.Bool("ht", false, ": select toop as heuristic")
	euclideanPtr := flag.Bool("he", false, ": select euclidean as heuristic")
	filePtr := flag.String("f", "", ": -f thePuzzle.txt")
	flag.Parse()
	if *filePtr == "" {
		println("need a file !")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *manhattanPtr == true {
		calcHeuristicCost = append(calcHeuristicCost, manhattan)
	}
	if *toopPtr == true {
		calcHeuristicCost = append(calcHeuristicCost, tilesOutOfPlace)
	}
	if *euclideanPtr == true {
		calcHeuristicCost = append(calcHeuristicCost, euclidian)
	}
	if len(calcHeuristicCost) == 0 {
		println("need a heuristic !")
		flag.PrintDefaults()
		os.Exit(1)
	}
	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
	_ = file.Close()
	env.fileName = *filePtr
}

func main() {

	arg()
	parse(env.fileName)

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
