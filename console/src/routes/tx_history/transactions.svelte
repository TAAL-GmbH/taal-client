<script>
  // svelte-ignore unused-export-let
  export let location

  import { onMount } from 'svelte'
  import Transaction from './transaction.svelte'

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
  <table class="table">
    <tr>
      <th>Created At</th>
      <th>ID</th>
      <th>API Key used</th>
      <th>Data size [bytes]</th>
      <th>Filename</th>
    </tr>
    {#each transactions as transaction}
      <Transaction {transaction} />
    {/each}
  </table>
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
