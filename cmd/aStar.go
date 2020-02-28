/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   aStar.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 15:23:56 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/27 23:05:28 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

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
	// move tile
	return &childMove
}

func generateNextMoves(current *state) []*state {
	var childs []*state
	nbchilds := countNbChilds(current.zeroCoord)
	x0, y0 := current.zeroCoord.x, current.zeroCoord.y
	childs = make([]*state, nbchilds)

	i := 0
	if x0 != 0 { // moveLeftTile
		childs[i] = moveTile(coord{x0 - 1, y0}, current.zeroCoord, current)
		i++
	}
	if x0 != env.size-1 { // moveRightTile
		childs[i] = moveTile(coord{x0 + 1, y0}, current.zeroCoord, current)
		i++
	}
	if y0 != 0 { // moveUpTile
		childs[i] = moveTile(coord{x0, y0 - 1}, current.zeroCoord, current)
		i++
	}
	if y0 != env.size-1 { // moveDownTile
		childs[i] = moveTile(coord{x0, y0 + 1}, current.zeroCoord, current)
	}
	return childs
}

func moveToClosed(toMove *state) {
	toMove.isOpen = false
	env.closedSet[toMove.index] = toMove
}

func aStar() {
	success := false
	for !(env.openedSet.isEmpty()) || success == false {
		var current = env.openedSet.pullLowestCost()
		if current.heuristicCost == 0 { // If current is the solution
			env.finalState = current
			success = true
			break
		}
		var childs []*state = generateNextMoves(current)
		moveToClosed(current)
		for _, child := range childs {
			if _, ok := env.allSets[child.index]; !ok { // If child is a new state, calc heuristic and sort it in openedSet
				calcHeuristicCost(child)
				env.openedSet.insertWithCostPriority(child)
				env.allSets[child.index] = child
			} else if _, ok := env.closedSet[child.index]; ok { // If child was already opened, ignore this child
				child = nil
			} else if val, ok := env.allSets[child.index]; ok && val.isOpen { // If child is already in openset, keep the one with better score
				calcHeuristicCost(child)
				if val.totalCost > child.totalCost {
					val.parent = child.parent
					val.totalCost = child.totalCost
				}
				child = nil
			}
		}
	}
	if success {
		fmt.Println("ALLEZ A LA TEUF")
		printSolve(env.finalState)
		fmt.Printf("OpenSet = %d\nCloseSet = %d\n", len(env.openedSet.tab), len(env.closedSet))
	}
}
