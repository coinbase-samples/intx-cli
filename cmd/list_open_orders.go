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

var listOpenOrdersCmd = &cobra.Command{
	Use:   "list-open-orders",
	Short: "List open orders based on filter criteria.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, portfolioId, err := utils.InitClientAndPortfolioId(cmd, true)
		if err != nil {
			return fmt.Errorf("cannot initialize from environment: %w", err)
		}

		resultLimit, _ := utils.GetFlagIntValue(cmd, utils.ResultLimitFlag)
		resultOffset, _ := utils.GetFlagIntValue(cmd, utils.ResultOffsetFlag)

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.ListOpenOrdersRequest{
			PortfolioId:   portfolioId,
			InstrumentId:  utils.GetFlagStringValue(cmd, utils.InstrumentIdFlag),
			ClientOrderId: utils.GetFlagStringValue(cmd, utils.ClientOrderIdFlag),
			EventType:     utils.GetFlagStringValue(cmd, utils.EventTypeFlag),
			RefDatetime:   utils.GetFlagStringValue(cmd, utils.RefDatetimeFlag),
			ResultLimit:   resultLimit,
			ResultOffset:  resultOffset,
		}

		request.Pagination = utils.CreatePaginationParams(request.RefDatetime, request.ResultLimit, request.ResultOffset)

		response, err := client.ListOpenOrders(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot list open orders: %w", err)
		}

		return utils.PrintJsonResponse(cmd, response)
	},
}

func init() {
	cmdConfigs := []utils.CommandConfig{
		{
			Command: listOpenOrdersCmd,
			FlagConfig: []utils.FlagConfig{
				{
					FlagName:     utils.PortfolioIdFlag,
					Shorthand:    "p",
					Usage:        "Portfolio ID(s). Uses environment variable if blank",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.InstrumentIdFlag,
					Shorthand:    "i",
					Usage:        "Filter open orders by instrument ID",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.ClientOrderIdFlag,
					Shorthand:    "c",
					Usage:        "Filter open orders by client order ID",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.EventTypeFlag,
					Shorthand:    "e",
					Usage:        "Filter open orders by event type",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.RefDatetimeFlag,
					Usage:        "Filter open orders by datetime",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.ResultLimitFlag,
					DefaultValue: 0,
					Usage:        "Limit the number of open orders returned",
					Required:     false,
				},
				{
					FlagName:     utils.ResultOffsetFlag,
					DefaultValue: 0,
					Usage:        "Offset for the list of open orders returned",
					Required:     false,
				},
				{
					FlagName:     utils.TimeFromFlag,
					Usage:        "Filter open orders from this time",
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
