/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   utils.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 19:25:44 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/27 13:55:48 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"strings"
)

func arrayToString(a []int, delim string) string {
	return strings.Replace(fmt.Sprint(a), " ", delim, -1)
}

func array2Dto1D(arr2D [][]int) []int {
	var tab []int = make([]int, 0)
	for _, line := range arr2D {
		tab = append(tab, line...)
	}
	return tab
}

func arr2DtoCoord(arr2D [][]int) []coord {
	size := env.size * env.size
	var tab []coord = make([]coord, size)
	for y, line := range arr2D {
		for x, tile := range line {
			tab[tile] = coord{x, y}
		}
	}
	return tab
}

func duplicateState2D(original [][]int) [][]int {
	duplicate := make([][]int, len(original))
	for i := range original {
		duplicate[i] = make([]int, len(original[i]))
		copy(duplicate[i], original[i])
	}
	return duplicate
}

func getInitialZeroCoord(size int) coord {
	n := size
	coord := coord{0, 0}
	if n%2 == 0 {
		coord.x--
	} else {
		n--
	}
	coord.x += n / 2
	coord.y += n / 2
	return coord
}

func searchZeroCoord(state [][]int) coord {
	for y, line := range state {
		for x, tile := range line {
			if tile == 0 {
				return coord{x, y}
			}
		}
	}
	return coord{-1, -1}
}
