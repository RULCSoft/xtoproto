// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package service implements the XToProtoService gRPC service.
package service

import (
	"context"

	sgrpcpb "github.com/google/xtoproto/proto/service"
)

// FileReaderFunc is a function that returns the content of a file or an error.
type FileReaderFunc func(ctx context.Context, path string) ([]byte, error)

// FileWriterFunc is a function that outputs a source code file to the given path.
type FileWriterFunc func(ctx context.Context, path string, data []byte) error

type service struct {
	defaultWorkspaceDir string
	readFile            FileReaderFunc
	writeFile           FileWriterFunc
}

// New returns a new XToProtoService.
//
// defaultWorkspaceDir will be used for the workspace directory when no value
// is passed in with the request.
func New(defaultWorkspaceDir string, readFile FileReaderFunc, writeFile FileWriterFunc) sgrpcpb.XToProtoServiceServer {
	return &service{
		defaultWorkspaceDir,
		readFile,
		writeFile,
	}
}
