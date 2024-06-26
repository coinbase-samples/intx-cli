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

package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coinbase-samples/intx-sdk-go"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strconv"
	"time"
)

func getDefaultTimeoutDuration() time.Duration {
	envTimeout := os.Getenv("intxCliTimeout")
	if envTimeout != "" {
		if value, err := strconv.Atoi(envTimeout); err == nil && value > 0 {
			return time.Duration(value) * time.Second
		}
	}
	return 7 * time.Second
}

func GetContextWithTimeout() (context.Context, context.CancelFunc) {
	timeoutDuration := getDefaultTimeoutDuration()
	return context.WithTimeout(context.Background(), timeoutDuration)
}

func GetClientFromEnv() (*intx.Client, error) {
	credentials := &intx.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("INTX_CREDENTIALS")), credentials); err != nil {
		return nil, fmt.Errorf("cannot unmarshal credentials: %w", err)
	}

	client := intx.NewClient(credentials, http.Client{})
	return client, nil
}

func InitClientAndPortfolioId(cmd *cobra.Command, needPortfolioId bool) (client *intx.Client, portfolioId string, err error) {
	client, err = GetClientFromEnv()
	if err != nil {
		err = fmt.Errorf("cannot get client from environment: %w", err)
		return
	}

	if needPortfolioId {
		portfolioId, err = GetPortfolioId(cmd, client)
		if err != nil {
			return
		}
	}

	return
}

func GetFlagStringValue(cmd *cobra.Command, flagName string) string {
	value, _ := cmd.Flags().GetString(flagName)
	return value
}

func GetFlagBoolValue(cmd *cobra.Command, flagName string) *bool {
	value, err := cmd.Flags().GetBool(flagName)
	if err != nil {
		return nil
	}
	return &value
}

func MarshalJson(data interface{}, format bool) ([]byte, error) {
	if format {
		return json.MarshalIndent(data, "", JsonIndent)
	}
	return json.Marshal(data)
}

func CheckFormatFlag(cmd *cobra.Command) (bool, error) {
	formatFlagValue, err := cmd.Flags().GetBool(FormatFlag)
	if err != nil {
		return false, fmt.Errorf("cannot read format flag: %w", err)
	}
	return formatFlagValue, nil
}

func GetPortfolioId(cmd *cobra.Command, client *intx.Client) (string, error) {
	portfolioId, err := cmd.Flags().GetString(PortfolioIdFlag)
	if err != nil {
		return "", fmt.Errorf("error retrieving portfolio ID: %w", err)
	}

	if portfolioId == "" {
		if client == nil || client.Credentials == nil {
			return "", errors.New("client or client credentials are nil")
		}
		portfolioId = client.Credentials.PortfolioId
		if portfolioId == "" {
			return "", errors.New("portfolio ID is not provided in both flag and client credentials")
		}
	}

	return portfolioId, nil
}

func PrintJsonResponse(cmd *cobra.Command, response interface{}) error {
	jsonResponse, err := FormatResponseAsJson(cmd, response)
	if err != nil {
		return err
	}

	fmt.Println(jsonResponse)
	return nil
}

func FormatResponseAsJson(cmd *cobra.Command, response interface{}) (string, error) {
	shouldFormat, err := CheckFormatFlag(cmd)
	if err != nil {
		return "", err
	}

	jsonResponse, err := MarshalJson(response, shouldFormat)
	if err != nil {
		return "", fmt.Errorf("cannot marshal response to JSON: %w", err)
	}

	return string(jsonResponse), nil
}

func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func GetFlagIntValue(cmd *cobra.Command, flagName string) (int, error) {
	valueStr, err := cmd.Flags().GetString(flagName)
	if err != nil {
		return 0, err
	}
	if valueStr == "" {
		return 0, nil
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func CreatePaginationParams(refDatetime string, resultLimit, resultOffset int) *intx.PaginationParams {
	if refDatetime != "" || resultLimit != ZeroInt || resultOffset != ZeroInt {
		return &intx.PaginationParams{
			RefDatetime:  refDatetime,
			ResultLimit:  resultLimit,
			ResultOffset: resultOffset,
		}
	}
	return nil
}

type CommandConfig struct {
	Command    *cobra.Command
	FlagConfig []FlagConfig
}

type FlagConfig struct {
	FlagName     string
	Shorthand    string
	Usage        string
	DefaultValue interface{}
	Required     bool
}

func RegisterCommandConfigs(root *cobra.Command, cmdConfigs []CommandConfig) {
	for _, config := range cmdConfigs {
		for _, flag := range config.FlagConfig {
			switch defaultValue := flag.DefaultValue.(type) {
			case string:
				config.Command.Flags().StringP(flag.FlagName, flag.Shorthand, defaultValue, flag.Usage)
			case int:
				config.Command.Flags().IntP(flag.FlagName, flag.Shorthand, defaultValue, flag.Usage)
			case bool:
				config.Command.Flags().BoolP(flag.FlagName, flag.Shorthand, defaultValue, flag.Usage)
			}

			if flag.Required {
				if err := config.Command.MarkFlagRequired(flag.FlagName); err != nil {
					fmt.Printf("could not mark flag %s as required: %v\n", flag.FlagName, err)
				}
			}
			root.AddCommand(config.Command)
		}
	}
}
