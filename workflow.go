// Copyright 2018, Goomba project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package workflow

import (
	"github.com/google/uuid"

	"github.com/goombaio/dag"
)

// Workflow ...
type Workflow struct {
	ID          uuid.UUID
	Name        string
	Description string
	dag         *dag.DAG
}

// NewWorkflow ...
func NewWorkflow(name string, description string) *Workflow {
	w := &Workflow{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		dag:         dag.NewDAG(),
	}

	return w
}

// AddTask ...
func (w *Workflow) AddTask(task *Task) error {
	taskVertex := dag.NewVertex(task)

	err := w.dag.AddVertex(taskVertex)
	return err
}
