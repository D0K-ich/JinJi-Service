package rest

const (
	//User section
	TableUsers          	= "users"
	TableUserSessions   	= "user_sessions"
	TableAdmins         	= "admins"
	TableAdminSessions  	= "admin_sessions"
	TablePromoCodes     	= "promocodes"
	TableTariffs        	= "tariffs"
	TableUsersPromocodes	= "users_promocodes"
	TableUsersSites     	= "users_sites"

	CookieNameUser  		= "user"
	CookieNameAdmin 		= "admin"

	StatusCreated 			= "created"
	StatusPending 			= "pending" // FOR ECOM
	StatusFailed  			= "failed"
	StatusSucceed 			= "succeed"

	TypeEcom    			= "ecom"
	TypeBilling 			= "billing"

	TypeBillOut 			= "bill-out"
	TypeBuySubscription 	= "buy_subscription"

	StateUnconfirmed 		= "unconfirmed"
	StateBlocked     		= "blocked"
	StateActive      		= "active"

	ColumnTariffId        	= "tariff_id"
	ColumnAuthCode        	= "auth_code"
	ColumnBalance         	= "balance"
	ColumnTariffBalance   	= "tariff_balance"
	ColumnUserSettings      = "settings"
	ColumnPrivate         	= "private"
	ColumnUsage           	= "usage"
	ColumnState           	= "state"
	ColumnTariffExpiration	= "tariff_expiration"
	ColumnSld				= "sld"
)
