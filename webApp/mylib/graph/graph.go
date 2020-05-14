package graph

type Node struct {
	Id   int
	Dist int
}

type PQ []Node

func Insert(pq *PQ, node Node) {
	*pq = append(*pq, node)
	i := len(*pq) - 1
	for i > 1 && (*pq)[i/2].Dist > (*pq)[i].Dist {
		tmp := (*pq)[i]
		(*pq)[i] = (*pq)[i/2]
		(*pq)[i/2] = tmp
		i = i / 2
	}
}

func Extract(pq *PQ) Node {
	node := (*pq)[1]
	(*pq)[1] = (*qp)[len(*pq)-1]
	*pq = (*qp)[len(*pq)-1]
	update(*pq, 1)
	return node
}

func update(pq *PQ, index int) {
	var min int
	l := 2 * index
	r := 2*index + 1

	if l <= len(pq)-1 && pq[l].Dist < pq[index].Dist {
		min = l
	} else {
		min = index
	}

	if r <= len(pq)-1 && pq[r].Dist < pq[min].Dist {
		min = r
	}
	if min != index {
		tmp := pq[index]
		pq[index] = pq[min]
		pq[min] = tmp
		update(pq, min)
	}
}
