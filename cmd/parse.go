/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parse.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 15:19:49 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/26 15:20:16 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"bufio"
	//"fmt"
	"strconv"
	"strings"

	//"fmt"
	"log"
	"os"
	"regexp"
)

func nik(str [][]byte, twoD [][]int) [][]int {
	i := 0
	if len(twoD) == 0 {
		twoD = make([][]int, env.size)

	}
	for _, elem := range str {
		//println(twoD);
		//println(string(elem));
		//println(len(twoD));
		i, _ = strconv.Atoi(string(elem))

		twoD[len(twoD)-1] = append(twoD[len(twoD)-1], i)

	}
	return twoD
}

func parse() int {
	env.size = -1
	var twoD [][]int
	re := regexp.MustCompile(`[#][\S ]*`)   //enleve les commentaires
	re2 := regexp.MustCompile(`[^0-9 $]`)   //verifie qu'il n'y ai pas des caracteres de merde
	re3 := regexp.MustCompile(`^([0-9]+)$`) // recupere la size du taquin
	re4 := regexp.MustCompile(`([0-9]+)`)   // recupere une ligne du taquin
	str := ""
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //lecture du file
		text := []byte(scanner.Text())
		str := re.ReplaceAll(text, []byte(""))
		if string(str) != "" { // si la ligne n'est pas un commentaire
			str2 := re2.ReplaceAll(str, []byte(""))
			if !strings.EqualFold(string(str2), string(str)) { //si la ligne contient un caractere indesirable
				println("not well formated")
				break
			}
			if env.size == -1 { //si la taille du taquin n'a pas encore été set
				str3 := re3.Find(str2)
				if string(str3) != "" { //si la ligne contient un simple nombre
					env.size, err = strconv.Atoi(string(str3))
				}
			} else {
				str4 := re4.FindAll(str2, -1)
				//println(string(str4[0]))
				if len(str4) != env.size {
					println("not well formated")
					break
				}
				twoD = nik(str4, twoD)
				//println(twoD[0][0]);

			}
		}

		//str += scanner.Text();
	}
	print(str)
	return 6
}
