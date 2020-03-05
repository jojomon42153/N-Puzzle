/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   main.go                                          .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2020/02/26 13:48:27 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/03/05 17:10:06 by jojomoon    ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
/* ************************************************************************** */
package main

import (
	"flag"
	"log"
	"os"
)

func calculHeuristique(state *state) {
	for _, function := range calcHeuristicCost {
		function(state)
	}
}

func arg() {
	manhattanPtr := flag.Bool("hm", false, ": select manhattan as heuristic")
	toopPtr := flag.Bool("ht", false, ": select toop as heuristic")
	euclideanPtr := flag.Bool("he", false, ": select euclidean as heuristic")
	filePtr := flag.String("f", "", ": -f thePuzzle.txt")
	flag.Parse()
	if *filePtr == "" {
		log.Fatal("Argument error: File needed")
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
}

func main() {

	ch = make(map[string]chan int, 2)
	ch["nbOpened"] = make(chan int)
	ch["nbClosed"] = make(chan int)
	go updateNbOpened()
	go updateNbClosed()

	arg()
	parse(env.fileName)
	aStar()
	for _, chanel := range ch {
		chanel <- exit
		close(chanel)
	}
}
