package garbagecollector

import "fmt"

type Object struct {
	data   string
	marked bool
	refs   []*Object // References to other objects
}

type Heap struct {
	objects []*Object
}

func NewObject(data string) *Object {
	return &Object{data: data, refs: []*Object{}}
}

func (h *Heap) AddObject(o *Object) {
	h.objects = append(h.objects, o)
}

func (h *Heap) mark(o *Object) {
	fmt.Println("o", o.data, o.marked, o.refs)
	if o == nil {
		return
	}
	for _, ref := range o.refs {
		o.marked = true
		ref.marked = true
		h.mark(ref)
	}
}

func (h *Heap) MarkAll() {
	for _, obj := range h.objects {
		if obj.marked {
			continue
		}
		h.mark(obj)
	}
}

func (h *Heap) Sweep() {
	newObjects := []*Object{}
	for _, obj := range h.objects {
		if obj.marked {
			obj.marked = false
			newObjects = append(newObjects, obj)
		}
	}
	h.objects = newObjects
}

func (h *Heap) CollectGarbage() {
	h.MarkAll()
	h.Sweep()
}

func Run() {
	heap := &Heap{}

	// Create some objects
	a := NewObject("A")
	b := NewObject("B")
	c := NewObject("C")

	// Setup references
	a.refs = []*Object{b}
	b.refs = []*Object{c}
	// c is now unreachable since no other object refers to it

	// Add objects to the heap
	heap.AddObject(a)
	heap.AddObject(b)
	heap.AddObject(c)

	// Before GC, all objects are reachable
	fmt.Println("Before GC:")
	fmt.Println(len(heap.objects)) // Should print 3

	b.refs = []*Object{}

	// Simulate garbage collection
	heap.CollectGarbage()

	// After GC, object c should be collected
	fmt.Println("After GC:")
	fmt.Println(len(heap.objects)) // Should print 2 (since c is unreachable)
}
