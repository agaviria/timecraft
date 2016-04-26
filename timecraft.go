// Copyright 2016 Alejandro R. Gaviria
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/agaviria/timecraft/cmd"
	"github.com/codegangsta/cli"
)

// Version is the application semantic version
const Version = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "TimeCraft"
	app.Usage = "Time tracking tasks"
	app.Version = Version
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		cmd.Setup,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
