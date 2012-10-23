// Copyright 2012 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"github.com/globocom/tsuru/cmd"
	. "launchpad.net/gocheck"
	"net/http"
)

func (s *S) TestEnvGetInfo(c *C) {
	e := EnvGet{}
	i := e.Info()
	c.Assert(i.Name, Equals, "env-get")
	c.Assert(i.Usage, Equals, "env-get <appname> [ENVIRONMENT_VARIABLE1] [ENVIRONMENT_VARIABLE2] ...")
	c.Assert(i.Desc, Equals, "retrieve environment variables for an app.")
	c.Assert(i.MinArgs, Equals, 1)
}

func (s *S) TestEnvGetRun(c *C) {
	var stdout, stderr bytes.Buffer
	result := "DATABASE_HOST=somehost\n"
	context := cmd.Context{
		Args:   []string{"someapp", "DATABASE_HOST"},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	client := cmd.NewClient(&http.Client{Transport: &transport{msg: result, status: http.StatusOK}})
	err := (&EnvGet{}).Run(&context, client)
	c.Assert(err, IsNil)
	c.Assert(stdout.String(), Equals, result)
}

func (s *S) TestEnvGetRunWithMultipleParams(c *C) {
	var stdout, stderr bytes.Buffer
	result := "DATABASE_HOST=somehost\nDATABASE_USER=someuser"
	params := []string{"someapp", "DATABASE_HOST", "DATABASE_USER"}
	context := cmd.Context{
		Args:   params,
		Stdout: &stdout,
		Stderr: &stderr,
	}
	client := cmd.NewClient(&http.Client{Transport: &transport{msg: result, status: http.StatusOK}})
	err := (&EnvGet{}).Run(&context, client)
	c.Assert(err, IsNil)
	c.Assert(stdout.String(), Equals, result)
}

func (s *S) TestEnvSetInfo(c *C) {
	e := EnvSet{}
	i := e.Info()
	c.Assert(i.Name, Equals, "env-set")
	c.Assert(i.Usage, Equals, "env-set <appname> <NAME=value> [NAME=value] ...")
	c.Assert(i.Desc, Equals, "set environment variables for an app.")
	c.Assert(i.MinArgs, Equals, 2)
}

func (s *S) TestEnvSetRun(c *C) {
	var stdout, stderr bytes.Buffer
	result := "variable(s) successfully exported\n"
	context := cmd.Context{
		Args:   []string{"someapp", "DATABASE_HOST=somehost"},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	client := cmd.NewClient(&http.Client{Transport: &transport{msg: result, status: http.StatusOK}})
	err := (&EnvSet{}).Run(&context, client)
	c.Assert(err, IsNil)
	c.Assert(stdout.String(), Equals, result)
}

func (s *S) TestEnvSetRunWithMultipleParams(c *C) {
	var stdout, stderr bytes.Buffer
	result := "variable(s) successfully exported\n"
	params := []string{"someapp", "DATABASE_HOST=somehost", "DATABASE_USER=user"}
	context := cmd.Context{
		Args:   params,
		Stdout: &stdout,
		Stderr: &stderr,
	}
	client := cmd.NewClient(&http.Client{Transport: &transport{msg: result, status: http.StatusOK}})
	err := (&EnvSet{}).Run(&context, client)
	c.Assert(err, IsNil)
	c.Assert(stdout.String(), Equals, result)
}

func (s *S) TestEnvUnsetInfo(c *C) {
	e := EnvUnset{}
	i := e.Info()
	c.Assert(i.Name, Equals, "env-unset")
	c.Assert(i.Usage, Equals, "env-unset <appname> <ENVIRONMENT_VARIABLE1> [ENVIRONMENT_VARIABLE2]")
	c.Assert(i.Desc, Equals, "unset environment variables for an app.")
	c.Assert(i.MinArgs, Equals, 2)
}

func (s *S) TestEnvUnsetRun(c *C) {
	var stdout, stderr bytes.Buffer
	result := "variable(s) successfully unset\n"
	context := cmd.Context{
		Args:   []string{"someapp", "DATABASE_HOST"},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	client := cmd.NewClient(&http.Client{Transport: &transport{msg: result, status: http.StatusOK}})
	err := (&EnvUnset{}).Run(&context, client)
	c.Assert(err, IsNil)
	c.Assert(stdout.String(), Equals, result)
}

func (s *S) TestRequestEnvUrl(c *C) {
	result := "DATABASE_HOST=somehost"
	client := cmd.NewClient(&http.Client{Transport: &transport{msg: result, status: http.StatusOK}})
	args := []string{"someapp", "DATABASE_HOST"}
	b, err := requestEnvUrl("GET", args, client)
	c.Assert(err, IsNil)
	c.Assert(b, Equals, result)
}