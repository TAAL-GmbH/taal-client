<script>
  import chartjs from 'chart.js'
  import { GetFormatedTxDate } from '../tx_history/transaction_format_functions.svelte'

  let valueLabels = {}
  export let transactions

  let chartValues = []
  let chartLabels = []

  let ctx
  let chart
  let chartCanvas
  function unique(value, index, self) {
    return self.indexOf(value) === index
  }

  $: if (transactions.length > 0) {
    chartLabels = transactions
      .reverse()
      .map((tx) => GetFormatedTxDate(tx.created_at))
      .filter(unique)

    chartLabels.forEach((element) => {
      valueLabels[element] = 0
    })

    transactions.forEach(
      (tx) => valueLabels[GetFormatedTxDate(tx.created_at)]++
    )
    chartValues = Object.values(valueLabels)

    const config = {
      type: 'line',
      data: {
        labels: chartLabels,
        datasets: [
          {
            label: 'Transactions',
            // backgroundColor: 'rgb(255, 99, 132)',
            borderColor: 'rgb(255, 99, 132)',
            data: chartValues,
          },
        ],
      },
    }

    ctx = chartCanvas.getContext('2d')
    chart = new chartjs(ctx, config)
  }
</script>

<canvas bind:this={chartCanvas} id="transactionsChart" />
