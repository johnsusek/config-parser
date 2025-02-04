// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)


func TestLogNormal0(t *testing.T) {
	parser := &parsers.Log{}
	line := strings.TrimSpace("log global")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestLogNormal1(t *testing.T) {
	parser := &parsers.Log{}
	line := strings.TrimSpace("log stdout format short daemon # send log to systemd")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestLogNormal2(t *testing.T) {
	parser := &parsers.Log{}
	line := strings.TrimSpace("log stdout format raw daemon # send everything to stdout")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestLogNormal3(t *testing.T) {
	parser := &parsers.Log{}
	line := strings.TrimSpace("log stderr format raw daemon notice # send important events to stderr")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestLogNormal4(t *testing.T) {
	parser := &parsers.Log{}
	line := strings.TrimSpace("log 127.0.0.1:514 local0 notice # only send important events")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestLogNormal5(t *testing.T) {
	parser := &parsers.Log{}
	line := strings.TrimSpace("log 127.0.0.1:514 local0 notice notice # same but limit output level")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestLogNormal6(t *testing.T) {
	parser := &parsers.Log{}
	line := strings.TrimSpace("log 127.0.0.1:1515 len 8192 format rfc5424 local2 info")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestLogFail0(t *testing.T) {
	parser := &parsers.Log{}
	line := strings.TrimSpace("log")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result()
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
func TestLogFail1(t *testing.T) {
	parser := &parsers.Log{}
	line := strings.TrimSpace("---")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result()
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
func TestLogFail2(t *testing.T) {
	parser := &parsers.Log{}
	line := strings.TrimSpace("--- ---")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result()
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
