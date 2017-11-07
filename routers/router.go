// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"microservice-wallet-server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/JOB_EXECUTION_LOG",
			beego.NSInclude(
				&controllers.JOBEXECUTIONLOGController{},
			),
		),

		beego.NSNamespace("/JOB_STATUS_TRACE_LOG",
			beego.NSInclude(
				&controllers.JOBSTATUSTRACELOGController{},
			),
		),

		beego.NSNamespace("/job_execution_log",
			beego.NSInclude(
				&controllers.JobExecutionLogController{},
			),
		),

		beego.NSNamespace("/job_status_trace_log",
			beego.NSInclude(
				&controllers.JobStatusTraceLogController{},
			),
		),

		beego.NSNamespace("/tb_act_account",
			beego.NSInclude(
				&controllers.TbActAccountController{},
			),
		),

		beego.NSNamespace("/tb_act_block",
			beego.NSInclude(
				&controllers.TbActBlockController{},
			),
		),

		beego.NSNamespace("/tb_act_contract_abi",
			beego.NSInclude(
				&controllers.TbActContractAbiController{},
			),
		),

		beego.NSNamespace("/tb_act_contract_event",
			beego.NSInclude(
				&controllers.TbActContractEventController{},
			),
		),

		beego.NSNamespace("/tb_act_contract_info",
			beego.NSInclude(
				&controllers.TbActContractInfoController{},
			),
		),

		beego.NSNamespace("/tb_act_contract_storage",
			beego.NSInclude(
				&controllers.TbActContractStorageController{},
			),
		),

		beego.NSNamespace("/tb_act_transaction",
			beego.NSInclude(
				&controllers.TbActTransactionController{},
			),
		),

		beego.NSNamespace("/tb_act_transaction_ex",
			beego.NSInclude(
				&controllers.TbActTransactionExController{},
			),
		),

		beego.NSNamespace("/tb_act_withdraw",
			beego.NSInclude(
				&controllers.TbActWithdrawController{},
			),
		),

		beego.NSNamespace("/tb_exchange_wallet_config",
			beego.NSInclude(
				&controllers.TbExchangeWalletConfigController{},
			),
		),

		beego.NSNamespace("/tb_user_address",
			beego.NSInclude(
				&controllers.TbUserAddressController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
