package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math"
	"mylib/graph"
	"net/http"
	"time"
)

type Greeting struct {
	Message string
	Time    time.Time
}

type RequestData struct {
	Num   int
	Paths []Paths
	Start int
	End   int
}

type Paths struct {
	From int
	To   int
	Dist int
}

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/sample.html")
	m := Greeting{
		Message: "Hello Go lang!",
		Time:    time.Now(),
	}
	t.Execute(w, m)
}

func dijkstra(data RequestData) {
	n := data.Num
	paths := data.Paths
	start := data.Start
	end := data.End

	visited := make([]bool, n)
	dist := make([]int, n)

	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt64
	}

	dist[start] = 0
	var pq graph.PQ
	var dummy graph.Node
	dummy.Id = -1
	dummy.Dist = -1
	pq = append(pq, dummy)

	var node graph.Node
	node.Id = start
	node.Dist = dist[start]
	pq = append(pq, node)

	for len(pq) > 1 {
		node = graph.Extract(&pq)
		visited[node.Id] = true
		for i := 0; i < len(paths); i++ {
			if paths[i].From == node.Id {
				to := paths[i].To
				if visited[to] {
					continue
				}
				if dist[to] > dist[node.Id]+paths[i].Dist {
					dist[to] = dist[node.Id] + paths[i].Dist
					var toNode graph.Node
					toNode.Id = to
					toNode.Dist = dist[to]
					graph.Insert(&pq, toNode)

				}
			}
		}

	}
}

func calc(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var data RequestData
	json.Unmarshal(body, &data)

	dijkstra(data)

}

func main() {
	// cssの読込
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/sample", calc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe failed.", err)
	}
}
