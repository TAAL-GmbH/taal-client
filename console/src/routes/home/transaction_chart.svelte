<script>
  import chartjs from 'chart.js'
  import { onMount } from 'svelte'
  import {GetDateFromISODateString} from '../util/format_functions.svelte'

  export let transactions
  export let valueFunction = () =>{}
  export let valueLabel = ''
  export let datasetLabel = ''

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
            labelString: valueLabel,
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
            label: datasetLabel,
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

      valueFunction(reverseTxs, valueLabels)
      chartValues = Object.values(valueLabels)
    }

    chart.data.labels = chartLabels
    chart.data.datasets[0].data = chartValues
    chart.update()
  }

  $: OnTransactionsChange(transactions)
</script>

<canvas width="200" height="40" bind:this={chartCanvas} id="transactionsChart" />
