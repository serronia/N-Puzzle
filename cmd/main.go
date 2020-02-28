/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: gaennuye <gaennuye@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 13:48:27 by jmonneri          #+#    #+#             */
/*   Updated: 2020/02/28 11:58:13 by gaennuye         ###   ########lyon.fr   */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func createDataSet(n int) {
	var initial *state = &state{
		parent:        nil,
		state2D:       make([][]int, n),
		state1D:       make([]int, 0),
		coord:         nil,
		initialCost:   0,
		heuristicCost: 0,
		totalCost:     0,
		isOpen:        true,
	}
	// init state 2D
	// 35 34 33 32 31 30
	// 29 28 27 26 25 24
	// 23 22 21 20 19 18
	// 17 16 15 14 13 12
	// 11 10 9  8  7  6
	// 5  4  3  2  1  0
	for i := 0; i < n; i++ {
		initial.state2D[i] = make([]int, n)
	}
	var i int = n*n - 1
	for y, line := range initial.state2D {
		for x := range line {
			initial.state2D[y][x] = i
			i--
		}
	}
	initial.state2D[0][0] = 6
	initial.state2D[0][1] = 8
	initial.state2D[0][2] = 1
	initial.state2D[1][0] = 2
	initial.state2D[1][1] = 7
	initial.state2D[1][2] = 4
	initial.state2D[2][0] = 3
	initial.state2D[2][1] = 5
	initial.state2D[2][2] = 0
	initial.zeroCoord = searchZeroCoord(initial.state2D)
	// init state1D
	for _, line := range initial.state2D {
		initial.state1D = append(initial.state1D, line...)
	}
	initial.index = arrayToString(initial.state1D, ",")
	calcHeuristicCost = manhattan
	calcHeuristicCost(initial)

	// fmt.Println(initial.state1D)
	// fmt.Println(env.finalState.state1D)

	fmt.Println(checkSolvability(initial, n))

	tilesOutOfPlace(initial)
	euclidian(initial)
	env.openedSet.tab[0] = initial
	env.allSets[initial.index] = initial
}

func main() {
	n := parse()
	n = 3
	initEnv(n)
	createDataSet(n)

	// var test3 = []int {1,2,3,0,8,7,4,6,5}
	// fmt.Println(env.finalState.state1D)
	// euclidian(test3)
	// aStar()
}
