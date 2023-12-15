package serializer

import "work/model"

type Task struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	View      int    `json:"view"`
	Status    int    `json:"status"`
	CreatedAt int64  `json:"created_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

func BuildTask(it model.Task) Task {
	it.AddView()
	return Task{
		ID:        it.ID,
		Title:     it.Title,
		Content:   it.Content,
		Status:    it.Status,
		View:      it.View(),
		CreatedAt: it.CreatedAt.Unix(),
		StartTime: it.StartTime,
		EndTime:   it.EndTime,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, it := range items {
		task := BuildTask(it)
		tasks = append(tasks, task)
	}
	return tasks
}
