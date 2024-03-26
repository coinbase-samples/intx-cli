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

var listTransfersCmd = &cobra.Command{
	Use:   "list-transfers",
	Short: "List transfers based on filter criteria.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, portfolioIds, err := utils.InitClientAndPortfolioId(cmd, true)
		if err != nil {
			return fmt.Errorf("cannot initialize from environment: %w", err)
		}

		timeFrom, _ := cmd.Flags().GetString(utils.TimeFromFlag)
		timeTo, _ := cmd.Flags().GetString(utils.TimeToFlag)
		status, _ := cmd.Flags().GetString(utils.StatusFlag)
		transferType, _ := cmd.Flags().GetString(utils.TypeFlag)

		resultLimit, _ := utils.GetFlagIntValue(cmd, utils.ResultLimitFlag)
		resultOffset, _ := utils.GetFlagIntValue(cmd, utils.ResultOffsetFlag)

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.ListTransfersRequest{
			PortfolioIds: portfolioIds,
			TimeFrom:     timeFrom,
			TimeTo:       timeTo,
			Status:       status,
			Type:         transferType,
			ResultLimit:  resultLimit,
			ResultOffset: resultOffset,
		}

		if request.TimeFrom != "" || request.TimeTo != "" || request.ResultLimit != utils.ZeroInt || request.ResultOffset != utils.ZeroInt {
			request.Pagination = &intx.PaginationParams{
				RefDatetime:  request.TimeFrom,
				ResultLimit:  request.ResultLimit,
				ResultOffset: request.ResultOffset,
			}
		}

		response, err := client.ListTransfers(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot list transfers: %w", err)
		}

		return utils.PrintJsonResponse(cmd, response)
	},
}

func init() {
	cmdConfigs := []utils.CommandConfig{
		{
			Command: listTransfersCmd,
			FlagConfig: []utils.FlagConfig{
				{
					FlagName:     utils.PortfolioIdFlag,
					Shorthand:    "p",
					Usage:        "Portfolio ID(s). Uses environment variable if blank, supports comma-separated values for multiple IDs",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.TimeFromFlag,
					Shorthand:    "f",
					Usage:        "Filter transfers from this time",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.TimeToFlag,
					Shorthand:    "t",
					Usage:        "Filter transfers to this time",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.StatusFlag,
					Shorthand:    "s",
					Usage:        "Filter transfers by status",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.TypeFlag,
					Shorthand:    "y",
					Usage:        "Filter transfers by type",
					DefaultValue: "",
					Required:     false,
				},
				{
					FlagName:     utils.ResultLimitFlag,
					DefaultValue: 0,
					Usage:        "Limit the number of transfers returned",
					Required:     false,
				},
				{
					FlagName:     utils.ResultOffsetFlag,
					DefaultValue: 0,
					Usage:        "Offset for the list of transfers returned",
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
