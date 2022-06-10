<script>
  import { TxDataSize } from '../util/format_functions.svelte'
  import TransactionChart from './transaction_chart.svelte'
  import TransactionsInfo from '../util/transactions_info.svelte'
  import Notifications from 'svelte-notifications'
  import Health from './health.svelte'
  import { onMount } from 'svelte'
  import Fa from 'svelte-fa'
  import { faSpinner } from '@fortawesome/free-solid-svg-icons'

  let transactionInfos
  let statNrOfTransactions = 0
  let statCombinedSize = 0
  let dataSizeYLabel = 'Transaction size [B]'

  let dataSizeDivisionFactor = 1
  let timeUnit = 'day'

  let xValuesInfos
  let yValuesInfosNrOfTxs
  let yValuesInfosDataSize

  onMount(() => {
    var today = new Date()
    var tomorrow = new Date(today)
    tomorrow.setDate(tomorrow.getDate() + 1)
  })

  function formatDataSizeDivFactor(dataSizeValue) {
    var nrOfDigits = dataSizeValue.toString().length
    if (nrOfDigits >= 10) {
      // size >= 1 GB
      return { divFactor: 1000000000, unit: 'GB' }
    }
    if (nrOfDigits >= 7) {
      // size >= 1 MB
      return { divFactor: 1000000, unit: 'MB' }
    }
    if (nrOfDigits >= 4) {
      // size >= 1 kB
      return { divFactor: 1000, unit: 'kB' }
    }

    return { divFactor: 1, unit: 'B' }
  }

  function OnTransactionsChange(transactionInfos) {
    if (transactionInfos == null) {
      return
    }

    statNrOfTransactions = transactionInfos.length

    statCombinedSize = transactionInfos
      .map((tx) => tx.data_bytes)
      .reduce((partialSum, a) => partialSum + a, 0)

    xValuesInfos = transactionInfos.map((item) => item.timestamp)
    yValuesInfosNrOfTxs = transactionInfos.map((item) => item.count)
    yValuesInfosDataSize = transactionInfos.map((item) => item.data_bytes)

    dataSizeYLabel =
      'Transaction size [' +
      formatDataSizeDivFactor(Math.max(...yValuesInfosDataSize)).unit +
      ']'

    dataSizeDivisionFactor = formatDataSizeDivFactor(
      Math.max(...yValuesInfosDataSize)
    ).divFactor
  }

  $: OnTransactionsChange(transactionInfos)
</script>

<div class="field">
  <div class="columns">
    <div class="column is-four-fifths">
      <Notifications>
        <TransactionsInfo bind:transactionInfos bind:timeUnit />
      </Notifications>
    </div>
    <div class="column ">
      <Health />
    </div>
  </div>
  {#if transactionInfos}
    <div class="field">
      <h1>Number of transactions: {statNrOfTransactions}</h1>
      <TransactionChart
        yAxisLabel="Nr of transactions"
        bind:xValues={xValuesInfos}
        bind:yValues={yValuesInfosNrOfTxs}
        bind:timeUnit
        datasetLabel="# Tx"
      />
    </div>
  {/if}
  {#if transactionInfos}
    <div class="field">
      <h1>
        Combined data size of transactions: {TxDataSize(statCombinedSize)}
      </h1>
      <TransactionChart
        bind:yAxisLabel={dataSizeYLabel}
        bind:xValues={xValuesInfos}
        bind:yValues={yValuesInfosDataSize}
        bind:divisionFactor={dataSizeDivisionFactor}
        bind:timeUnit
        datasetLabel="Tx data size"
      />
    </div>
  {:else}
    <div id="spinner" class="field">
      <Fa icon={faSpinner} size="3x" pulse />
    </div>
  {/if}
</div>

<style>
  #spinner {
    padding-top: 100px;
    margin: auto;
    width: 0%;
  }
  h1 {
    font-size: 1em;
    font-weight: bold;
    display: block;
  }
</style>
