/*
 * Copyright (c) 2024 OceanBase.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package agent

import (
	"github.com/spf13/cobra"

	"github.com/oceanbase/obshell/client/utils/printer"
	"github.com/oceanbase/obshell/agent/config"
	"github.com/oceanbase/obshell/agent/constant"
	"github.com/oceanbase/obshell/agent/global"
	ocsagentlog "github.com/oceanbase/obshell/agent/log"
)

func initStartCmd(c *cobra.Command) {
	c.Hidden = false
	c.Short = "Start the background running obshell agent process."
	c.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		ocsagentlog.InitLogger(config.DefaultClientLoggerConifg())
		global.InitGlobalVariable()
	}

	c.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		printer.PrintHelpFunc(cmd, []string{})
	})
	c.Example = startCmdExample()
	hideFlags(c, constant.FLAG_PID)
}

func hideFlags(c *cobra.Command, names ...string) {
	for _, n := range names {
		c.Flags().MarkHidden(n)
	}
}

func startCmdExample() string {
	return `  obshell agent start --ip 192.168.1.1
  obshell agent start --ip 192.168.1.1 -P 2886 `
}
