/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parse.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: gaennuye <gaennuye@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 15:19:49 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/28 17:19:38 by gaennuye         ###   ########lyon.fr   */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func parse() int {
	return 3
}

func checkSolvability(initial *state) bool {

	// fmt.Println(env.finalState.state1D)
	inv := 0 //number of inversions
	sqSize := env.size * env.size
	// test1 := []int {1,3,7,4,8,6,5,2,0}
	tab := initial.state1D
	// final = env.finalState.state1D
	zeroPos := (initial.zeroCoord.x + 1) * (initial.zeroCoord.y + 1) - 1
	fmt.Println(initial.state1D)
	
	fmt.Println(" zeropos", zeroPos)

	for i := 0; i < sqSize; i++ {
		for j := i + 1; j < sqSize; j++ {
			fmt.Println("i : ", i, " j : ", j, "tab[i]", tab[i], "tab[j]", tab[j])
			if (tab[i] != 0 && tab[j] != 0 && tab[i] > tab[j]) {
				fmt.Println("cond")
				inv++
			}
		}
	}
	fmt.Println("inv", inv)
	fmt.Println(zeroPos, env.size, zeroPos/env.size)

	//Width is odd
	if env.size % 2 == 1 {
		return inv % 2 == 1 // True if number of inversions is odd, if even false
	} else if (zeroPos / env.size) % 2 == 0 {  //width is even, 0 is on an even row from bottom
		fmt.Println("else if", inv % 1)
		return inv % 2 == 1
	} else {
		fmt.Println("else", inv % 2)
		return inv % 2 == 0
	}
}
