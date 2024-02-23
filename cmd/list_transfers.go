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
		client, err := utils.GetClientFromEnv()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		portfolioIds, err := utils.GetPortfolioId(cmd, client)
		if err != nil {
			return fmt.Errorf("failed to get portfolio ID: %w", err)
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

		jsonResponse, err := utils.FormatResponseAsJson(cmd, response)
		if err != nil {
			return err
		}

		fmt.Println(jsonResponse)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listTransfersCmd)

	listTransfersCmd.Flags().StringP(utils.PortfolioIdFlag, "p", "", "Portfolio ID(s). Uses environment variable if blank, supports comma-separated values for multiple IDs")
	listTransfersCmd.Flags().StringP(utils.TimeFromFlag, "f", "", "Filter transfers from this time")
	listTransfersCmd.Flags().StringP(utils.TimeToFlag, "t", "", "Filter transfers to this time")
	listTransfersCmd.Flags().StringP(utils.StatusFlag, "s", "", "Filter transfers by status")
	listTransfersCmd.Flags().StringP(utils.TypeFlag, "y", "", "Filter transfers by type")
	listTransfersCmd.Flags().Int(utils.ResultLimitFlag, 0, "Limit the number of transfers returned")
	listTransfersCmd.Flags().Int(utils.ResultOffsetFlag, 0, "Offset for the list of transfers returned")
	listTransfersCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
