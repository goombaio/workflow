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

// Workflow type represents a workflow engine.
type Workflow struct {
	// Unique ID for this workflow engine.
	// Used for traceability, metrics, monitoring, etc ...
	ID uuid.UUID

	// Name of this workflow engine.
	Name string

	// Description of this worflow and its purpose.
	Description string

	// A Directed acyclic graph or DAG  describes the workflow processes and
	// how are they related each other.
	graph *dag.DAG
}

// NewWorkflow creates a new workflow engine given a name and description.
//
// It generates an unique internal ID which will be used to identify this
// concrete workflow engine among others in order to support traceability,
// metrics, monitoring, etc ...
//
// The workflow is defined by a graph, a directed acyclic graph or DAG. This
// graph will describe the inner processes belonging to it.
func NewWorkflow(name string, description string) *Workflow {
	w := &Workflow{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		graph:       dag.NewDAG(),
	}

	return w
}
