package shell

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
)

type fish struct {
	path string
}

func (f *fish) GetPath() string {
	return f.path
}

func (f *fish) SetupWrapper(clientPath string) string {
	return fmt.Sprintf("%[1]s --mode=wrapper", clientPath)
}

var fishHooksTmpl = `
function __shell_logger_preexec -e fish_preexec
  env \
    {{.StartTimeEnv}}=(date -u +"%Y-%m-%dT%H:%M:%SZ") \
    {{.ReturnCodeEnv}}=$status \
    {{.CommandEnv}}=$history[1] \
    shell_logger --mode=submit
end
`

func (f *fish) SetupHooks(clientPath string) string {
	tmpl, err := template.New("fish-hook").Parse(fishHooksTmpl)
	if err != nil {
		panic(err)
	}
	return renderHooks(tmpl, clientPath)
}

func (f *fish) InWrapper() bool {
	return f.GetSocketPath() != ""
}

func (f *fish) GetSocketPath() string {
	return os.Getenv(socketEnv)
}

func (f *fish) SetSocketPath(socketPath string) {
	os.Setenv(socketEnv, socketPath)
}

func (f *fish) GetStartTime() (int, error) {
	return strconv.Atoi(os.Getenv(startTimeEnv))
}

func (f *fish) GetCommand() string {
	return os.Getenv(commandEnv)
}

func (f *fish) GetReturnCode() (int, error) {
	return strconv.Atoi(os.Getenv(returnCodeEnv))
}

func (f *fish) GetEndTime() (int, error) {
	return strconv.Atoi(os.Getenv(endTimeEnv))
}
