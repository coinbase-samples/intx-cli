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
		client, portfolioId, err := utils.InitClientAndPortfolioId(cmd, true)
		if err != nil {
			return fmt.Errorf("cannot initialize from environment: %w", err)
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

		return utils.PrintJsonResponse(cmd, response)
	},
}

func init() {
	cmdConfigs := []utils.CommandConfig{
		{
			Command: modifyOrderCmd,
			FlagConfig: []utils.FlagConfig{
				{
					FlagName:     utils.OrderIdFlag,
					Shorthand:    "i",
					Usage:        "Order ID (Required)",
					DefaultValue: "",
					Required:     true,
				},
				{
					FlagName:     utils.SizeFlag,
					Shorthand:    "s",
					Usage:        "Order size in base asset units",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.LimitPriceFlag,
					Shorthand:    "l",
					Usage:        "Limit price for the order",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.StopPriceFlag,
					Shorthand:    "p",
					Usage:        "Stop price for the order",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.PortfolioIdFlag,
					Shorthand:    "o",
					Usage:        "Portfolio ID. Uses environment variable if blank",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.ClientOrderIdFlag,
					Shorthand:    "c",
					Usage:        "Client order id value",
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
