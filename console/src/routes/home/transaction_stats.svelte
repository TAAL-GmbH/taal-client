<script>
  import { TxDataSize } from '../util/format_functions.svelte'
  import TransactionChart from './transaction_chart.svelte'
  import { GetDateFromISODateString } from '../util/format_functions.svelte'
  import TransactionsInfo from '../util/transactions_info.svelte'
  import Notifications from 'svelte-notifications'

  let transactions
  let nrOfTransactions = 0
  let combinedSize = 0

  function countElements(txs, valueLabels) {
    txs.forEach((tx) => valueLabels[GetDateFromISODateString(tx.created_at)]++)
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
    nrOfTransactions = transactions.length
    var dataSizes = transactions.map((tx) => tx.data_bytes)

    if (dataSizes.length == 0) {
      combinedSize = 0
    }

    dataSizes.forEach((element) => {
      combinedSize += element
    })
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
      valueLabel="Number of transactions in time"
      valueFunction={countElements}
      datasetLabel="# Tx"
      bind:transactions
    />
  </div>
  <div class="field">
    <h1>Combined data size of transactions: {TxDataSize(combinedSize)}</h1>
    <TransactionChart
      valueLabel="Transaction size [byte]"
      valueFunction={sumDataSize}
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
