package flags

import (
	"github.com/urfave/cli"
)

var Verbose = cli.BoolFlag{
	Name:  "verbose,v",
	Usage: `Print verbose log for operation.`,
}

var LogDirectory = cli.StringFlag{
	Name:  "log-dir",
	Usage: `Path to log directory where stdout/stderr log files will be placed (default: "${cwd}/logs/${time}")`,
}

var Concurrency = cli.Int64Flag{
	Name:  "concurrency,C",
	Usage: `Throttle concurrency number for running jobs`,
	Value: 8,
}

// DryRun ...
// var DryRun = cli.BoolFlag{
// 	Name:  "dry-run",
// 	Usage: `Print the pipeline(s) that would be run and then exit. (default: false)`,
// }

var Provider = cli.StringFlag{
	Name:  "provider,p",
	Usage: `Job service provider, either of [aws, gcp, vbox, hyperv]`,
	Value: "aws",
}

// Tasks ...
var Tasks = cli.StringFlag{
	Name:  "tasks",
	Usage: `Path to CSV of task parameters, expected to specify --env, --input, --input-recursive and --output-recursive. (required)`,
}

// Image ...
var Image = cli.StringFlag{
	Name:  "image",
	Usage: `Image name from Docker Hub or other Docker image service.`,
	Value: "ubuntu:14.04",
}

// Script ...
var Script = cli.StringFlag{
	Name:  "script",
	Usage: `Local path to a script to run inside the job's Docker container. (required)`,
}

// Shared ...
var Shared = cli.StringSliceFlag{
	Name:  "shared,S",
	Usage: `Shared data URL on cloud storage bucket. (e.g. s3://~)`,
}

// Keep ...
var Keep = cli.BoolFlag{
	Name:  "keep",
	Usage: `Keep instances created for computing event after everything gets done`,
}

// MinCoresFlag ...
// var MinCoresFlag = cli.UintFlag{
// 	Name:  "min-cores",
// 	Usage: `Minimum CPU cores for each job.`,
// 	Value: 1,
// }

// MinRAMFlag ...
// var MinRAMFlag = cli.Float64Flag{
// 	Name:  "min-ram",
// 	Usage: `Minimum RAM per job in GB.`,
// 	Value: 4,
// }

// Disksize ...
var Disksize = cli.UintFlag{
	Name:  "disk-size",
	Usage: `Size of data disk to attach for each job in GB.`,
	Value: 64,
}

// Env flag allows to provide environment variables to each container, instead of using "--env" in tasks file.
// This flag is expected to be used for specifying common environment variables to containers,
// but at the same time, which is specific for project or every command line.
var Env = cli.StringSliceFlag{
	Name:  "env,E",
	Usage: `Environment variables to pass to all the workflow containers`,
}

// SharedDataDisksize ...
var SharedDataDisksize = cli.IntFlag{
	Name:  "shareddata-disksize",
	Usage: `Disk size of shared data instance (in GB)`,
	Value: 64,
}
