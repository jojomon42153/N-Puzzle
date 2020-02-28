/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   init.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: gaennuye <gaennuye@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 11:03:41 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/28 11:51:16 by gaennuye         ###   ########lyon.fr   */
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

func initFinalState() *state {
	var fState state = state{
		parent:        nil,
		state2D:       initFinalState2D(),
		state1D:       nil,
		coord:         nil,
		zeroCoord:	   getInitialZeroCoord(env.size),
		initialCost:   0,
		heuristicCost: 0,
		totalCost:     0,
		isOpen:        false}
	fState.state1D = array2Dto1D(fState.state2D)
	fState.coord = arr2DtoCoord(fState.state2D)
	return &fState
}

func initEnv(n int) {
	env.size = n
	env.openedSet = &openedSet{tab: make([]*state, 1)}
	env.closedSet = make(map[string]*state)
	env.allSets = make(map[string]*state)
	env.finalState = initFinalState()
	env.stats = stats{0, 0, 0}
}
