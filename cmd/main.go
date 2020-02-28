/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 13:48:27 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/28 18:38:42 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	calcHeuristicCost = manhattan
	fmt.Println("ici2")
	parse("ressources/correctInput/taquin_dim4_0.txt")

	ch = make(map[string]chan int, 2)
	ch["nbOpened"] = make(chan int)
	ch["nbClosed"] = make(chan int)
	go updateNbOpened()
	go updateNbClosed()

	aStar()

	for _, chanel := range ch {
		close(chanel)
	}
}
