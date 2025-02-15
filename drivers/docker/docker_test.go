// Copyright 2016 Iron.io
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/comfucios/runner/common"
	"github.com/comfucios/runner/drivers"
	"github.com/vrischmann/envconfig"
)

type taskDockerTest struct {
	id     string
	input  io.Reader
	output io.Writer
}

func (f *taskDockerTest) Command() string { return "" }
func (f *taskDockerTest) EnvVars() map[string]string {
	return map[string]string{}
}
func (f *taskDockerTest) Labels() map[string]string          { return nil }
func (f *taskDockerTest) Id() string                         { return f.id }
func (f *taskDockerTest) Group() string                      { return "" }
func (f *taskDockerTest) Image() string                      { return "iron/hello" }
func (f *taskDockerTest) Timeout() time.Duration             { return 30 * time.Second }
func (f *taskDockerTest) Logger() (stdout, stderr io.Writer) { return f.output, nil }
func (f *taskDockerTest) WriteStat(drivers.Stat)             { /* TODO */ }
func (f *taskDockerTest) Volumes() [][2]string               { return [][2]string{} }
func (f *taskDockerTest) WorkDir() string                    { return "" }
func (f *taskDockerTest) Close()                             {}
func (f *taskDockerTest) Input() io.Reader                   { return f.input }

func TestDockerIsAlive(t *testing.T) {

	env := common.NewEnvironment(func(e *common.Environment) {})
	dkr := NewDocker(env, drivers.Config{})
	ctx := context.Background()

	isAlive := dkr.IsAlive(ctx)
	if !isAlive {
		t.Fatal("Docker is not working")
	}
}

func TestRunnerDocker(t *testing.T) {
	env := common.NewEnvironment(func(e *common.Environment) {})
	dkr := NewDocker(env, drivers.Config{})
	ctx := context.Background()

	task := &taskDockerTest{"test-docker", nil, nil}

	cookie, err := dkr.Prepare(ctx, task)
	if err != nil {
		t.Fatal("Couldn't prepare task test")
	}
	defer cookie.Close()

	result, err := cookie.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if result.Status() != "success" {
		t.Fatal("Test should successfully run the image")
	}
}

func TestRunnerDockerStdin(t *testing.T) {
	env := common.NewEnvironment(func(e *common.Environment) {})
	dkr := NewDocker(env, drivers.Config{})
	ctx := context.Background()

	input := `{"name": "test"}`
	var output bytes.Buffer

	task := &taskDockerTest{"test-docker-stdin", bytes.NewBufferString(input), &output}

	cookie, err := dkr.Prepare(ctx, task)
	if err != nil {
		t.Fatal("Couldn't prepare task test")
	}
	defer cookie.Close()

	result, err := cookie.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if result.Status() != "success" {
		t.Error("Test should successfully run the image")
	}

	expect := "Hello test!"
	got := output.String()
	if !strings.Contains(got, expect) {
		t.Errorf("Test expected output to contain '%s', got '%s'", expect, got)
	}
}

func TestConfigLoadMemory(t *testing.T) {
	if err := os.Setenv("MEMORY_PER_JOB", "128M"); err != nil {
		t.Fatalf("Could not set MEMORY_PER_JOB: %v", err)
	}

	var conf drivers.Config
	if err := envconfig.Init(&conf); err != nil {
		t.Fatalf("Could not read config: %v", err)
	}

	if conf.Memory != 128*1024*1024 {
		t.Fatalf("Memory read from config should match 128M, got %d", conf.Memory)
	}
}
