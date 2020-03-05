/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   aStar.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 15:23:56 by jmonneri          #+#    #+#             */
/*   Updated: 2020/03/05 19:22:22 by jojomoon         ###   ########lyon.fr   */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

/*
	Evalue le coût heuristique du state en le passant dans toutes les fonctions heuristiques demandées en paramètre
*/
func calculHeuristique(state *state) {
	for _, function := range calcHeuristicCost {
		function(state)
	}
}

/*
	Compte le nombre de coups possibles via la position du zéro dans le taquin
*/
func countNbChilds(zeroCoord coord) int {
	ret := 2
	if zeroCoord.x != 0 && zeroCoord.x != env.size-1 {
		ret++
	}
	if zeroCoord.y != 0 && zeroCoord.y != env.size-1 {
		ret++
	}
	return ret
}

/*
	Crée un state a partir d'un state existant en échangeant les 2 tuiles passées en paramètre
*/
func moveTile(originCoord coord, destCoord coord, model *state) *state {
	var a coord = originCoord
	var b coord = destCoord
	var state2D = duplicateState2D(model.state2D)
	state2D[a.y][a.x], state2D[b.y][b.x] = state2D[b.y][b.x], state2D[a.y][a.x]
	var state1D = array2Dto1D(state2D)
	var childMove state = state{
		parent:        model,
		state2D:       state2D,
		state1D:       state1D,
		coord:         nil,
		zeroCoord:     originCoord,
		initialCost:   model.initialCost + 1,
		heuristicCost: 0,
		totalCost:     0,
		index:         arrayToString(state1D, ","),
		isOpen:        true,
	}
	return &childMove
}

/*
	Génère tous les coups possibles à partir du state envoyé en paramètre
*/
func generateNextMoves(current *state) []*state {
	var childs []*state
	nbchilds := countNbChilds(current.zeroCoord)
	x0, y0 := current.zeroCoord.x, current.zeroCoord.y
	childs = make([]*state, nbchilds)

	i := 0
	if x0 != 0 { // Bouge la tuile a gauche de la tuile 0
		childs[i] = moveTile(coord{x0 - 1, y0}, current.zeroCoord, current)
		i++
	}
	if x0 != env.size-1 { // Bouge la tuile a droite de la tuile 0
		childs[i] = moveTile(coord{x0 + 1, y0}, current.zeroCoord, current)
		i++
	}
	if y0 != 0 { // Bouge la tuile en haut de la tuile 0
		childs[i] = moveTile(coord{x0, y0 - 1}, current.zeroCoord, current)
		i++
	}
	if y0 != env.size-1 { // Bouge la tuile en bas de la tuile 0
		childs[i] = moveTile(coord{x0, y0 + 1}, current.zeroCoord, current)
	}
	return childs
}

/*
	Passe un state dans le closedSet
*/
func moveToClosed(toMove *state) {
	ch["nbClosed"] <- incr // On incremente la stat du nombre de states closed
	toMove.isOpen = false
	env.closedSet[toMove.index] = toMove
}

func aStar() {
	success := false
	for !(env.openedSet.isEmpty()) || success == false { // Tant que l'openSet n'est pas vide ou que la solution n'est pas trouvée
		var current = env.openedSet.pullLowestCost() // On prend la meilleure solution de l'openSet
		if current.heuristicCost == 0 {              // Si on a trouvé la solution, sortir
			env.finalState = current
			success = true
			break
		}
		env.stats.nbEvaluated++
		var childs []*state = generateNextMoves(current) // On génere les coups suivants
		moveToClosed(current)                            // On passe le state évalué dans le closedSet
		for _, child := range childs {                   // Pour chaque state enfant...
			if previous, ok := env.allSets[child.index]; !ok { // Si le state enfant n'appartient a aucun set...
				calculHeuristique(child)
				env.openedSet.insertWithCostPriority(child) // On le rajoute en le triant dans l'openSet
				env.allSets[child.index] = child
			} else if previous.initialCost > child.initialCost { // Si l'initialCost du state previous (déjà existant) est plus cher que le nouveau...
				previous.initialCost = child.initialCost // ... on copie l'initialCost du nouveau dans le previous...
				previous.parent = child.parent           // ... et pareil avec le parent
				if !previous.isOpen {                    // Si le previous est dans le closedSet...
					delete(env.closedSet, previous.index) // ... on l'enleve du closedSet...
					ch["nbClosed"] <- decr
					previous.isOpen = true
					env.openedSet.insertWithCostPriority(child) // ... et on le met dans l'openSet
					env.allSets[previous.index] = previous
				}
			}
		}
	}
	if success {
		printSolve(env.finalState)
		fmt.Printf("OpenSet = %d\nCloseSet = %d\n", len(env.openedSet.tab), len(env.closedSet))
		fmt.Println("ALLEZ A LA TEUF")
	}
}
