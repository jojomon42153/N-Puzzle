/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   init.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 11:03:41 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/26 11:53:37 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

func initFinalState2D(n int) [][]int {
	var tab [][]int
	return tab
}

func initFinalState1D(n int) []int {
	var tab []int
	return tab
}

func initFinalCoord(n int) []*coord {
	var tab []*coord
	return tab
}

func initEnv(n int) {
	env.openedSet = make(map[string]*state)
	env.closedSet = make(map[string]*state)
	env.finalState2D = initFinalState2D(n)
	env.finalState1D = initFinalState1D(n)
	env.finalCoord = initFinalCoord(n)
	env.stats = stats{0, 0, 0}
}

func initState() {

}
