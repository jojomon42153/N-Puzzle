/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parse.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 15:19:49 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/28 19:45:15 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func checkSolvability(initial *state) bool {

	inv := 0 //number of inversions
	sqSize := env.size * env.size
	tab := initial.state1D
	zeroPos := env.size*(initial.zeroCoord.y) + initial.zeroCoord.x

	for i := 0; i < sqSize; i++ {
		for j := i + 1; j < sqSize; j++ {
			if tab[i] != 0 && tab[j] != 0 && tab[i] > tab[j] {
				inv++
			}
		}
	}

	//Width is odd
	if env.size%2 == 1 {
		return inv%2 == 1 // True if number of inversions is odd, if even false
	} else if (zeroPos/env.size)%2 == 0 { //width is even, 0 is on an even row from bottom
		fmt.Println("else if", inv%1)
		return inv%2 == 1 // True if number of inversions is odd, if even false
	} else {
		fmt.Println("else", inv%2)
		return inv%2 == 0
	}
}

func fillLines(str [][]byte, twoD [][]int, lines int) [][]int {
	i := 0
	if lines == 0 {
		twoD = make([][]int, env.size)
	}
	for _, elem := range str {
		i, _ = strconv.Atoi(string(elem))
		twoD[lines] = append(twoD[lines], i)
	}
	return twoD
}

func parse(fileName string) int {
	env.size = -1
	var twoD [][]int
	re := regexp.MustCompile(`[#][\S ]*`)   // enleve les commentaires
	re2 := regexp.MustCompile(`[^0-9 $]`)   // verifie qu'il n'y ait pas des caracteres de merde
	re3 := regexp.MustCompile(`^([0-9]+)$`) // recupere la size du taquin
	re4 := regexp.MustCompile(`([0-9]+)`)   // recupere une ligne du taquin
	file, err := os.Open(fileName)
	lines := 0
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //lecture du file
		text := []byte(scanner.Text())
		str := re.ReplaceAll(text, []byte(""))
		if string(str) != "" { // si la ligne n'est pas un commentaire
			str2 := re2.ReplaceAll(str, []byte(""))
			if lines == env.size {
				log.Fatal("not well formated: too much lines")
				twoD = nil
				os.Exit(1)
			}
			if !strings.EqualFold(string(str2), string(str)) { //si la ligne contient un caractere indesirable
				log.Fatal("not well formated")
				twoD = nil
				os.Exit(1)
			}
			if env.size == -1 { //si la taille du taquin n'a pas encore été set
				str3 := re3.Find(str2)
				if string(str3) != "" { //si la ligne contient un simple nombre
					env.size, err = strconv.Atoi(string(str3))
				}
			} else {
				str4 := re4.FindAll(str2, -1)
				if len(str4) != env.size {
					log.Fatal("not well formated")
					twoD = nil
					os.Exit(1)
				}
				twoD = fillLines(str4, twoD, lines)
				lines++
			}
		}
	}

	//Print twoD tab
	if len(twoD) != env.size {
		twoD = nil
		log.Fatal("Lines missing or wrong size")
		os.Exit(1)
	}
	print("print twoD:\n")
	for _, elem := range twoD {
		for _, e := range elem {
			print(e)
			print(" ")
		}
		print("\n")
	}

	var oneD []int
	maps := make(map[int]bool)
	max := env.size * env.size
	for _, elem := range twoD {
		for _, e := range elem {
			if _, ok := maps[e]; !ok && e < max {
				maps[e] = true
				oneD = append(oneD, e)
			}
		}
	}
	print("print oneD:\n")
	for _, e := range oneD {
		print(e)
		print(" ")
	}
	//Print oneD tabd
	print("\n")
	if len(oneD) != max {
		log.Fatal("wrong format: uncorrect numbers\n")
		twoD = nil
		oneD = nil
		os.Exit(1)
	}
	print("what do we return ?\n")
	createFirstState(twoD)
	return 6
}
