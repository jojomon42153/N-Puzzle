/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   stats.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/28 14:15:43 by jmonneri          #+#    #+#             */
/*   Updated: 2020/03/05 19:22:54 by jojomoon         ###   ########lyon.fr   */
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
