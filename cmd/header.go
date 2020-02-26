/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   header.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:52:04 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/26 10:41:36 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

type coord struct {
	x int
	y int
}

// Va etre utilise par A*
type state struct {
	parent        *state
	state2D       [][]int
	state1D       []int
	initialCost   int
	heuristicCost int
	totalCost     int
}

type stats struct {
	nbOpened       int
	nbMaxAllocated int
	nbMoves        int
}

type env struct {
	openedSet    map[string]*state
	closedSet    map[string]*state
	finalCoord   []*coord
	finalState2D [][]int
	finalState1D []int
	stats        *stats
}
