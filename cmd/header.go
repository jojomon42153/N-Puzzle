/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   header.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:52:04 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/28 15:12:59 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"math"
)

type coord struct {
	x int
	y int
}

// Structure des etats
// parent => pointe sur le state precedent
// state2D => tableau a 2 dimensions du state
// state1D => tableau a 1 dimension du state
// initialCost => nombre de coups pour arriver a ce state
// heuristicCost => cout determine par l'heurisitque
// totalCost => somme de heuristicCost et initialCost
// isOpen => true si ce state est dans l'openSet
type state struct {
	parent        *state
	index         string
	state2D       [][]int
	state1D       []int
	coord         []coord
	zeroCoord     coord
	initialCost   int
	heuristicCost int
	totalCost     int
	isOpen        bool
}

// Structure utilisée pour stocker les informations statistiques generales
// nbOpened => nombre total de states ayant ete dans l'openedSet
// nbClosed => nombre maximal de states presents en mémoire au meme moment (max(nbOpenedSet+nbClosedSet))
// nbMoves => nombre de coups totaux de la solution finale
type stats struct {
	nbOpened    int
	nbClosed    int
	nbTotal     int
	nbMaxOpened int
	nbMaxTotal  int
	nbMoves     int
}

// Structure de la priority queue (file de priorité)
// tab => tableau des openedStates triés par la methode heapsort (sur le tas)
// isEmpty => retourne true si la file est vide
// insertWithCostPriority => insère dans la file en triant du totalCost minimum au totalCost maximum par un heapsort
// pullLowestCost => retire le premier element de la liste (ici le moins cher donc) et reconstruit le tas
type openedSet struct {
	tab []*state
}

func (me *openedSet) isEmpty() bool {
	return len(me.tab) == 0
}

func (me *openedSet) insertWithCostPriority(new *state) {
	me.tab = append(me.tab, new)
	sorted := false
	for !sorted {
		newIndex := len(me.tab) - 1
		parentIndex := int(math.Floor(float64((newIndex - 1) / 2)))
		if me.tab[newIndex].totalCost < me.tab[parentIndex].totalCost {
			me.tab[newIndex], me.tab[parentIndex] = me.tab[parentIndex], me.tab[newIndex]
		} else {
			sorted = true
		}
	}
}

func (me *openedSet) pullLowestCost() *state {
	bestState := me.tab[0]
	if len(me.tab) == 1 {
		me.tab = make([]*state, 0)
		return bestState
	}
	// Now we will construct back the "tree"
	me.tab[0] = me.tab[len(me.tab)-1]
	me.tab = me.tab[:len(me.tab)-1]
	// The tree is constructed but the first value is not sorted
	sorted := false
	toSortIndex := 0

	for !sorted {
		var bestChildIndex int
		leftChildIndex := toSortIndex*2 + 1
		rightChildIndex := leftChildIndex + 1

		if len(me.tab) == leftChildIndex+1 { // Si le tableau sarrete sur une branche gauche
			bestChildIndex = leftChildIndex
		} else if len(me.tab) < rightChildIndex+1 && len(me.tab) != leftChildIndex+1 {
			bestChildIndex = toSortIndex

		} else if me.tab[leftChildIndex].totalCost > me.tab[rightChildIndex].totalCost {
			bestChildIndex = rightChildIndex
		} else {
			bestChildIndex = leftChildIndex
		}
		if me.tab[toSortIndex].totalCost > me.tab[bestChildIndex].totalCost {
			me.tab[toSortIndex], me.tab[bestChildIndex] = me.tab[bestChildIndex], me.tab[toSortIndex]
			toSortIndex = bestChildIndex
		} else {
			sorted = true
		}
	}
	return bestState
}

// Structure utilisee pour stocker toute information generale relative au projet
// openedSet => map de hashage de tous les states dans l'état opened
// closedSet => map de hashage de tous les states dans l'état opened
// finalState2D => tableau 2D definissant le final state
// finalState1D => tableau 1D definissant le final state
// finalCoord => tableau de pointeur sur coordonnee definissant le final state
// stats => structure stats
var env struct {
	openedSet  *openedSet
	closedSet  map[string]*state
	allSets    map[string]*state
	size       int
	finalState *state
	stats      stats
}

var calcHeuristicCost func(*state)
var ch map[string]chan int

const (
	incr = 0
	decr = 1
	exit = 2
)
