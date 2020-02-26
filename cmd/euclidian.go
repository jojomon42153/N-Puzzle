/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   euclidian.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: gaennuye <gaennuye@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 14:19:20 by gaennuye          #+#    #+#             */
/*   Updated: 2020/02/26 15:14:46 by gaennuye         ###   ########lyon.fr   */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"math"
)

func euclidian(twoD [][]int, goal[][]int) float64 {
	
	size := len(goal[0])
	sum := 0.0

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (goal[i][j] != twoD[i][j]) {

				goal_c := finalCoord(twoD[i][j])

				current_err := math.Pow(float64(i - goal_c[0]), 2 ) + math.Pow(float64(j - goal_c[1]), 2) 


				for m := 0; m < size; m++ {
					for n := 0; n < size; n++ {
						if (goal[i][j] == twoD[m][n]) {
							// fmt.Println(goal[i][j])
							current_err := math.Pow(float64(i - m), 2) + math.Pow(float64(j - n), 2)
							// fmt.Println(current_err)
							sum += current_err
						}
					}
				}
			}
		}
	}

	fmt.Println(sum)

	return sum
}