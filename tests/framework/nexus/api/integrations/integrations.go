// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: BUSL-1.1

package integrations

//go:generate protoc --proto_path=../:. --go_out=. --go_opt=paths=source_relative --rangerrpc_out=. integrations.proto
