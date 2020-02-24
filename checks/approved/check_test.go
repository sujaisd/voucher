package approved

import (
	"context"
	"testing"

	"github.com/grafeas/voucher"
	"github.com/grafeas/voucher/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApprovedCheck(t *testing.T) {
	ctx := context.Background()
	imageData, err := voucher.NewImageReference("gcr.io/voucher-test-project/apps/staging/voucher-internal@sha256:73d506a23331fce5cb6f49bfb4c27450d2ef4878efce89f03a46b27372a88430")
	require.NoErrorf(t, err, "failed to get ImageData: %s", err)
	buildDetail := repository.BuildDetail{RepositoryURL: "https://github.com/grafeas/voucher-internal", Commit: "efgh6543"}
	commitURL := "https://github.com/grafeas/voucher-internal/commit/efgh6543"

	cases := []struct {
		name                 string
		defaultBranchCommits []repository.CommitRef
		isSigned             bool
		status               string
		pullRequest          repository.PullRequest
		shouldPass           bool
		err                  error
	}{
		{
			name:                 "Should pass",
			defaultBranchCommits: []repository.CommitRef{{URL: commitURL}},
			isSigned:             true,
			status:               "SUCCESS",
			pullRequest:          repository.PullRequest{IsMerged: true, MergeCommit: repository.CommitRef{URL: commitURL}, HasRequiredApprovals: true},
			shouldPass:           true,
			err:                  nil,
		},
		{
			name:                 "Not built off default branch",
			defaultBranchCommits: []repository.CommitRef{{URL: "otherCommit"}},
			isSigned:             true,
			status:               "SUCCESS",
			pullRequest:          repository.PullRequest{IsMerged: true, MergeCommit: repository.CommitRef{URL: commitURL}, HasRequiredApprovals: true},
			shouldPass:           false,
			err:                  ErrNotOnDefaultBranch,
		},
		{
			name:                 "Commit not signed",
			defaultBranchCommits: []repository.CommitRef{{URL: commitURL}},
			isSigned:             false,
			status:               "SUCCESS",
			pullRequest:          repository.PullRequest{IsMerged: true, MergeCommit: repository.CommitRef{URL: commitURL}, HasRequiredApprovals: true},
			shouldPass:           false,
			err:                  ErrNotSigned,
		},
		{
			name:                 "Commit not a merge commit",
			defaultBranchCommits: []repository.CommitRef{{URL: commitURL}},
			isSigned:             true,
			status:               "SUCCESS",
			pullRequest:          repository.PullRequest{IsMerged: true, MergeCommit: repository.CommitRef{URL: "otherURL"}, HasRequiredApprovals: true},
			shouldPass:           false,
			err:                  ErrNotMergeCommit,
		},
		{
			name:                 "Commit PR does not have required approvals",
			defaultBranchCommits: []repository.CommitRef{{URL: commitURL}},
			isSigned:             true,
			status:               "SUCCESS",
			pullRequest:          repository.PullRequest{IsMerged: true, MergeCommit: repository.CommitRef{URL: commitURL}, HasRequiredApprovals: false},
			shouldPass:           false,
			err:                  ErrMissingRequiredApprovals,
		},
		{
			name:                 "CI check not successful",
			defaultBranchCommits: []repository.CommitRef{{URL: commitURL}},
			isSigned:             true,
			status:               "FAILURE",
			pullRequest:          repository.PullRequest{IsMerged: true, MergeCommit: repository.CommitRef{URL: commitURL}, HasRequiredApprovals: true},
			shouldPass:           false,
			err:                  ErrNotPassedCI,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			commit := repository.Commit{
				URL:                    commitURL,
				Status:                 testCase.status,
				IsSigned:               testCase.isSigned,
				AssociatedPullRequests: []repository.PullRequest{testCase.pullRequest},
			}
			defaultBranch := repository.Branch{Name: "production", CommitRefs: testCase.defaultBranchCommits}

			metadataClient := new(voucher.MockMetadataClient)
			metadataClient.On("GetBuildDetail", ctx, imageData).Return(buildDetail, nil)

			repositoryClient := new(repository.MockClient)
			repositoryClient.On("GetCommit", ctx, buildDetail).Return(commit, nil)
			repositoryClient.On("GetDefaultBranch", ctx, buildDetail).Return(defaultBranch, nil)

			orgCheck := new(check)
			orgCheck.SetMetadataClient(metadataClient)
			orgCheck.SetRepositoryClient(repositoryClient)

			status, err := orgCheck.Check(ctx, imageData)

			assert.Equal(t, testCase.shouldPass, status)
			if testCase.err != nil {
				assert.EqualError(t, testCase.err, err.Error())
			}
		})
	}
}
