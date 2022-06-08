<script>
  import Card from './transaction_card.svelte'
  export let transactions
  export let distinctAPIKeys
</script>

{#if transactions}
<div class="columns is-multiline">

  {#each transactions as transaction}
    <Card transaction={transaction} distinctAPIKeys={distinctAPIKeys} />
  {/each}
</div>
{:else}
  <p class="loading">loading...</p>
{/if}
