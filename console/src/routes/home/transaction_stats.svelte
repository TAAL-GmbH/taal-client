<script>
  import { TxDataSize } from '../tx_history/transaction_format_functions.svelte'
  import TransactionChart from './transaction_chart.svelte'

  import TransactionsInfo from '../util/transactions_info.svelte'

  let transactions
  let nrOfTransactions = 0
  let combinedSize = 0

  function OnTransactionsChange(transactions) {
    if (transactions == null) {
      return
    }
    nrOfTransactions = transactions.length
    var dataSizes = transactions.map((tx) => tx.data_bytes)

    if (dataSizes.length == 0) {
      combinedSize = 0
    }

    dataSizes.forEach((element) => {
      combinedSize += element
    })
  }

  $: OnTransactionsChange(transactions)
</script>

<div class="field">
  <TransactionsInfo bind:transactions />
  <h1>Total number of transactions: {nrOfTransactions}</h1>
  <h1>Total size of transaction data: {TxDataSize(combinedSize)}</h1>
  <TransactionChart bind:transactions />
</div>
