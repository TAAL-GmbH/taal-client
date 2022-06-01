<script>
  import { onMount } from 'svelte'
  import {TxDataSize} from '../tx_history/transaction_format_functions.svelte'
  let transactions
  let nrOfTransactions = 0
  let combinedSize = 0.0

  onMount(async () => {
    await fetch(`${BASE_URL}/api/v1/transactions/info`)
      .then((r) => r.json())
      .then((data) => {
        transactions = data.transactions
        nrOfTransactions = transactions.length

        var dataSizes = transactions.map((a) => a.data_bytes)
        dataSizes.forEach(element => {
          combinedSize += element
        });
      })
  })
</script>

<div class="field">
  <h1>Total number of transactions: {nrOfTransactions}</h1>
  <h1>Total size of transaction data: {TxDataSize(combinedSize)}</h1>
</div>
