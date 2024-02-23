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

var getHistoricalFundingRatesCmd = &cobra.Command{
	Use:   "get-historical-funding-rates",
	Short: "Get historical funding rates for instrument.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := utils.GetClientFromEnv()
		if err != nil {
			return fmt.Errorf("failed to initialize client: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.GetHistoricalFundingRequest{
			InstrumentId: utils.GetFlagStringValue(cmd, utils.InstrumentIdFlag),
		}

		response, err := client.GetHistoricalFundingRates(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot get funding rates: %w", err)
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
	rootCmd.AddCommand(getHistoricalFundingRatesCmd)

	getHistoricalFundingRatesCmd.Flags().StringP(utils.InstrumentIdFlag, "i", "", "Instrument ID (Required)")
	getHistoricalFundingRatesCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getHistoricalFundingRatesCmd.MarkFlagRequired(utils.InstrumentIdFlag)
}
