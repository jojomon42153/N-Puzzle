/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parse.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: gaennuye <gaennuye@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 15:19:49 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/28 11:50:09 by gaennuye         ###   ########lyon.fr   */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func parse() int {
	return 3
}

func checkSolvability(initial *state, size int) int {

	counter := 0
	test1 := []int{1,8,7,4,5,2,6,3,9}
	test2 := []int{2,6,3,1,7,4,8,5,9}
	test3 := []int{5,2,1,6,8,4,7,3,9}
	// test3 := []int{1,2,3,4,6,5,8,7,9}
	// test4 := []int{13, 2, 3, 12, 9, 11, 1, 10, 16, 6, 4, 14, 15, 8, 7, 5}
	test4 := []int{8,2,3,9,11,12,7,13,15,14,10,6,1,5,4}
	// fake3 := genFakeFinal1D(5)
	fmt.Println("bubble : ", bubblesort(initial.state1D))
	
	fmt.Println("bubble : ", bubblesort(test1))
	fmt.Println("bubble : ", bubblesort(test2))
	fmt.Println("bubble : ", bubblesort(test3))
	fmt.Println("bubble : ", bubblesort(test4))


	return counter
}

func genFakeFinal1D(n int) []int {
	n = n*n
	var tab = make([]int, n)
	fmt.Println(n)
	for i := 0; i < n - 1; i++ {
		tab[i] = i + 1
	}
	tab[n-1] = n
	fmt.Println("fake tab: ", tab)
	return tab
}

func bubblesort(items []int) int {
    var (
        n = len(items)
        sorted = false
	)
	
	count := 0

    for !sorted {
        swapped := false
        for i := 0; i < n-1; i++ {
            if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				count++
                swapped = true
            }
        }
        if !swapped {
            sorted = true
        }
        n = n - 1
	}
	return count
}