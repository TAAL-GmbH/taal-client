<script>
    import { onMount } from "svelte";
    import Transaction from "./transaction.svelte";
    
    let transactions;

    onMount(async () => {
    await fetch(`${BASE_URL}/api/v1/transactions`)
      .then(r => r.json())
      .then(data => {
        transactions = data.transactions;
      });
  })

</script>

<h1>Transactions</h1>

{#if transactions}
  {#each transactions as transaction }
    <ul>
      <li>    
        <Transaction {transaction} />
      </li>
    </ul>
  {/each}
{:else}
  <p class="loading">loading...</p>
{/if}
