/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 13:48:27 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/26 13:54:34 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func main() {
<<<<<<< HEAD

	size := 5
	tab := make([][]int, size)
	for i := 0; i < size; i++ {
		tab[i] = make([]int, size);
	}

    count := 1;
    i := 0;
    j := 0;

    for count < size * size {
        for j < size {
            if (tab[i][j] == 0 && count < size * size) {
                tab[i][j] = count;
                count++;
            }
            if (j + 1 >= size || ( j + 1 < size && tab[i][j+1] != 0)) {
                break
            }
            j++;
        }
        for i < size {
            if (tab[i][j] == 0 && count < size * size) {
                tab[i][j] = count;
                count++;
            }
            if (i + 1 >= size || ( i + 1 < size && tab[i + 1][j] != 0)) {
                break
            }
            i++;
        }
        for j >= 0 {
            if (tab[i][j] == 0 && count < size * size) {
                tab[i][j] = count;
                count++;
            }
            if (j - 1 < 0 || ( j - 1 >= 0 && tab[i][j - 1] != 0)) {
                break
            }
            j--;
        }
        for i >= 0 {
            if (tab[i][j] == 0 && count < size * size) {
                tab[i][j] = count;
                count++;
            }
            if (i - 1 < 0 || ( i - 1 >= 0 && tab[i-1][j] != 0)) {
                break
            }
            i--;
        }
    }

    k := 0
    for k < size{
        l := 0
        for l < size {
            fmt.Printf("%3d", tab[k][l])
            l++
        }
        k++
        fmt.Printf("\n")
    }
=======
	initEnv(6)
	fmt.Println(env)
>>>>>>> 7ccebc3a707c1c06a2466e3bdce483dbdbca9513
}
