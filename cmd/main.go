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

func arg(){
	manhattanPtr := flag.Bool("hm", false, ": select manhattan as heuristic");
	toopPtr := flag.Bool("ht", false, ": select toop as heuristic");
	euclideanPtr := flag.Bool("he", false, ": select euclidean as heuristic");
	filePtr := flag.String("f", "", ": -f thePuzzle.txt");
	//wordPtr := flag.String("word", "foo", "a string")
	flag.Parse()
	if *filePtr == ""{
		println("need a file !");
		flag.PrintDefaults();
		os.Exit(1);
	}
	//if *manhattanPtr == false && *toopPtr == false && *euclideanPtr == false{
	//	println("need a heuristic ! Only one");
	//	flag.PrintDefaults();
	//	os.Exit(1);
	//}
	if (*manhattanPtr == true){
		calcHeuristicCost = append(calcHeuristicCost, manhattan)
	}
	if (*toopPtr == true){
		calcHeuristicCost = append(calcHeuristicCost, )
	}
	if (*euclideanPtr == true){
	}
	//if (i != 1){
	//	println("need a heuristic ! Only one");
	//	flag.PrintDefaults();
	//	os.Exit(1);
	//}
	fmt.Println("file : ", *filePtr)
	fmt.Println("hm : ", *manhattanPtr)
	fmt.Println("ht : ", *toopPtr)
	fmt.Println("he : ", *euclideanPtr)
	fmt.Println("tail:", flag.Args())

	file, err := os.Open(*filePtr);
	if err != nil {
		log.Fatal(err)
	}
	println(file);

}

func main() {
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
