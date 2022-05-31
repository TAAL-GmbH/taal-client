<script>
  import Card from './transaction_card.svelte'
  export let transactions
</script>

{#if transactions}
<div class="columns is-multiline">

  {#each transactions as transaction}
    <Card {transaction} />
  {/each}
</div>
{:else}
  <p class="loading">loading...</p>
{/if}
