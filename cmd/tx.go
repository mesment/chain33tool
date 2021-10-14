/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"chain33tool/util"
)

func init() {
	rootCmd.AddCommand(TxCmd())
}

// txCmd represents the tx command
func TxCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "tx",
		Short: " tx manage",
		Args: cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(AlterCmd(),DecodeTxCmd())

	return cmd
}

func AlterCmd() *cobra.Command  {
	var cmd = &cobra.Command{
		Use:  "alter",
		Short: "alter tx info",
		Run:AlterTransaction,
	}

	addAlterFlags(cmd)
	return cmd
}

func DecodeTxCmd() *cobra.Command  {
	var cmd = &cobra.Command{
		Use:  "decode",
		Short: "decode tx info",
		Run:DecodeTransaction,
	}

	addDecodeFlags(cmd)
	return cmd
}

func addAlterFlags(cmd *cobra.Command)  {
	cmd.Flags().StringP("data", "d", "", "signed transaction data")
	cmd.Flags().StringP("to", "t", "", "receiver account address")
	cmd.Flags().Float64P("amount", "v", 0, "transaction amount")
}

func addDecodeFlags(cmd *cobra.Command)  {
	cmd.Flags().StringP("data", "d", "", "signed transaction data")
}


func DecodeTransaction(cmd *cobra.Command, args[]string)  {
	signedTx,_ := cmd.Flags().GetString("data")

	if signedTx == "" {
		fmt.Println("data 不能为空")
		fmt.Println(cmd.UsageString())
	}
	err := util.PrintHexTxDetail(signedTx)
	if err != nil {
		fmt.Println(err)
	}
}


func AlterTransaction(cmd *cobra.Command, args[]string)  {
	signedTx,_ := cmd.Flags().GetString("data")
	toAddr, _ := cmd.Flags().GetString("to")
	amt, _ := cmd.Flags().GetFloat64("amount")

	if signedTx == "" {
		fmt.Println("data 不能为空")
		fmt.Println(cmd.UsageString())
	}
	tx, err := util.DecodeHexTransaction(signedTx)
	if err != nil {
		fmt.Println(err)
		return
	}

	if amt > 0 && toAddr != "" {
		newtx, err := util.ChangeTx(tx, toAddr, amt)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("result in HEX:", util.EncodeToString(newtx))
		return
	} else if amt > 0 {
		newtx, err := util.ChangeTxAmt(tx, amt)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("result in HEX:", util.EncodeToString(newtx))
		return
	} else if toAddr != "" {
		newtx:= util.ChangeTxToAddress(tx, toAddr)
		fmt.Println("result in HEX:", util.EncodeToString(newtx))
		return
	}
	fmt.Println(cmd.UsageString())
}
