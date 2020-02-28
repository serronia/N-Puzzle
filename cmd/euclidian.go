/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   euclidian.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: gaennuye <gaennuye@student.le-101.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/02/26 14:19:20 by gaennuye          #+#    #+#             */
/*   Updated: 2020/02/27 14:15:00 by gaennuye         ###   ########lyon.fr   */
/*                                                                            */
/* ************************************************************************** */

package main

import "math"


/* Euclidian heuristic */

func old_euclidian(twoD [][]int) int {

	sum := 0.0

	for i := 0; i < env.size; i++ {
		for j := 0; j < env.size; j++ {
			if (twoD[i][j] != env.finalState.state2D[i][j]) {

				goal_c := env.finalState.coord[twoD[i][j]]
				current_err := math.Pow(float64(i - goal_c.x), 2 ) + math.Pow(float64(j - goal_c.y), 2) 
				sum += current_err
			}
		}
	}

	return int(sum)
}