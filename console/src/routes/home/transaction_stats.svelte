<script>
  import { TxDataSize } from '../util/format_functions.svelte'
  import TransactionChart from './transaction_chart.svelte'
  import { GetDateFromISODateString } from '../util/format_functions.svelte'
  import TransactionsInfo from '../util/transactions_info.svelte'
  import Notifications from 'svelte-notifications'

  let transactions
  let nrOfTransactions = 0
  let combinedSize = 0
  let dataSizeYLabel = 'Transaction size [B]'

  function countElements(txs, valueLabels) {
    txs.forEach((tx) => valueLabels[GetDateFromISODateString(tx.created_at)]++)
  }

  function formatDataSize(label) {
    var nrOfDigits = label.toString().length
    var unit = ''
    var divisionFactor = 1
    if (nrOfDigits >= 10) {
      // size >= 1 GB
      divisionFactor = 1000000000
      unit = 'GB'
    }
    if (nrOfDigits >= 7) {
      // size >= 1 MB
      divisionFactor = 1000000
      unit = 'MB'
    }
    if (nrOfDigits >= 4) {
      // size >= 1 kB
      divisionFactor = 1000
      unit = 'kB'
    }

    // return label / divisionFactor
    return { label: label / divisionFactor, unit: unit }
  }

  function sumDataSize(txs, valueLabels) {
    txs.forEach(
      (tx) =>
        (valueLabels[GetDateFromISODateString(tx.created_at)] += tx.data_bytes)
    )
  }

  function OnTransactionsChange(transactions) {
    if (transactions == null) {
      return
    }
    combinedSize = 0
    nrOfTransactions = transactions.length
    var dataSizes = transactions.map((tx) => tx.data_bytes)

    if (dataSizes.length == 0) {
      combinedSize = 0
    }

    dataSizes.forEach((element) => {
      combinedSize += element
    })

    var labelUnit = formatDataSize(combinedSize)
    dataSizeYLabel = 'Transaction size ['+labelUnit.unit+']' 
    console.log(dataSizeYLabel)
  }

  $: OnTransactionsChange(transactions)
</script>

<div class="field">
  <Notifications>
    <TransactionsInfo bind:transactions />
  </Notifications>
  <div class="field">
    <h1>Number of transactions: {nrOfTransactions}</h1>
    <TransactionChart
      valueLabel="Nr of transactions"
      valueFunction={countElements}
      datasetLabel="# Tx"
      bind:transactions
    />
  </div>
  <div class="field">
    <h1>Combined data size of transactions: {TxDataSize(combinedSize)}</h1>
    <TransactionChart
      bind:valueLabel={dataSizeYLabel}
      valueFunction={sumDataSize}
      labelFormatFunc={formatDataSize}
      datasetLabel="Tx data size"
      bind:transactions
    />
  </div>
</div>

<style>
  h1 {
    font-size: 1em;
    font-weight: bold;
    display: block;
  }
</style>
