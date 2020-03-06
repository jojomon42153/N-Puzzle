/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   stats.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: gaennuye <gaennuye@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/28 14:15:43 by gaennuye          #+#    #+#             */
/*   Updated: 2020/03/06 16:16:08 by gaennuye         ###   ########lyon.fr   */
/*                                                                            */
/* ************************************************************************** */

package main

func updateNbOpened() {
	for true {
		if code := <-ch["nbOpened"]; code == incr {
			env.stats.nbOpened++
			env.stats.nbTotal++
			if env.stats.nbMaxOpened < env.stats.nbOpened {
				env.stats.nbMaxOpened = env.stats.nbOpened
			}
		} else if code == decr {
			env.stats.nbOpened--
			env.stats.nbTotal--
		} else if code == exit {
			return
		}
	}
}

func updateNbClosed() {
	for true {
		if code := <-ch["nbClosed"]; code == incr {
			env.stats.nbClosed++
			if env.stats.nbMaxClosed < env.stats.nbClosed {
				env.stats.nbMaxClosed = env.stats.nbClosed
			}
		} else if code == decr {
			env.stats.nbClosed--
		} else if code == exit {
			return
		}
	}
}
