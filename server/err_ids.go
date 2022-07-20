package server

const (
	errReadTxNotAuthorized                       = 1
	errWriteFailedToGetTxs                       = 2
	errWriteTxNotAuthorized                      = 3
	errWriteApiKeyUnknown                        = 4
	errWriteFailedToReadApiKey                   = 5
	errWriteEmptyMimeType                        = 10
	errWriteEmptyBody                            = 11
	errWriteCouldNotReadBody                     = 12
	errReadTx                                    = 13
	errWriteFailedToReturnOpReturnOutput         = 14
	errWriteFailedToSignTx                       = 15
	errWriteFailedToSubmitTxs                    = 16
	errRegisterFailedToCheckApiKey               = 17
	errRegisterApiKeyHasAlreadyBeenRegistered    = 18
	errRegisterFailedToGenerateNewPrivateKey     = 19
	errRegisterFailedToSignPrivateKeywithApiKey  = 20
	errRegisterFailedToRegisterApiKey            = 21
	errRegisterFailedToGetKeysFromPrivateKey     = 22
	errRegisterFailedToStoreKeys                 = 23
	errAPIKeysFailedToGetKeys                    = 24
	errWriteInsertTransaction                    = 25
	errAPIKeysRevoke                             = 26
	errGetTransactions                           = 27
	errHealthDB                                  = 28
	errGetTransactionsHoursBackMustBeInteger     = 29
	errGetTransactionsHoursBackMustBePositive    = 30
	errGetTransactionInfo                        = 31
	errGetTransactionInfoHoursBackMustBeInteger  = 32
	errGetTransactionInfoHoursBackMustBePositive = 33
	errGetTransactionInfoGetGranularity          = 34
	errGetTransactionInfoFillDates               = 35
	errPutSettingsUpdateSettings                 = 36
	errPutSettingsGetDuration                    = 37
	errPutSettingsBind                           = 38
	errPutSettingsGetJson                        = 39
)
