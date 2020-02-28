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
	"strconv"
	"strings"
	"log"
	"os"
	"regexp"
)

func fillLines(str [][]byte, twoD [][]int, lines int) [][]int {
	i := 0
	if lines  == 0 {
		twoD = make([][]int, env.size)
	}
	for _, elem := range str {
		i, _ = strconv.Atoi(string(elem))
		twoD[lines] = append(twoD[lines], i)
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
	file, err := os.Open("test.txt")
    lines := 0
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
            if lines == env.size{
				println("not well formated: too much lines")
                twoD = nil;
                return 0;
            }
			if !strings.EqualFold(string(str2), string(str)) { //si la ligne contient un caractere indesirable
				println("not well formated")
                twoD = nil;
				return 0
			}
			if env.size == -1 { //si la taille du taquin n'a pas encore été set
				str3 := re3.Find(str2)
				if string(str3) != "" { //si la ligne contient un simple nombre
					env.size, err = strconv.Atoi(string(str3))
				}
			} else {
				str4 := re4.FindAll(str2, -1)
				if len(str4) != env.size {
					println("not well formated")
                    twoD = nil;
					return 0
				}
				twoD = fillLines(str4, twoD, lines)
                lines++
			}
		}
	}

    //Print twoD tab
    if len(twoD) != env.size{
        twoD = nil;
        println("Lines missing or wrong size")
        return 0
    }
    print("print twoD:\n")
    for _, elem := range twoD {
        for _, e := range elem {
            print(e)
            print(" ")
        }
        print("\n")
    }

   var oneD []int; 
   maps := make(map[int] bool)
   max := env.size * env.size
    for _ , elem := range twoD {
        for _, e := range elem {
            if _, ok := maps[e]; !ok && e < max {
                maps[e] = true
                oneD = append(oneD, e)
            }
        }
    }
    print("print oneD:\n")
    for _, e := range oneD {
            print(e)
            print(" ")
    }
    //Print oneD tab
    print("\n")
    if len(oneD) != max{
        print("wrong format: uncorrect numbers\n")
        twoD = nil;
        oneD = nil;
        return 0;
    }
	print("what do we return ?\n")
    return 6
}
