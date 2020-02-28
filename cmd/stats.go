/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   stats.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/28 14:15:43 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/28 15:17:13 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

func updateNbOpened() {
	for true {
		if code := <-ch["nbOpened"]; code == incr {
			env.stats.nbOpened++
			env.stats.nbTotal++
		} else if code == decr {
			env.stats.nbOpened--
			env.stats.nbTotal--
		} else if code == exit {
			return
		}
		if env.stats.nbOpened > env.stats.nbMaxOpened {
			env.stats.nbMaxOpened = env.stats.nbOpened
		}
	}
}

func updateNbClosed() {
	for true {
		if code := <-ch["nbClosed"]; code == incr {
			env.stats.nbClosed++
			env.stats.nbTotal++
		} else if code == decr {
			env.stats.nbClosed--
			env.stats.nbTotal--
		} else if code == exit {
			return
		}
	}
}
