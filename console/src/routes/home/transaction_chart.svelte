<script>
  import chartjs from 'chart.js'
  import { onMount } from 'svelte'

  export let transactions

  let chartValues = []
  let chartLabels = []

  let initialChartValues
  let initialChartLabels

  let ctx
  let chart
  let chartCanvas
  function unique(value, index, self) {
    return self.indexOf(value) === index
  }

  const options = {
    responsive: true,
    title: {
      display: false,
      text: 'Transaction history',
    },
    scales: {
      yAxes: [
        {
          ticks: {
            suggestedMin: 0,
            suggestedMax: 7,
            fontSize: fontSize,
          },
          scaleLabel: {
            display: true,
            labelString: 'Number of transactions',
            fontSize: fontSize,
          },
        },
      ],
      xAxes: [
        {
          type: 'time',
          distribution: 'linear',
          time: {
            unit: 'day',
          },
          ticks: {
            fontSize: fontSize,
          },
        },
      ],
    },
  }

  const fontSize = 18

  export function GetDateFromISODateString(dateString) {
    var date = new Date(dateString)
    return date.toISOString().split('T')[0]
  }

  onMount(() => {
    var today = new Date()
    var tomorrow = new Date(today)
    tomorrow.setDate(tomorrow.getDate() + 1)

    initialChartLabels = [
      GetDateFromISODateString(today),
      GetDateFromISODateString(tomorrow),
    ]
    initialChartValues = [0, 0]

    const config = {
      type: 'bar',
      options: options,
      data: {
        labels: initialChartLabels,
        datasets: [
          {
            label: 'Transactions per day',
            borderColor: 'rgb(255, 99, 132)',
            lineTension: '0.2',
            data: initialChartValues,
            fill: false,
          },
        ],
      },
    }
    ctx = chartCanvas.getContext('2d')
    chart = new chartjs(ctx, config)
  })

  function OnTransactionsChange(transactions) {
    if (transactions == null) {
      return
    }

    chartLabels = initialChartLabels
    chartValues = initialChartValues

    if (transactions.length > 0) {
      var reverseTxs = []
      var valueLabels = {}
      reverseTxs = transactions.reverse() // Order of time axis is ascending

      chartLabels = reverseTxs
        .map((tx) => GetDateFromISODateString(tx.created_at))
        .filter(unique)

      chartLabels.forEach((element) => {
        valueLabels[element] = 0
      })

      reverseTxs.forEach(
        (tx) => valueLabels[GetDateFromISODateString(tx.created_at)]++
      )
      chartValues = Object.values(valueLabels)
    }

    chart.data.labels = chartLabels
    chart.data.datasets[0].data = chartValues
    chart.update()
  }

  $: OnTransactionsChange(transactions)
</script>

<canvas bind:this={chartCanvas} id="transactionsChart" />
