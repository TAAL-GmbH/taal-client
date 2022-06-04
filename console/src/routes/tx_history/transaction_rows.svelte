<script>
import TransactionRow from './transaction_row.svelte'
export let transactions
export let distinctAPIKeys

</script>

{#if transactions}
  <table class="table is-fullwidth">
    <thead>
      <tr>
        <th>Created at (UTC)</th>
        <th>ID</th>
        <th>API key used</th>
        <th>Data size</th>
        <th>Filename</th>
        <th></th>
      </tr>
    </thead>
    <tbody>
      {#each transactions as transaction}
        <TransactionRow transaction={transaction} distinctAPIKeys={distinctAPIKeys} />
      {/each}
    </tbody>
  </table>
{:else}
  <p class="loading">loading...</p>
{/if}
