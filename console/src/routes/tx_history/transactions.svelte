<script>
  // svelte-ignore unused-export-let
  export let location

  import TransactionRows from './transaction_rows.svelte'
  import Cards from './transaction_cards.svelte'
  import TransactionsInfo from '../util/transactions_info.svelte'
  import { setButtonClassIsSuccess } from '../util/control_functions.svelte'

  let transactions
  let distinctAPIKeys = []

  let viewButtonClasses = ['button is-success is-selected', 'button']
  let listViewSelected = true

  function unique(value, index, self) {
    return self.indexOf(value) === index
  }

  function clickListViewButton(e) {
    listViewSelected = true
    viewButtonClasses = setButtonClassIsSuccess(0, viewButtonClasses)
  }

  function clickGridViewButton(e) {
    listViewSelected = false
    viewButtonClasses = setButtonClassIsSuccess(1, viewButtonClasses)
  }

  function onChangeTransactions(transactions) {
    if (transactions == null) {
      return
    }
    var apiKeys = transactions.map((a) => a.api_key)
    distinctAPIKeys = apiKeys.filter(unique)
  }

  $: onChangeTransactions(transactions)
</script>

<h1>Transaction History</h1>
<div class="buttons has-addons is-left">
  <button class={viewButtonClasses[0]} on:click={clickListViewButton}
    >List view</button
  >
  <button class={viewButtonClasses[1]} on:click={clickGridViewButton}
    >Grid view</button
  >
</div>
<TransactionsInfo bind:transactions />

{#if transactions}
  {#if listViewSelected}
    <TransactionRows {transactions} {distinctAPIKeys} />
  {:else}
    <Cards {transactions} {distinctAPIKeys} />
  {/if}
{:else}
  <p class="loading">loading...</p>
  <button class="button is-loading" />
{/if}

<style>
  h1 {
    font-size: 1.4em;
    font-weight: bold;
    display: block;
  }

  button {
    min-width: 100px;
  }
</style>
