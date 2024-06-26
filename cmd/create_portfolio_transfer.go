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

var createPortfolioTransferCmd = &cobra.Command{
	Use:   "create-portfolio-transfer",
	Short: "Create a new transfer.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, portfolioId, err := utils.InitClientAndPortfolioId(cmd, true)
		if err != nil {
			return fmt.Errorf("cannot initialize from environment: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.CreatePortfolioTransferRequest{
			From:    portfolioId,
			To:      utils.GetFlagStringValue(cmd, utils.ToFlag),
			AssetId: utils.GetFlagStringValue(cmd, utils.AssetIdFlag),
			Amount:  utils.GetFlagStringValue(cmd, utils.AmountFlag),
		}

		response, err := client.CreatePortfolioTransfer(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot create transfer: %w", err)
		}

		return utils.PrintJsonResponse(cmd, response)
	},
}

func init() {
	cmdConfigs := []utils.CommandConfig{
		{
			Command: createPortfolioTransferCmd,
			FlagConfig: []utils.FlagConfig{
				{
					FlagName:     utils.ToFlag,
					Shorthand:    "t",
					Usage:        "Name of destination portfolio (Required)",
					DefaultValue: "",
					Required:     true,
				},
				{
					FlagName:     utils.AssetIdFlag,
					Shorthand:    "i",
					Usage:        "ID of asset to be transferred (Required)",
					DefaultValue: "",
					Required:     true,
				},
				{
					FlagName:     utils.AmountFlag,
					Shorthand:    "a",
					Usage:        "Amount of asset to be transferred (Required)",
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
