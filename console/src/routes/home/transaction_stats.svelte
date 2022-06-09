<script>
  import { TxDataSize } from '../util/format_functions.svelte'
  import TransactionChart from './transaction_chart.svelte'
  import { GetDateFromISODateString } from '../util/format_functions.svelte'
  import TransactionsInfo from '../util/transactions_info.svelte'
  import Notifications from 'svelte-notifications'
  import { onMount } from 'svelte'

  let transactions
  let statNrOfTransactions = 0
  let statCombinedSize = 0
  let dataSizeYLabel = 'Transaction size [B]'
  
  let dataSizeDivisionFactor = 1

  let initialChartValues
  let initialChartLabels

  let xValues
  let yValuesNrOfTxs
  let yValuesDataSize

  onMount(() => {
    var today = new Date()
    var tomorrow = new Date(today)
    tomorrow.setDate(tomorrow.getDate() + 1)

    initialChartLabels = [
      GetDateFromISODateString(today),
      GetDateFromISODateString(tomorrow),
    ]
    initialChartValues = [0, 0]
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

  function unique(value, index, self) {
    return self.indexOf(value) === index
  }

  function sumDataSize(txs, valueLabels) {
    txs.forEach(
      (tx) =>
        (valueLabels[GetDateFromISODateString(tx.created_at)] += tx.data_bytes)
    )
  }

  function countElements(txs, valueLabels) {
    txs.forEach((tx) => valueLabels[GetDateFromISODateString(tx.created_at)]++)
  }

  function getYValuesGroupedByXValues(txs, xValues, func) {
    var valueLabels = {}

    xValues.forEach((element) => {
      valueLabels[element] = 0
    })

    func(txs, valueLabels)
    return Object.values(valueLabels)
  }

  function OnTransactionsChange(transactions) {
    if (transactions == null) {
      return
    }

    statNrOfTransactions = transactions.length

    statCombinedSize = 0
    var dataSizes = transactions.map((tx) => tx.data_bytes)

    if (dataSizes.length == 0) {
      statCombinedSize = 0
    }

    dataSizes.forEach((element) => {
      statCombinedSize += element
    })

    xValues = initialChartLabels
    yValuesDataSize = initialChartValues
    yValuesNrOfTxs = initialChartValues

    if (transactions.length > 0) {
      var reverseTxs = []
      reverseTxs = transactions.reverse() // Order of time axis is ascending
      xValues = reverseTxs
        .map((tx) => GetDateFromISODateString(tx.created_at))
        .filter(unique)

      yValuesNrOfTxs = getYValuesGroupedByXValues(
        reverseTxs,
        xValues,
        countElements
      )
      yValuesDataSize = getYValuesGroupedByXValues(
        reverseTxs,
        xValues,
        sumDataSize
      )
      yValuesNrOfTxs = getYValuesGroupedByXValues(reverseTxs, xValues, countElements)
      yValuesDataSize = getYValuesGroupedByXValues(reverseTxs, xValues, sumDataSize)
    }

    dataSizeYLabel = 'Transaction size [' + formatDataSizeDivFactor(Math.max(...yValuesDataSize)).unit + ']'
    
    dataSizeDivisionFactor = formatDataSizeDivFactor(Math.max(...yValuesDataSize)).divFactor
  }

  $: OnTransactionsChange(transactions)
</script>

<div class="field">
  <Notifications>
    <TransactionsInfo bind:transactions />
  </Notifications>
  <div class="field">
    <h1>Number of transactions: {statNrOfTransactions}</h1>
    <TransactionChart
      yAxisLabel="Nr of transactions"
      bind:xValues
      bind:yValues={yValuesNrOfTxs}
      datasetLabel="# Tx"
    />
  </div>
  <div class="field">
    <h1>Combined data size of transactions: {TxDataSize(statCombinedSize)}</h1>
    <TransactionChart
      bind:yAxisLabel={dataSizeYLabel}
      bind:xValues
      bind:yValues={yValuesDataSize}
      bind:divicionFactor={dataSizeDivisionFactor}
      datasetLabel="Tx data size"
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
