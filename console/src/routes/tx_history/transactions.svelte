<script>
  // svelte-ignore unused-export-let
  export let location

  import { onMount } from 'svelte'
  import TransactionRows from './transaction_rows.svelte'
  import Cards from './transaction_cards.svelte'

  let transactions
  let distinctAPIKeys = []

  let viewButtonClasses = ['button is-success is-selected', 'button']
  let listViewSelected = true

  let filterButtonClasses = [
    'button is-success is-selected',
    'button',
    'button',
  ]

  const hours24h = 24
  const hours7d = 168
  const hours30d = 720

  let hoursBack = hours24h

  function setButtonClassIsSuccess(setIndex, buttonArray) {
    for (let index = 0; index < buttonArray.length; index++) {
      if (index == setIndex) {
        buttonArray[index] = 'button is-success'
        continue
      }
      buttonArray[index] = 'button'
    }
    console.log(buttonArray)
    return buttonArray
  }

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
  function click24Hours(e) {
    hoursBack = hours24h
  }
  function click7Days(e) {
    hoursBack = hours7d
  }
  function click30Days(e) {
    hoursBack = hours30d
  }

  function onChangeHoursBack(hoursBack) {
    console.log(hoursBack)
    switch (hoursBack) {
      case hours24h:
        filterButtonClasses = setButtonClassIsSuccess(0, filterButtonClasses)
        break
      case hours7d:
        filterButtonClasses = setButtonClassIsSuccess(1, filterButtonClasses)
        break
      case hours30d:
        filterButtonClasses = setButtonClassIsSuccess(2, filterButtonClasses)
      default:
        break
    }

    fetch(`${BASE_URL}/api/v1/transactions/info?hours_back=${hoursBack}`)
      .then((r) => r.json())
      .then((data) => {
        transactions = data.transactions
        var apiKeys = transactions.map((a) => a.api_key)
        distinctAPIKeys = apiKeys.filter(unique)
      })
  }

  $: onChangeHoursBack(hoursBack)
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
<div class="field has-addons">
  <p class="control">
    <button class={filterButtonClasses[0]} on:click={click24Hours}
      >24 hours</button
    >
  </p>
  <p class="control">
    <button class={filterButtonClasses[1]} on:click={click7Days}>
      7 days</button
    >
  </p>
  <p class="control">
    <button class={filterButtonClasses[2]} on:click={click30Days}
      >30 days</button
    >
  </p>
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

  button {
    min-width: 100px;
  }
</style>
