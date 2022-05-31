<script>
  // svelte-ignore unused-export-let
  export let location

  import { onMount } from 'svelte'
  import TransactionRows from './transaction_rows.svelte'

  let transactions

  onMount(async () => {
    await fetch(`${BASE_URL}/api/v1/transactions/info`)
      .then((r) => r.json())
      .then((data) => {
        transactions = data.transactions
      })
  })
</script>

<h1>Transactions</h1>

{#if transactions}
 <TransactionRows {transactions} />
{:else}
  <p class="loading">loading...</p>
{/if}

<style>
  h1 {
    font-size: 1.4em;
    font-weight: bold;
    display: block;
  }
</style>
