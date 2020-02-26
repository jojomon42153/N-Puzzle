/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   init.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 11:03:41 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/26 13:52:13 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

func initFinalState2D() [][]int {
	size := env.size
	tab := make([][]int, size)
	for i := 0; i < size; i++ {
		tab[i] = make([]int, size)
	}

	count := 1
	i := 0
	j := 0

	for count < size*size {
		for j < size {
			if tab[i][j] == 0 && count < size*size {
				tab[i][j] = count
				count++
			}
			if j+1 >= size || (j+1 < size && tab[i][j+1] != 0) {
				break
			}
			j++
		}
		for i < size {
			if tab[i][j] == 0 && count < size*size {
				tab[i][j] = count
				count++
			}
			if i+1 >= size || (i+1 < size && tab[i+1][j] != 0) {
				break
			}
			i++
		}
		for j >= 0 {
			if tab[i][j] == 0 && count < size*size {
				tab[i][j] = count
				count++
			}
			if j-1 < 0 || (j-1 >= 0 && tab[i][j-1] != 0) {
				break
			}
			j--
		}
		for i >= 0 {
			if tab[i][j] == 0 && count < size*size {
				tab[i][j] = count
				count++
			}
			if i-1 < 0 || (i-1 >= 0 && tab[i-1][j] != 0) {
				break
			}
			i--
		}
	}
	return tab
}

func initFinalState1D() []int {
	var tab []int = make([]int, 0)
	for _, line := range env.finalState2D {
		tab = append(tab, line...)
	}
	return tab
}

func initFinalCoord() []coord {
	size := env.size * env.size
	var tab []coord = make([]coord, size)
	for y, line := range env.finalState2D {
		for x, tile := range line {
			tab[tile] = coord{x, y}
		}
	}
	return tab
}

func initEnv(n int) {
	env.size = n
	env.openedSet = make(map[string]*state)
	env.closedSet = make(map[string]*state)
	env.finalState2D = initFinalState2D()
	env.finalState1D = initFinalState1D()
	env.finalCoord = initFinalCoord()
	env.stats = stats{0, 0, 0}
}

func initState() {

}
