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

var getMarginOverrideCmd = &cobra.Command{
	Use:   "set-margin-override",
	Short: "Set margin override for portfolio.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, portfolioId, err := utils.InitClientAndPortfolioId(cmd, true)
		if err != nil {
			return fmt.Errorf("cannot initialize from environment: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.SetMarginOverrideRequest{
			PortfolioId:    portfolioId,
			MarginOverride: utils.GetFlagStringValue(cmd, utils.MarginOverrideFlag),
		}

		response, err := client.SetMarginOverride(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot set margin override: %w", err)
		}

		return utils.PrintJsonResponse(cmd, response)
	},
}

func init() {
	rootCmd.AddCommand(getMarginOverrideCmd)

	getMarginOverrideCmd.Flags().StringP(utils.PortfolioIdFlag, "i", "", "Portfolio ID. Uses environment variable if blank")
	getMarginOverrideCmd.Flags().StringP(utils.MarginOverrideFlag, "m", "", "Margin Override string")
	getMarginOverrideCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
