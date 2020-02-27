/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   header.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:52:04 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/27 10:03:43 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

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
	nbOpened       int
	nbMaxAllocated int
	nbMoves        int
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

func (me *openedSet) insertWithCostPriority() {

}

func (me *openedSet) pullLowestCost() *state {
	return &state{}
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
