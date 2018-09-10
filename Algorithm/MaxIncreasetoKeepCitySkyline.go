package Algorithm

func MaxIncreaseKeepingSkyline(grid [][]int) int {
	//计算地标高度
	var topBottom []int
	var leftRight []int
	for _, v1 := range grid {
		max := 0
		for _, v2 := range v1 {
			if v2 > max {
				max = v2
			}
		}
		leftRight = append(leftRight,max)
	}

	innerLength := len(grid[0])
	for i:=0;i<innerLength;i++ {
		max := 0
		for k1,_ := range grid {
			if grid[k1][i] > max {
				max = grid[k1][i]
			}
		}
		topBottom = append(topBottom,max)
	}

	//生成新的地标图
	var gridNew [][]int
	for _,v1 := range leftRight {
		var gridTmp []int
		for _,v2 := range topBottom {
			if v2 > v1 {
				gridTmp= append(gridTmp,v1)
			} else {
				gridTmp= append(gridTmp,v2)
			}
		}
		gridNew = append(gridNew,gridTmp)
	}

	//计算差额
	balance := 0
	for k1,v1 := range gridNew {
		for  k2,v2 := range v1 {
			if v2 - grid[k1][k2] > 0 {
				balance += v2 - grid[k1][k2]

			}
		}
	}
	return balance
}