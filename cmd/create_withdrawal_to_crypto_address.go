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
		client, err := utils.GetClientFromEnv()
		if err != nil {
			return fmt.Errorf("failed to initialize client: %w", err)
		}

		portfolioId, err := utils.GetPortfolioId(cmd, client)
		if err != nil {
			return err
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

		jsonResponse, err := utils.FormatResponseAsJson(cmd, response)
		if err != nil {
			return err
		}

		fmt.Println(jsonResponse)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createWithdrawalToCryptoAddressCmd)

	createWithdrawalToCryptoAddressCmd.Flags().StringP(utils.AssetIdFlag, "i", "", "ID of asset to be withdrawn (Required)")
	createWithdrawalToCryptoAddressCmd.Flags().StringP(utils.AmountFlag, "a", "", "Amount of asset to be withdrawn (Required)")
	createWithdrawalToCryptoAddressCmd.Flags().StringP(utils.NetworkArnIdFlag, "n", "", "Network Arn Id for withdrawal (Required)")
	createWithdrawalToCryptoAddressCmd.Flags().StringP(utils.AddressFlag, "d", "", "Address for withdrawal (Required)")
	createWithdrawalToCryptoAddressCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	createWithdrawalToCryptoAddressCmd.MarkFlagRequired(utils.AssetIdFlag)
	createWithdrawalToCryptoAddressCmd.MarkFlagRequired(utils.AmountFlag)
	createWithdrawalToCryptoAddressCmd.MarkFlagRequired(utils.NetworkArnIdFlag)
	createWithdrawalToCryptoAddressCmd.MarkFlagRequired(utils.AddressFlag)

}
