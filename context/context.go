package context

import "context"

func PathsDownGrade(ctx context.Context, quick, slow func()) {
	quick()
	if ctx.Value("grade") == "slow" {
		return
	}
	slow()
}
