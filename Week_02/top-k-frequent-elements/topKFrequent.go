package top_k_frequent_elements

import "container/heap"

type kv struct{
	key int
	val int
}

type MHeap []kv
func (h MHeap)Len()int{return len(h)}
func (h MHeap)Less(i,j int)bool{return h[i].val>h[j].val}
func (h MHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MHeap)Push(x interface{}){
	*h = append(*h,x.(kv))
}
func (h*MHeap)Pop() interface{}{
	top:= (*h)[len(*h)-1]
	*h=(*h)[:len(*h)-1]
	return top
}

func topKFrequent(nums []int, k int) []int {
	res:=[]int{}
	m:=make(map[int]int)
	for _,v:=range nums{
		m[v]=  m[v]+1
	}
	h:=&MHeap{}
	heap.Init(h)
	for k,v:=range m{
		heap.Push(h,kv{k,v})
	}
	for i:=0;i<k;i++{
		o:=heap.Pop(h).(kv)
		res=append(res,o.key)
	}
	return res
}
