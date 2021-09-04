package timewheel

import (
	"errors"
	"time"
)
type BasicJob func(interface{})
/**
timewheel
一个指针指向当前位置
一个固定大小的数组，数组元素是一个集合
集合的元素应该包含周期数，需要转几圈，每次指针指向该元素的时候圈数减少1，如果为0就执行任务
 */
type Job interface{
	OnExecuting(interface{})
}

func (b BasicJob) OnExecuting(v interface{}) {
	b(v)
}

type TimeWheel struct {
	bucketList []*Bucket
	cursor  int
	stop    bool
	period  time.Duration
	call    Job
}
func NewTimeWheel(period time.Duration,N int, call Job) *TimeWheel{
	bucketList := make([]*Bucket,N)
	for i,_ := range bucketList {
		bucketList[i] = NewBucket()
	}
	return &TimeWheel{
		bucketList: bucketList,
		cursor :0,
		period: period,
		stop : false,
		call : call,
	}
}

func (t *TimeWheel) Start() {
	ticker := time.NewTicker(t.period)
	for range ticker.C {
		if t.stop == true {
			break
		}

		t.cursor = (t.cursor + 1)%t.Len()
		bucket := t.bucketList[t.cursor]
		//log.Printf("cursor %v: %v", t.cursor, bucket)
		for _, task := range bucket.tasks {
			if task.circle == 0 {
				t.call.OnExecuting(task.callbackParam)
				delete(bucket.tasks, task.taskid)
				continue
			}
			task.circle = task.circle - 1
		}
	}
}

//runnning Time means what time you want this task to run
func (t *TimeWheel) AddTask(taskid string , runningTime time.Time, callbackParam interface{}) {
	totalTime := int64(time.Second) * (runningTime.Unix()-time.Now().Unix())
	cursor := totalTime/(int64(t.period)*int64(t.Len()))
	var task = NewTask(taskid, int(cursor), callbackParam)
	taskpos := totalTime %(int64(t.period)*int64(t.Len()))/int64(t.period) + int64(t.cursor)
	if totalTime %(int64(t.period)*int64(t.Len())) % int64(t.period) != 0 {
		taskpos += 1
	}
	taskpos  = taskpos % int64(t.Len())
	t.bucketList[taskpos].AddTask(task)
}


func (t *TimeWheel) Stop(){
	t.stop = true
}

func (t *TimeWheel) Len () int {
	return len(t.bucketList)
}

type Bucket struct {
	tasks map[interface{}]*Task
}

func NewBucket() *Bucket{
	return &Bucket{
		tasks : make(map[interface{}]*Task),
	}
}

func (b *Bucket)AddTask(task *Task) error{
     if _,ok := b.tasks[task.taskid]; ok {
     	return errors.New("task already add")
	 }
	 b.tasks[task.taskid] = task
	 return nil
}

type Task struct{
	taskid string //任务ID
	circle int //周期数
	callbackParam interface{}
}

func NewTask(taskid string, circle int, callbackParam interface{}) *Task{
	return &Task{
		circle: circle,
		taskid: taskid,
		callbackParam: callbackParam,
	}
}
