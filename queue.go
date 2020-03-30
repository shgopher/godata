package godata

type Queue struct {
	val  []interface{}
	leng int
}
// return cap of the slice.
func NewQueue(count int) *Queue{
	return &Queue{
		val:  make([]interface{},0,count),
		leng: 0,
	}
}
func (q *Queue) Length() int {
	return q.leng
}
// return length of the queue
func (q *Queue) Push(value interface{}) int {
	q.val = append(q.val, value)
	q.leng++
	return q.Length()
}
// return the top of queue,and delete a node of the queue.
func (q *Queue) Pop() interface{} {
	var value interface{}
	if q.Length() > 1 {
		value = q.val[0]
		q.val = q.val[1:]
		q.leng--
	} else if q.Length() == 1 {
		value = q.val[0]
		q.val = nil
		q.leng--
	} else {
		return nil
	}

	return value
}
func (q *Queue) Top() interface{} {
	if q.Length() > 0 {
		return q.val[0]
	}
	return nil
}
