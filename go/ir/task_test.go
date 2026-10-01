package ir

import "testing"

func TestTaskWaitMarksDependenciesTransitivelyDone(t *testing.T) {
	leaf := &task{done: make(chan unit)}
	child := &task{done: make(chan unit)}
	root := &task{done: make(chan unit)}
	child.addEdge(leaf)
	root.addEdge(child)
	leaf.markDone()
	child.markDone()
	root.markDone()

	root.wait()
	for name, task := range map[string]*task{"root": root, "child": child, "leaf": leaf} {
		if !task.isTransitivelyDone() {
			t.Errorf("%s was not marked transitively done", name)
		}
	}
}
