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

package workflow_test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/goombaio/workflow"
)

func TestWorkflowInstance(t *testing.T) {
	expected_name := "workflow_1"
	expected_description := "workflow_1 description"

	workflow1 := workflow.NewWorkflow(expected_name, expected_description)

	if workflow1.ID == uuid.Nil {
		t.Fatalf("Workflow ID expected to be not nil.\n")
	}

	if workflow1.Name != expected_name {
		t.Fatalf("Workflow name expected to be %q but got %s.\n", expected_name, workflow1.Name)
	}

	if workflow1.Description != expected_description {
		t.Fatalf("Workflow name expected to be %q but got %q.\n", expected_description, workflow1.Description)
	}
}
