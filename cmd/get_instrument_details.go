/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"fmt"
	"github.com/coinbase-samples/intx-cli/utils"
	"github.com/coinbase-samples/intx-sdk-go"
	"github.com/spf13/cobra"
)

var getInstrumentDetailsCmd = &cobra.Command{
	Use:   "get-instrument-details",
	Short: "Get details for instrument.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, _, err := utils.InitClientAndPortfolioId(cmd, false)
		if err != nil {
			return fmt.Errorf("cannot initialize from environment: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.GetInstrumentRequest{
			InstrumentId: utils.GetFlagStringValue(cmd, utils.InstrumentIdFlag),
		}

		response, err := client.GetInstrument(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot get instrument: %w", err)
		}

		return utils.PrintJsonResponse(cmd, response)
	},
}

func init() {
	cmdConfigs := []utils.CommandConfig{
		{
			Command: getInstrumentDetailsCmd,
			FlagConfig: []utils.FlagConfig{
				{
					FlagName:     utils.InstrumentIdFlag,
					Shorthand:    "i",
					Usage:        "Instrument ID (Required)",
					DefaultValue: "",
					Required:     true,
				},
				{
					FlagName:     utils.FormatFlag,
					Shorthand:    "z",
					Usage:        "Pass true for formatted JSON. Default is false",
					DefaultValue: false,
					Required:     false,
				},
			},
		},
	}

	utils.RegisterCommandConfigs(rootCmd, cmdConfigs)
}
