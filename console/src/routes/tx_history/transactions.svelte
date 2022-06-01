<script>
  // svelte-ignore unused-export-let
  export let location

  import { onMount } from 'svelte'
  import TransactionRows from './transaction_rows.svelte'
  import Cards from './transaction_cards.svelte'

  let transactions
  let classListViewButton = 'button is-success is-selected'
  let classGridViewButton = 'button'

  let listViewSelected = true

  let distinctAPIKeys = []

  function unique(value, index, self) {
    return self.indexOf(value) === index
  }

  onMount(async () => {
    await fetch(`${BASE_URL}/api/v1/transactions/info`)
      .then((r) => r.json())
      .then((data) => {
        transactions = data.transactions
        var apiKeys = transactions.map((a) => a.api_key)
        distinctAPIKeys = apiKeys.filter(unique)
      })
  })

  function clickListViewButton(e) {
    classListViewButton = 'button is-success is-selected'
    classGridViewButton = 'button'
    listViewSelected = true
  }

  function clickGridViewButton(e) {
    classListViewButton = 'button'
    classGridViewButton = 'button is-success is-selected'
    listViewSelected = false
  }
</script>

<h1>Transaction History</h1>
<div class="buttons has-addons is-left">
  <button class={classListViewButton} on:click={clickListViewButton}
    >List view</button
  >
  <button class={classGridViewButton} on:click={clickGridViewButton}
    >Grid view</button
  >
</div>

{#if transactions}
  {#if listViewSelected}
    <TransactionRows {transactions} {distinctAPIKeys} />
  {:else}
    <Cards {transactions} {distinctAPIKeys} />
  {/if}
{:else}
  <p class="loading">loading...</p>
  <button class='button is-loading'></button>
{/if}

<style>
  h1 {
    font-size: 1.4em;
    font-weight: bold;
    display: block;
  }
</style>
