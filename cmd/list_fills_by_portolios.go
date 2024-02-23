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

var listFillsCmd = &cobra.Command{
	Use:   "list-fills",
	Short: "List fills by portfolios.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := utils.GetClientFromEnv()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		portfolioIds, err := utils.GetPortfolioId(cmd, client)
		if err != nil {
			return fmt.Errorf("failed to get portfolio ID: %w", err)
		}

		resultLimit, _ := utils.GetFlagIntValue(cmd, utils.ResultLimitFlag)
		resultOffset, _ := utils.GetFlagIntValue(cmd, utils.ResultOffsetFlag)
		timeFrom, _ := cmd.Flags().GetString(utils.TimeFromFlag)

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.ListFillsByPortfoliosRequest{
			PortfolioIds:  portfolioIds,
			OrderId:       utils.GetFlagStringValue(cmd, utils.OrderIdFlag),
			ClientOrderId: utils.GetFlagStringValue(cmd, utils.ClientOrderIdFlag),
			RefDatetime:   utils.GetFlagStringValue(cmd, utils.RefDatetimeFlag),
			ResultLimit:   resultLimit,
			ResultOffset:  resultOffset,
			TimeFrom:      timeFrom,
		}

		request.Pagination = utils.CreatePaginationParams(request.RefDatetime, request.ResultLimit, request.ResultOffset)

		response, err := client.ListFillsByPortfolios(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot list fills: %w", err)
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
	rootCmd.AddCommand(listFillsCmd)

	listFillsCmd.Flags().StringP(utils.PortfolioIdFlag, "i", "", "Portfolio ID(s). Uses environment variable if blank")
	listFillsCmd.Flags().StringP(utils.OrderIdFlag, "o", "", "Filter fills by order ID")
	listFillsCmd.Flags().StringP(utils.ClientOrderIdFlag, "c", "", "Filter fills by client order ID")
	listFillsCmd.Flags().StringP(utils.RefDatetimeFlag, "", "", "Filter fills by reference datetime")
	listFillsCmd.Flags().Int(utils.ResultLimitFlag, 0, "Limit the number of fills returned")
	listFillsCmd.Flags().Int(utils.ResultOffsetFlag, 0, "Offset for the list of fills returned")
	listFillsCmd.Flags().StringP(utils.TimeFromFlag, "", "", "Filter fills from this time")
	listFillsCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
