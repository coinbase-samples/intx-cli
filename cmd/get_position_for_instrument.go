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

var getPositionForInstrumentCmd = &cobra.Command{
	Use:   "get-position-for-instrument",
	Short: "Get position for instrument.",
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

		request := &intx.GetInstrumentPositionRequest{
			PortfolioId:  portfolioId,
			InstrumentId: utils.GetFlagStringValue(cmd, utils.InstrumentIdFlag),
		}

		response, err := client.GetInstrumentPosition(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot get instrument position: %w", err)
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
	rootCmd.AddCommand(getPositionForInstrumentCmd)

	getPositionForInstrumentCmd.Flags().StringP(utils.PortfolioIdFlag, "p", "", "Portfolio ID. Uses environment variable if blank")
	getPositionForInstrumentCmd.Flags().StringP(utils.InstrumentIdFlag, "i", "", "ID of the Instrument, e.g. ETH-USDC (Required)")
	getPositionForInstrumentCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
