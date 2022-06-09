<script>
  import chartjs, { CategoryScale } from 'chart.js'
  import { onMount } from 'svelte'

  export let labelFormatFunc = (label) => {
    return { label: label, unit: '' }
  }
  export let valueLabel = ''
  export let datasetLabel = ''

  export let chartValues = []
  export let chartLabels = []

  let unit = ''
  let ctx
  let chart
  let chartCanvas

  const options = {
    responsive: true,
    title: {
      display: false,
      text: 'Transaction history',
    },
    scales: {
      yAxes: [
        {
          afterUpdate: function (axis) {
            axis.width = 100
          },

          ticks: {
            suggestedMin: 0,
            suggestedMax: 7,
            fontSize: fontSize,
            callback: function (label, index, labels) {
              var labelUnit = labelFormatFunc(label)
              unit = labelUnit.unit
              return labelUnit.label
            },
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
    const config = {
      type: 'bar',
      options: options,
      data: {
        datasets: [
          {
            label: datasetLabel,
            borderColor: 'rgb(255, 99, 132)',
            lineTension: '0.2',
            fill: false,
          },
        ],
      },
    }
    ctx = chartCanvas.getContext('2d')
    chart = new chartjs(ctx, config)
  })

  function OnTransactionsChange(chartValues, chartLabels, valueLabel) {
    if (chart == null) {
      return
    }
    chart.data.labels = chartLabels
    chart.data.datasets[0].data = chartValues
    chart.options.scales.yAxes[0].scaleLabel.labelString = valueLabel
    chart.update()
  }

  $: OnTransactionsChange(chartValues, chartLabels, valueLabel)
</script>

<canvas
  width="200"
  height="40"
  bind:this={chartCanvas}
  id="transactionsChart"
/>
