package golang

import (
	"github.com/devstream-io/devstream/internal/pkg/configmanager"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/github"
	"github.com/devstream-io/devstream/internal/pkg/statemanager"
)

func Read(options configmanager.RawOptions) (statemanager.ResourceStatus, error) {
	operator := &plugininstaller.Operator{
		PreExecuteOperations: plugininstaller.PreExecuteOperations{
			validate,
			github.BuildWorkFlowsWrapper(workflows),
		},
		GetStatusOperation: github.GetActionStatus,
	}

	status, err := operator.Execute(options)
	if err != nil {
		return nil, err
	}
	return status, nil
}
