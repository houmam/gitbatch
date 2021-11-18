package app

import (
	"os"
	"testing"

	"github.com/isacikgoz/gitbatch/internal/git"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func TestQuick(t *testing.T) {
	th := git.InitTestRepositoryFromLocal(t)
	defer th.CleanUp(t)

	var tests = []struct {
		inp1 []string
		inp2 string
		log  zerolog.Logger
	}{
		{
			[]string{th.DirtyRepoPath()},
			"fetch",
			log.Output(zerolog.ConsoleWriter{Out: os.Stderr}),
		}, {
			[]string{th.DirtyRepoPath()},
			"pull",
			log.Output(zerolog.ConsoleWriter{Out: os.Stderr}),
		},
	}
	for _, test := range tests {
		err := quick(test.inp1, test.inp2, test.log)
		require.NoError(t, err)
	}
}
