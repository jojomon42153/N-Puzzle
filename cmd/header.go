/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   header.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:52:04 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/26 10:59:23 by jmonneri         ###   ########.fr       */
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
type state struct {
	parent        *state
	state2D       [][]int
	state1D       []int
	initialCost   int
	heuristicCost int
	totalCost     int
	isOpen        bool
}

// structure utilisée pour stocker les informations statistiques generales
// nbOpened => nombre total de states ayant ete dans l'openedSet
// nbClosed => nombre maximal de states presents en mémoire au meme moment (max(nbOpenedSet+nbClosedSet))
// nbMoves => nombre de coups totaux de la solution finale
type stats struct {
	nbOpened       int
	nbMaxAllocated int
	nbMoves        int
}

// structure utilisee pour stocker toute information generale relative au projet
// openedSet => map de hashage de tous les states dans l'état opened
// closedSet => map de hashage de tous les states dans l'état opened
// finalCoord => tableau de pointeur sur coordonnee definissant le final state
// finalState2D => tableau 2D definissant le final state
// finalState1D => tableau 1D definissant le final state
// stats => pointeur sur la structure stats
var env struct {
	openedSet    map[string]*state
	closedSet    map[string]*state
	size 		 int
	finalCoord   []coord
	finalState2D [][]int
	finalState1D []int
	stats        *stats
}
