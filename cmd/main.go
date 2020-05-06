/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jojomoon <jojomoon@student.42lyon.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 13:48:27 by gaennuye          #+#    #+#             */
/*   Updated: 2020/05/06 19:03:08 by jojomoon         ###   ########lyon.fr   */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func arg() bool {
	greedyPtr := flag.Bool("gs", false, ": select greedySearch as algorithm (Astar by default)")
	manhattanPtr := flag.Bool("hm", false, ": select manhattan as heuristic")
	toopPtr := flag.Bool("ht", false, ": select toop as heuristic")
	euclideanPtr := flag.Bool("he", false, ": select euclidean as heuristic")
	filePtr := flag.String("f", "", ": -f thePuzzle.txt")
	flag.Parse()
	if *filePtr == "" {
		fmt.Println("Usage of npuzzle:")
		flag.PrintDefaults()
		log.Fatal("Argument error: File needed")
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
		fmt.Println("No heuristics are selected, so calcs will be done with all heurisitcs")
		calcHeuristicCost = append(calcHeuristicCost, manhattan)
		calcHeuristicCost = append(calcHeuristicCost, tilesOutOfPlace)
		calcHeuristicCost = append(calcHeuristicCost, euclidian)
	}
	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
	_ = file.Close()
	env.fileName = *filePtr
	return *greedyPtr
}

func main() {

	ch = make(map[string]chan int, 2)
	ch["nbOpened"] = make(chan int)
	ch["nbClosed"] = make(chan int)
	go updateNbOpened()
	go updateNbClosed()

	greedy := arg()
	parse(env.fileName)
	if !greedy {
		aStar()
	} else {
		greedySearch()
	}
	for _, chanel := range ch {
		chanel <- exit
		close(chanel)
	}
}
