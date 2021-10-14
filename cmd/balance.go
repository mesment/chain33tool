package cmd

import (
	"fmt"
	"github.com/33cn/chain33/types"
	"github.com/spf13/cobra"
	"chain33tool/util"
)




func init() {
	rootCmd.AddCommand(BalanceCmd())
}

// balanceCmd represents the balance command
func BalanceCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "balance",
		Short: "balance manage",
		Args: cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(BalanceGetCmd(),BalanceSetCmd()  )

	return cmd
}

func BalanceGetCmd() *cobra.Command  {
	var cmd = &cobra.Command{
		Use:  "get",
		Short: "get address balance",
		Run: GetBalance,
	}

	addBalanceGetFlags(cmd)
	return cmd
}

func BalanceSetCmd() *cobra.Command  {
	var cmd = &cobra.Command{
		Use:  "set",
		Short: "set address balance",
		Run: SetBalance,
	}

	addBalanceSetFlags(cmd)
	return cmd
}

func addBalanceGetFlags(cmd *cobra.Command)  {
	cmd.Flags().StringP("addr", "a", "", " account address")
}

func addBalanceSetFlags(cmd *cobra.Command)  {
	cmd.Flags().StringP("addr", "a", "", " account address")
	cmd.Flags().Float64P("amount", "v", 0, " account new balance")
}

func GetBalance(cmd *cobra.Command, args[]string)  {

	chain33Cfg := types.NewChain33Config(types.MergeCfg(types.ReadFile(cfgFile), ""))
	util.InitializeDB(chain33Cfg)
	defer util.CloseDB()

	addr, _ := cmd.Flags().GetString("addr")

	if addr == "" {
		fmt.Println("addr 不能为空")
		fmt.Println(cmd.UsageString())
		return
	}
	balance, err := util.GetDBBalance(addr)
	if err!= nil {
		fmt.Println(err)
		return
	}
	fmt.Println(balance)
}


func SetBalance(cmd *cobra.Command, args[]string)  {
	chain33Cfg := types.NewChain33Config(types.MergeCfg(types.ReadFile(cfgFile), ""))
	util.InitializeDB(chain33Cfg)
	defer util.CloseDB()
	addr, _ := cmd.Flags().GetString("addr")
	amt, _ := cmd.Flags().GetFloat64("amount")
	if addr == "" {
		fmt.Println("addr 不能为空")
		fmt.Println(cmd.UsageString())
		return
	}
	if amt <= 0 {
		fmt.Println("amount必须大于0")
		fmt.Println(cmd.UsageString())
		return
	}
	err := util.AlterAccountBalance(addr, amt)
	if err!= nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}