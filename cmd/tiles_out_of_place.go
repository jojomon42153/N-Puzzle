/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   tiles_out_of_place.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: gaennuye <gaennuye@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 13:36:48 by gaennuye          #+#    #+#             */
/*   Updated: 2020/02/27 14:15:07 by gaennuye         ###   ########lyon.fr   */
/*                                                                            */
/* ************************************************************************** */

package main


/* Tiles out of place heuristic */

func old_tilesOutOfPlace(oneD []int) int {

	counter := 0
	sq_size := env.size * env.size

	for i := 0; i < sq_size; i++ {
		if (oneD[i] != env.finalState.state1D[i]){
			counter++
		}
	}

	return counter
}