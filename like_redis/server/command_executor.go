package server

type CommandExecutor struct {
	storage *Storage
}

func NewCommandExecutor(storage *Storage) *CommandExecutor {
	return &CommandExecutor{
		storage: storage,
	}
}

func (e *CommandExecutor) Execute(cmd interface{}) (interface{}, error) {
	if cmd, ok := cmd.(SetCommandDefinition); ok {
		return e.storage.SetCounter(cmd.Key, cmd.Value)
	}
	if cmd, ok := cmd.(GetCommandDefinition); ok {
		return e.storage.GetCounter(cmd.Key)
	}
	if cmd, ok := cmd.(IncCommandDefinition); ok {
		return e.storage.IncrementCounter(cmd.Key, cmd.Value)
	}
	if cmd, ok := cmd.(DecCommandDefinition); ok {
		return e.storage.DecrementCounter(cmd.Key, cmd.Value)
	}

	return "unknown", nil
}
