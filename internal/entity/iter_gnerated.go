// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package entity

import "context"

type CommitIter interface {
	ForEach(context.Context, func(*Commit) error) error
	Slice(context.Context) ([]*Commit, error)
}

type RepositoryIter interface {
	ForEach(context.Context, func(*Repository) error) error
	Slice(context.Context) ([]*Repository, error)
}

type JobIter interface {
	ForEach(context.Context, func(*Job) error) error
	Slice(context.Context) ([]*Job, error)
}

type OrganizationIter interface {
	ForEach(context.Context, func(*Organization) error) error
	Slice(context.Context) ([]*Organization, error)
}

type PullRequestIter interface {
	ForEach(context.Context, func(*PullRequest) error) error
	Slice(context.Context) ([]*PullRequest, error)
}

type DeveloperExpIter interface {
	ForEach(context.Context, func(*DeveloperExp) error) error
	Slice(context.Context) ([]*DeveloperExp, error)
}

type FileIter interface {
	ForEach(context.Context, func(*File) error) error
	Slice(context.Context) ([]*File, error)
}

type ChangeMeasuresIter interface {
	ForEach(context.Context, func(*ChangeMeasures) error) error
	Slice(context.Context) ([]*ChangeMeasures, error)
}

type FileMeasuresIter interface {
	ForEach(context.Context, func(*FileMeasures) error) error
	Slice(context.Context) ([]*FileMeasures, error)
}

type FeedbackIter interface {
	ForEach(context.Context, func(*Feedback) error) error
	Slice(context.Context) ([]*Feedback, error)
}