<script>
  import { onMount } from 'svelte'
  let transactions
  let nrOfTransactions = 0

  onMount(async () => {
    await fetch(`${BASE_URL}/api/v1/transactions/info`)
      .then((r) => r.json())
      .then((data) => {
        transactions = data.transactions
        nrOfTransactions = transactions.unshift(0)
      })
  })
</script>

<div class="field">
  <h1>Total number of transactions: {nrOfTransactions}</h1>
</div>
