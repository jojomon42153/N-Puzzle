/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   header.go                                        .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2019/10/30 17:52:04 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/03/05 18:48:08 by jojomoon    ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
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
// nbOpened => nombre total de states ayant ete deployes
// nbMaxAllocated => nombre maximal de states presents en mémoire au meme moment (max(nbOpenedSet+nbClosedSet))
// nbMoves => nombre de coups totaux de la solution finale
type stats struct {
	nbOpened    int
	nbClosed    int
	nbTotal     int
	nbMaxOpened int
	nbMaxClosed int
	nbEvaluated int
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
	ch["nbOpened"] <- incr       // On incrémente le nombre de states open
	me.tab = append(me.tab, new) // On rajoute le state a la fin de l'arbre
	sorted := false
	newIndex := len(me.tab) - 1 // On enregistre son index
	for !sorted {
		parentIndex := int(math.Floor(float64((newIndex - 1) / 2)))     // On calcule l'index du parent
		if me.tab[newIndex].totalCost < me.tab[parentIndex].totalCost { // Si le coût du parent est plus cher que celui que l'on trie...
			me.tab[newIndex], me.tab[parentIndex] = me.tab[parentIndex], me.tab[newIndex] // ... on échange les 2
			newIndex = parentIndex
		} else if me.tab[newIndex].totalCost == me.tab[parentIndex].totalCost && me.tab[newIndex].initialCost < me.tab[parentIndex].initialCost { // Si les 2 coûts sont égaux mais que l'initialCost du parent est plus élevé...
			me.tab[newIndex], me.tab[parentIndex] = me.tab[parentIndex], me.tab[newIndex] // ... on échange les 2
			newIndex = parentIndex
		} else { // Ici le parent est prioritaire sur notre state donc on coupe le tri
			sorted = true
		}
	}
}

func (me *openedSet) pullLowestCost() *state {
	bestState := me.tab[0] // On sauvegarde le meilleur état
	ch["nbOpened"] <- decr // On décrémente le nombre de states open
	if len(me.tab) == 1 {
		me.tab = make([]*state, 0)
		return bestState
	}
	// On reconstruit le tas en mettant le dernier element a la place du premier
	me.tab[0] = me.tab[len(me.tab)-1]
	me.tab = me.tab[:len(me.tab)-1]
	// Le tas est construit mais il faut le retrier
	sorted := false
	toSortIndex := 0

	for !sorted {
		var bestChildIndex int
		leftChildIndex := toSortIndex*2 + 1 // On calcule l'index des enfants dans le tableau
		rightChildIndex := leftChildIndex + 1

		// On cherche quel noeud enfant a le meilleur score
		if len(me.tab) == leftChildIndex+1 { // Si le tableau s'arrête sur une branche gauche
			bestChildIndex = leftChildIndex
		} else if len(me.tab) < leftChildIndex+1 { // Sinon s'il n'y a pas de branche en dessous
			bestChildIndex = toSortIndex
		} else if me.tab[leftChildIndex].totalCost == me.tab[rightChildIndex].totalCost { // Si les 2 branches enfant ont le même coût total, on prend celui qui a le moindre initialCost
			if me.tab[rightChildIndex].initialCost <= me.tab[leftChildIndex].initialCost {
				bestChildIndex = rightChildIndex
			} else {
				bestChildIndex = leftChildIndex
			}
		} else if me.tab[leftChildIndex].totalCost < me.tab[rightChildIndex].totalCost { // Sinon si la branche gauche a le moindre coût
			bestChildIndex = leftChildIndex
		} else { // Ici la branche droite a le moindre coût
			bestChildIndex = rightChildIndex
		}
		// On echange le noeud actuel avec celui de la meilleure branche si elle a un coût moindre ou on arrête le tri
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
	fileName   string
	openedSet  *openedSet
	closedSet  map[string]*state
	allSets    map[string]*state
	size       int
	finalState *state
	stats      stats
}

var ch map[string]chan int
var calcHeuristicCost []func(*state)

const (
	incr = 0
	decr = 1
	exit = 2
)
