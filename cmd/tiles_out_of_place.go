/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   tiles_out_of_place.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: gaennuye <gaennuye@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 13:36:48 by gaennuye          #+#    #+#             */
/*   Updated: 2020/02/26 13:39:53 by gaennuye         ###   ########lyon.fr   */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

/* Tiles out of place heuristic */

func tilesOutOfPlace(oneD []int, goal []int) int {

	// fmt.Println(oneD)
	// fmt.Println(goal)

	counter := 0

	for i := 0; i < len(goal); i++ {
		if (goal[i] != oneD[i]){
			counter++
		}
	}

	// fmt.Println(counter)

	return(counter)
}