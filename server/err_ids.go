package server

const (
	errReadTxNotAuthorized                      = 1
	errWriteFailedToGetTxs                      = 2
	errWriteTxNotAuthorized                     = 3
	errWriteApiKeyUnknown                       = 4
	errWriteFailedToReadApiKey                  = 5
	errWriteEmptyMimeType                       = 10
	errWriteEmptyBody                           = 11
	errWriteCouldNotReadBody                    = 12
	errReadTx                                   = 13
	errWriteFailedToReturnOpReturnOutput        = 14
	errWriteFailedToSignTx                      = 15
	errWriteFailedToSubmitTxs                   = 16
	errRegisterFailedToCheckApiKey              = 17
	errRegisterApiKeyHasAlreadyBeenRegistered   = 18
	errRegisterFailedToGenerateNewPrivateKey    = 19
	errRegisterFailedToSignPrivateKeywithApiKey = 20
	errRegisterFailedToRegisterApiKey           = 21
	errRegisterFailedToGetKeysFromPrivateKey    = 22
	errRegisterFailedToStoreKeys                = 23
	errAPIKeysFailedToGetKeys                   = 24
	errWriteInsertTransaction                   = 25
	errAPIKeysRevoke                            = 26
	errGetTransactions                          = 27
)
