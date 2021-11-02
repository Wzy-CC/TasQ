package redis

import (
	"encoding/json"
	"log"
	"time"

	"TasQ/app/tasks"
)

// support redis cluster mod
type RedisClusterBroker struct{}

func (rc *RedisClusterBroker) Acquire(queueName string) *tasks.TasQ {
	task := tasks.TasQ{}
	vs, err := rc.BRPop(time.Duration(0), genQueueName(queueName)).Result()
	if err != nil {
		log.Panicf("failed to get task from redis: %s", err)
		return nil // never executed
	}
	v := []byte(vs[1])

	if err := json.Unmarshal(v, &task); err != nil {
		log.Panicf("failed to get task from redis: %s", err)
		return nil // never executed
	}

	return &task
}

func (rc *RedisClusterBroker) Enqueue(task *tasks.TasQ) string {
	taskBytes, err := json.Marshal(task)
	if err != nil {
		log.Panicf("failed to enquue task %+v: %s", task, err)
		return "" // never executed here
	}

	rc.Set(genTaskName(task.ID), taskBytes, time.Duration(r.TaskTTL)*time.Second)
	rc.LPush(genQueueName(task.QueueName), taskBytes)
	return task.ID
}

func (rc *RedisClusterBroker) Update(task *tasks.TasQ) {
	task.UpdatedAt = time.Now()
	taskBytes, err := json.Marshal(task)
	if err != nil {
		log.Panicf("failed to enquue task %+v: %s", task, err)
		return // never executed here
	}
	rc.Set(genTaskName(task.ID), taskBytes, time.Duration(r.TaskTTL)*time.Second)
}

func (rc *RedisClusterBroker) Cancel(task *tasks.TasQ) {
	// redis doesn't support ACK
	return true
}
