// Code generated by running "go generate" in github.com/neovim/go-client/nvim. DO NOT EDIT.

package nvim

// EmbedOptions specifies options for starting an embedded instance of Nvim.
//
// Deprecated: Use ChildProcessOption instead.
type EmbedOptions struct {
	// Logf log function for rpc.WithLogf.
	Logf func(string, ...interface{})

	// Dir specifies the working directory of the command. The working
	// directory in the current process is used if Dir is "".
	Dir string

	// Path is the path of the command to run. If Path = "", then
	// StartEmbeddedNvim searches for "nvim" on $PATH.
	Path string

	// Args specifies the command line arguments. Do not include the program
	// name (the first argument) or the --embed option.
	Args []string

	// Env specifies the environment of the Nvim process. The current process
	// environment is used if Env is nil.
	Env []string
}

// NewEmbedded starts an embedded instance of Nvim using the specified options.
//
// The application must call Serve() to handle RPC requests and responses.
//
// Deprecated: Use NewChildProcess instead.
func NewEmbedded(options *EmbedOptions) (*Nvim, error) {
	if options == nil {
		options = &EmbedOptions{}
	}
	path := options.Path
	if path == "" {
		path = "nvim"
	}

	return NewChildProcess(
		ChildProcessArgs(append([]string{"--embed"}, options.Args...)...),
		ChildProcessCommand(path),
		ChildProcessEnv(options.Env),
		ChildProcessDir(options.Dir),
		ChildProcessServe(false))
}

// ExecuteLua executes a Lua block.
//
// Deprecated: Use ExecLua instead.
func (v *Nvim) ExecuteLua(code string, result interface{}, args ...interface{}) error {
	if args == nil {
		args = []interface{}{}
	}
	return v.call("nvim_execute_lua", result, code, args)
}

// ExecuteLua executes a Lua block.
//
// Deprecated: Use ExecLua instead.
func (b *Batch) ExecuteLua(code string, result interface{}, args ...interface{}) {
	if args == nil {
		args = []interface{}{}
	}
	b.call("nvim_execute_lua", result, code, args)
}
