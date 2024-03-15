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

var updatePortfolioCmd = &cobra.Command{
	Use:   "update-portfolio",
	Short: "update a portfolio name.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, portfolioId, err := utils.InitClientAndPortfolioId(cmd, true)
		if err != nil {
			return fmt.Errorf("cannot initialize from environment: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.UpdatePortfolioRequest{
			Name:        utils.GetFlagStringValue(cmd, utils.NameFlag),
			PortfolioId: portfolioId,
		}

		response, err := client.UpdatePortfolio(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot update portfolio: %w", err)
		}

		return utils.PrintJsonResponse(cmd, response)
	},
}

func init() {
	rootCmd.AddCommand(updatePortfolioCmd)

	updatePortfolioCmd.Flags().StringP(utils.PortfolioIdFlag, "i", "", "Portfolio ID. Uses environment variable if blank (Required)")
	updatePortfolioCmd.Flags().StringP(utils.NameFlag, "n", "", "New name of the portfolio (Required)")
	updatePortfolioCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	updatePortfolioCmd.MarkFlagRequired(utils.PortfolioIdFlag)
	updatePortfolioCmd.MarkFlagRequired(utils.NameFlag)

}
