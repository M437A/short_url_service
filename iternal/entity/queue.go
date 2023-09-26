package entity

type Queue struct {
	items []string
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (q *Queue) Add(item string) {
	q.items = append(q.items, item)
}

func (q *Queue) Get() string {
	if len(q.items) == 0 {
		return ""
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
