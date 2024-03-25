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

var createCryptoAddressCmd = &cobra.Command{
	Use:   "create-crypto-address",
	Short: "Create crypto address.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, portfolioId, err := utils.InitClientAndPortfolioId(cmd, true)
		if err != nil {
			return fmt.Errorf("cannot initialize from environment: %w", err)
		}

		assetId, err := cmd.Flags().GetString(utils.AssetIdFlag)
		if err != nil {
			return fmt.Errorf("cannot cancel order: %w", err)
		}

		networkArnId, err := cmd.Flags().GetString(utils.NetworkArnIdFlag)
		if err != nil {
			return fmt.Errorf("cannot cancel order: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.CreateCryptoAddressRequest{
			PortfolioId:  portfolioId,
			AssetId:      assetId,
			NetworkArnId: networkArnId,
		}

		response, err := client.CreateCryptoAddress(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot create address: %w", err)
		}

		return utils.PrintJsonResponse(cmd, response)
	},
}

func init() {
	cmdConfigs := []utils.CommandConfig{
		{
			Command: createCryptoAddressCmd,
			FlagConfig: []utils.FlagConfig{
				{
					FlagName:     utils.AssetIdFlag,
					Shorthand:    "a",
					Usage:        "ID of the asset (Required)",
					DefaultValue: "",
					Required:     true,
				},
				{
					FlagName:     utils.NetworkArnIdFlag,
					Shorthand:    "n",
					Usage:        "Network Arn ID of the asset (Required)",
					DefaultValue: "",
					Required:     true,
				},
				{
					FlagName:     utils.PortfolioIdFlag,
					Shorthand:    "i",
					Usage:        "Portfolio ID. Uses environment variable if blank",
					DefaultValue: "",
					Required:     false,
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
