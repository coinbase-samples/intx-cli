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

		jsonResponse, err := utils.FormatResponseAsJson(cmd, response)
		if err != nil {
			return err
		}

		fmt.Println(jsonResponse)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createPortfolioTransferCmd)

	createPortfolioTransferCmd.Flags().StringP(utils.ToFlag, "t", "", "Name of destination portfolio (Required)")
	createPortfolioTransferCmd.Flags().StringP(utils.AssetIdFlag, "i", "", "ID of asset to be transferred (Required)")
	createPortfolioTransferCmd.Flags().StringP(utils.AmountFlag, "a", "", "Amount of asset to be transferred (Required)")
	createPortfolioTransferCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	createPortfolioTransferCmd.MarkFlagRequired(utils.ToFlag)
	createPortfolioTransferCmd.MarkFlagRequired(utils.AssetIdFlag)
	createPortfolioTransferCmd.MarkFlagRequired(utils.AmountFlag)
}
