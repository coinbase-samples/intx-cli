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

var validateCounterPartyIdCmd = &cobra.Command{
	Use:   "validate-counterparty-id",
	Short: "Validate a counterparty ID.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := utils.GetClientFromEnv()
		if err != nil {
			return fmt.Errorf("failed to initialize client: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.ValidateCounterpartyIdRequest{
			CounterpartyId: utils.GetFlagStringValue(cmd, utils.CounterpartyIdFlag),
		}

		response, err := client.ValidateCounterpartyId(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot validate counterparty ID: %w", err)
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
	rootCmd.AddCommand(validateCounterPartyIdCmd)

	validateCounterPartyIdCmd.Flags().StringP(utils.CounterpartyIdFlag, "i", "", "Counterparty ID ot be validated (Required)")
	validateCounterPartyIdCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	validateCounterPartyIdCmd.MarkFlagRequired(utils.CounterpartyIdFlag)
}
