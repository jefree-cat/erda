// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: stream.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*CommentIssueStreamCreateRequest)(nil)
var _ json.Unmarshaler = (*CommentIssueStreamCreateRequest)(nil)
var _ json.Marshaler = (*MRCommentInfo)(nil)
var _ json.Unmarshaler = (*MRCommentInfo)(nil)
var _ json.Marshaler = (*CommentIssueStreamBatchCreateRequest)(nil)
var _ json.Unmarshaler = (*CommentIssueStreamBatchCreateRequest)(nil)
var _ json.Marshaler = (*CommentIssueStreamBatchCreateResponse)(nil)
var _ json.Unmarshaler = (*CommentIssueStreamBatchCreateResponse)(nil)

// CommentIssueStreamCreateRequest implement json.Marshaler.
func (m *CommentIssueStreamCreateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CommentIssueStreamCreateRequest implement json.Marshaler.
func (m *CommentIssueStreamCreateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// MRCommentInfo implement json.Marshaler.
func (m *MRCommentInfo) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// MRCommentInfo implement json.Marshaler.
func (m *MRCommentInfo) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CommentIssueStreamBatchCreateRequest implement json.Marshaler.
func (m *CommentIssueStreamBatchCreateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CommentIssueStreamBatchCreateRequest implement json.Marshaler.
func (m *CommentIssueStreamBatchCreateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CommentIssueStreamBatchCreateResponse implement json.Marshaler.
func (m *CommentIssueStreamBatchCreateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CommentIssueStreamBatchCreateResponse implement json.Marshaler.
func (m *CommentIssueStreamBatchCreateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
