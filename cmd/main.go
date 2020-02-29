/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 13:48:27 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/29 00:44:19 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

func main() {
	ch = make(map[string]chan int, 2)
	ch["nbOpened"] = make(chan int)
	ch["nbClosed"] = make(chan int)
	go updateNbOpened()
	go updateNbClosed()

	calcHeuristicCost = manhattan
	parse("ressources/correctInput/taquin_dim4_1.txt")

	aStar()

	for _, chanel := range ch {
		close(chanel)
	}
}
