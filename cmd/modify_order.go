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

var modifyOrderCmd = &cobra.Command{
	Use:   "modify-order",
	Short: "Submit an order modification request.",
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

		request := &intx.ModifyOrderRequest{
			OrderId:       utils.GetFlagStringValue(cmd, utils.OrderIdFlag),
			ClientOrderId: utils.GetFlagStringValue(cmd, utils.ClientOrderIdFlag),
			PortfolioId:   portfolioId,
			Size:          utils.GetFlagStringValue(cmd, utils.SizeFlag),
			Price:         utils.GetFlagStringValue(cmd, utils.LimitPriceFlag),
			StopPrice:     utils.GetFlagStringValue(cmd, utils.StopPriceFlag),
		}

		response, err := client.ModifyOrder(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot modify order: %w", err)
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
	rootCmd.AddCommand(modifyOrderCmd)

	modifyOrderCmd.Flags().StringP(utils.OrderIdFlag, "i", "", "Order ID (Required)")
	modifyOrderCmd.Flags().StringP(utils.SizeFlag, "s", "", "Order size in base asset units")
	modifyOrderCmd.Flags().StringP(utils.LimitPriceFlag, "l", "", "Limit price for the order")
	modifyOrderCmd.Flags().StringP(utils.StopPriceFlag, "p", "", "Stop price for the order")
	modifyOrderCmd.Flags().StringP(utils.PortfolioIdFlag, "o", "", "Portfolio ID. Uses environment variable if blank")
	modifyOrderCmd.Flags().StringP(utils.ClientOrderIdFlag, "c", "", "Client order id value")
	modifyOrderCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	modifyOrderCmd.MarkFlagRequired(utils.OrderIdFlag)
}
