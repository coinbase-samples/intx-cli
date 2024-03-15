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

var createWithdrawalToCounterPartyIdCmd = &cobra.Command{
	Use:   "create-withdrawal-to-counterparty-id",
	Short: "Create a withdrawal to counterparty id.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, portfolioId, err := utils.InitClientAndPortfolioId(cmd, true)
		if err != nil {
			return fmt.Errorf("cannot initialize from environment: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.CreateWithdrawalToCounterpartyIdRequest{
			PortfolioId:    portfolioId,
			CounterpartyId: utils.GetFlagStringValue(cmd, utils.CounterpartyIdFlag),
			AssetId:        utils.GetFlagStringValue(cmd, utils.AssetIdFlag),
			Amount:         utils.GetFlagStringValue(cmd, utils.AmountFlag),
			Nonce:          utils.GetFlagStringValue(cmd, utils.NonceFlag),
		}

		response, err := client.CreateWithdrawalToCounterpartyId(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot create withdrawal: %w", err)
		}

		return utils.PrintJsonResponse(cmd, response)
	},
}

func init() {
	rootCmd.AddCommand(createWithdrawalToCounterPartyIdCmd)

	createWithdrawalToCounterPartyIdCmd.Flags().StringP(utils.CounterpartyIdFlag, "c", "", "ID of counterparty (Required)")
	createWithdrawalToCounterPartyIdCmd.Flags().StringP(utils.AssetIdFlag, "i", "", "ID of asset to be withdrawn (Required)")
	createWithdrawalToCounterPartyIdCmd.Flags().StringP(utils.AmountFlag, "a", "", "Amount of asset to be withdrawn (Required)")
	createWithdrawalToCounterPartyIdCmd.Flags().StringP(utils.NonceFlag, "n", "", "Nonce for withdrawal")
	createWithdrawalToCounterPartyIdCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	createWithdrawalToCounterPartyIdCmd.MarkFlagRequired(utils.CounterpartyIdFlag)
	createWithdrawalToCounterPartyIdCmd.MarkFlagRequired(utils.AssetIdFlag)
	createWithdrawalToCounterPartyIdCmd.MarkFlagRequired(utils.AmountFlag)
}
