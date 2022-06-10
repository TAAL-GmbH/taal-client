<script>
  // svelte-ignore unused-export-let
  export let location

  import TransactionRows from './transaction_rows.svelte'
  import Cards from './transaction_cards.svelte'
  import Transactions from '../util/transactions.svelte'
  import { setButtonClassIsSuccess } from '../util/control_functions.svelte'
  import Notifications from 'svelte-notifications'
  import Fa from 'svelte-fa'
  import { faList, faTh } from '@fortawesome/free-solid-svg-icons'

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

<div class="columns">
  <div class="column is-four-fifths">
    <h1>Transaction History</h1>
    <Notifications>
      <Transactions bind:transactions />
    </Notifications>
  </div>
  <div class="column">
    <div id="viewButtons" class="buttons has-addons is-left">
      <button class={viewButtonClasses[0]} on:click={clickListViewButton}
        ><Fa icon={faList} />&nbsp; List view</button
      >
      <button class={viewButtonClasses[1]} on:click={clickGridViewButton}
        ><Fa icon={faTh} />&nbsp; Grid view</button
      >
    </div>
  </div>
</div>
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

  #viewButtons {
    margin-top: 40px;
  }

  button {
    min-width: 100px;
  }
</style>
