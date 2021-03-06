func openLock(deadends []string, target string) int {
	deads := make(map[string]string)
	for _, dead := range deadends {
		deads[dead] = "ok"
	}
	visited := make(map[string]string)
	queue := []string{"0000"}
	var level int
	for len(queue) > 0 {
		qSize := len(queue)
		for i := 0; i < qSize; i++ {
			val := queue[0]
			queue = queue[1:]
			if _, ok := deads[val]; ok {
				continue // 处理deads， 结束这个Loop
			}
			if val == target {
				return level
			}
			for j := 0; j < 4; j++ { // 开始bfs
				next := getNextOrPrevValue(val[:j], val[j+1:], val[j], false) //  处理这个情况
				if _, ok := visited[next]; !ok {
					queue = append(queue, next)
					visited[next] = "ok"
				}
				prev := getNextOrPrevValue(val[:j], val[j+1:], val[j], true) // 往前往后走下
				if _, ok := visited[prev]; !ok {
					queue = append(queue, prev)
					visited[prev] = "ok"
				}
			}
		}
		level++
	}

	return -1
}

func getNextOrPrevValue(prefix, postfix string, char byte, prev bool) string {
	if prev {
		if char == '0' {
			char = '9'
		} else {
			char--
		}
	} else {
		if char == '9' {
			char = '0'
		} else {
			char++
		}
	}
	return prefix + string(char) + postfix
}