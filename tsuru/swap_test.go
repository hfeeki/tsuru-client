// Copyright 2015 tsuru-client authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"net/http"

	"github.com/tsuru/tsuru/cmd"
	"github.com/tsuru/tsuru/cmd/cmdtest"
	"gopkg.in/check.v1"
)

func (s *S) TestSwapInfo(c *check.C) {
	command := appSwap{}
	c.Assert(command.Info(), check.NotNil)
}

func (s *S) TestSwap(c *check.C) {
	var buf bytes.Buffer
	var called bool
	transport := cmdtest.ConditionalTransport{
		Transport: cmdtest.Transport{Status: http.StatusOK, Message: ""},
		CondFunc: func(r *http.Request) bool {
			called = true
			return r.Method == "PUT" && r.URL.Path == "/swap"
		},
	}
	context := cmd.Context{
		Args:   []string{"app1", "app2"},
		Stdout: &buf,
	}
	client := cmd.NewClient(&http.Client{Transport: &transport}, nil, manager)
	command := appSwap{}
	err := command.Run(&context, client)
	c.Assert(err, check.IsNil)
	c.Assert(called, check.Equals, true)
	expected := "Apps successfully swapped!\n"
	c.Assert(buf.String(), check.Equals, expected)
}

func (s *S) TestSwapWhenAppsAreNotEqual(c *check.C) {
	var buf bytes.Buffer
	var called int
	stdin := bytes.NewBufferString("yes")
	transportError := cmdtest.ConditionalTransport{
		Transport: cmdtest.Transport{Status: http.StatusPreconditionFailed, Message: "Apps are not equal."},
		CondFunc: func(r *http.Request) bool {
			called += 1
			return r.URL.RawQuery == "app1=app1&app2=app2&force=false"
		},
	}
	transportOk := cmdtest.ConditionalTransport{
		Transport: cmdtest.Transport{Status: http.StatusOK, Message: ""},
		CondFunc: func(r *http.Request) bool {
			called += 1
			return r.URL.RawQuery == "app1=app1&app2=app2&force=true"
		},
	}
	multiTransport := cmdtest.MultiConditionalTransport{
		ConditionalTransports: []cmdtest.ConditionalTransport{transportError, transportOk},
	}
	context := cmd.Context{
		Args:   []string{"app1", "app2"},
		Stdout: &buf,
		Stdin:  stdin,
	}
	client := cmd.NewClient(&http.Client{Transport: &multiTransport}, nil, manager)
	command := appSwap{}
	err := command.Run(&context, client)
	c.Assert(err, check.IsNil)
	c.Assert(called, check.Equals, 2)
}

func (s *S) TestSwapIsACommand(c *check.C) {
	var _ cmd.Command = &appSwap{}
}
