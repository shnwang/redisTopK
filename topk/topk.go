package topk

import "fmt"
import "sort"

type Data interface {
	Val() int
	Str() string
}

type TopK struct {
	data []Data
	k    int
}

func (t *TopK) Len() int {
	return len(t.data)
}

func (t *TopK) Swap(i, j int) {
	t.data[i], t.data[j] = t.data[j], t.data[i]
}

func (t *TopK) Less(i, j int) bool {
	return t.data[i].Val() < t.data[j].Val()
}

func (t *TopK) Insert(d Data) {
	if t.Len() < t.k {
		t.data = append(t.data, d)
		t.Sort()
	} else {
		min := t.data[0]
		if min.Val() < d.Val() {
			t.data = append(t.data, d)
			t.data = t.data[1:]
			t.Sort()
		}
	}
}

func (t *TopK) GetData() ([]Data){
    return t.data;
}

func (t *TopK) Min() Data {
    if len(t.data) == 0 {
        return nil;
    } else {
        return t.data[0];
    }
}

func (t *TopK) Max() Data {
    if len(t.data) == 0 {
        return nil;
    } else {
        return t.data[len(t.data)-1];
    }
}

func (t *TopK) Print() {
	for _, d := range t.data {
		fmt.Printf("%v\n", d.Str())
	}
	fmt.Printf("\n")
}

func (t *TopK) Sort() {
	sort.Sort(t)
}

func TopKNew(k int) *TopK {
	t := TopK{}
	t.k = k
	return &t
}
