package testdata

const (
	ExampleCLIImport = `import (
	"fmt"
	cli "grpc/cli/testapi"
	"os"

	"goa.design/goa"
	"google.golang.org/grpc"
)
`

	ExampleSingleHostCLIImport = `import (
	"fmt"
	cli "grpc/cli/single_host"
	"os"

	"goa.design/goa"
	"google.golang.org/grpc"
)
`
	ExamplePkgPathCLIImport = `import (
	"fmt"
	cli "my/pkg/path/grpc/cli/testapi"
	"os"

	"goa.design/goa"
	"google.golang.org/grpc"
)
`
	ExampleSingleHostPkgPathCLIImport = `import (
	"fmt"
	cli "my/pkg/path/grpc/cli/single_host"
	"os"

	"goa.design/goa"
	"google.golang.org/grpc"
)
`

ExampleCLICode = `func doGRPC(scheme, host string, timeout int, debug bool) (goa.Endpoint, interface{}, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("could not connect to gRPC server at %s: %v", host, err))
	}
	return cli.ParseEndpoint(conn)
}

func grpcUsageCommands() string {
	return cli.UsageCommands()
}

func grpcUsageExamples() string {
	return cli.UsageExamples()
}
`
)
