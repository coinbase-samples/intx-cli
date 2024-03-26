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

var createWithdrawalToCryptoAddressCmd = &cobra.Command{
	Use:   "create-withdrawal-to-crypto-address",
	Short: "Create a withdrawal to crypto address.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, portfolioId, err := utils.InitClientAndPortfolioId(cmd, true)
		if err != nil {
			return fmt.Errorf("cannot initialize from environment: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		addNetworkFeeToTotalPtr := utils.GetFlagBoolValue(cmd, utils.AddNetworkFeeToTotalFlag)
		addNetworkFeeToTotal := false
		if addNetworkFeeToTotalPtr != nil {
			addNetworkFeeToTotal = *addNetworkFeeToTotalPtr
		}

		request := &intx.CreateWithdrawalToCryptoAddressRequest{
			PortfolioId:          portfolioId,
			AssetId:              utils.GetFlagStringValue(cmd, utils.AssetIdFlag),
			Amount:               utils.GetFlagStringValue(cmd, utils.AmountFlag),
			AddNetworkFeeToTotal: addNetworkFeeToTotal,
			NetworkArnId:         utils.GetFlagStringValue(cmd, utils.NetworkArnIdFlag),
			Address:              utils.GetFlagStringValue(cmd, utils.AddressFlag),
			Nonce:                utils.GetFlagStringValue(cmd, utils.NonceFlag),
		}

		response, err := client.CreateWithdrawalToCryptoAddress(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot create withdrawal: %w", err)
		}

		return utils.PrintJsonResponse(cmd, response)
	},
}

func init() {
	cmdConfigs := []utils.CommandConfig{
		{
			Command: createWithdrawalToCryptoAddressCmd,
			FlagConfig: []utils.FlagConfig{
				{
					FlagName:     utils.AssetIdFlag,
					Shorthand:    "i",
					Usage:        "ID of asset to be withdrawn (Required)",
					DefaultValue: "",
					Required:     true,
				},
				{
					FlagName:     utils.AmountFlag,
					Shorthand:    "a",
					Usage:        "Amount of asset to be withdrawn (Required)",
					DefaultValue: "",
					Required:     true,
				},
				{
					FlagName:     utils.NetworkArnIdFlag,
					Shorthand:    "n",
					Usage:        "Network Arn Id for withdrawal (Required)",
					DefaultValue: "",
					Required:     true,
				},
				{
					FlagName:     utils.AddressFlag,
					Shorthand:    "d",
					Usage:        "Address for withdrawal (Required)",
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
