<script>
import TransactionRow from './transaction_row.svelte'
import DownloadButton from './download_button.svelte'
export let transactions

</script>

{#if transactions}
  <table class="table">
    <thead>
      <tr>
        <th>Created At (UTC)</th>
        <th>ID</th>
        <th>API Key used</th>
        <th>Data size [bytes]</th>
        <th>Filename</th>
      </tr>
    </thead>
    <tbody>
      {#each transactions as transaction}
        <TransactionRow {transaction} />
      {/each}
    </tbody>
  </table>
{:else}
  <p class="loading">loading...</p>
{/if}