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

var getPortfolioFillsCmd = &cobra.Command{
	Use:   "get-portfolio-fills",
	Short: "Get portfolio fills.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, portfolioId, err := utils.InitClientAndPortfolioId(cmd, true)
		if err != nil {
			return fmt.Errorf("cannot initialize from environment: %w", err)
		}

		resultLimit, _ := utils.GetFlagIntValue(cmd, utils.ResultLimitFlag)
		resultOffset, _ := utils.GetFlagIntValue(cmd, utils.ResultOffsetFlag)

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.GetPortfolioFillsRequest{
			PortfolioId:   portfolioId,
			OrderId:       utils.GetFlagStringValue(cmd, utils.OrderIdFlag),
			ClientOrderId: utils.GetFlagStringValue(cmd, utils.ClientOrderIdFlag),
			RefDatetime:   utils.GetFlagStringValue(cmd, utils.RefDatetimeFlag),
			ResultLimit:   resultLimit,
			ResultOffset:  resultOffset,
			TimeFrom:      utils.GetFlagStringValue(cmd, utils.TimeFromFlag),
		}

		request.Pagination = utils.CreatePaginationParams(request.RefDatetime, request.ResultLimit, request.ResultOffset)

		response, err := client.GetPortfolioFills(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot get portfolio fills: %w", err)
		}

		return utils.PrintJsonResponse(cmd, response)
	},
}

func init() {
	cmdConfigs := []utils.CommandConfig{
		{
			Command: getPortfolioFillsCmd,
			FlagConfig: []utils.FlagConfig{
				{
					FlagName:     utils.PortfolioIdFlag,
					Shorthand:    "p",
					Usage:        "Portfolio ID (required). Uses environment variable if blank",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.OrderIdFlag,
					Shorthand:    "i",
					Usage:        "Order ID",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.ClientOrderIdFlag,
					Shorthand:    "c",
					Usage:        "Client Order ID",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.RefDatetimeFlag,
					Shorthand:    "r",
					Usage:        "Reference datetime for the request",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.ResultLimitFlag,
					DefaultValue: utils.ZeroInt,
					Usage:        "Result limit",
					Required:     false,
				},
				{
					FlagName:     utils.ResultOffsetFlag,
					DefaultValue: utils.ZeroInt,
					Usage:        "Result offset",
					Required:     false,
				},
				{
					FlagName:  utils.TimeFromFlag,
					Shorthand: "t",
					Usage:     "Time from which to get fills",
					Required:  false,
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
