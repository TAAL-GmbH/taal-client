<script>
  import chartjs, { CategoryScale } from 'chart.js'
  import { onMount } from 'svelte'

  export let yAxisLabel = ''
  export let datasetLabel = ''

  export let yValues = []
  export let xValues = []
  export let divisionFactor = 1
  export let timeUnit = 'day'

  let ctx
  let chart
  let chartCanvas

  const options = {
    responsive: true,
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
              return label / divisionFactor
            },
          },
          scaleLabel: {
            display: true,
            labelString: yAxisLabel,
            fontSize: fontSize,
          },
        },
      ],
      xAxes: [
        {
          type: 'time',
          distribution: 'linear',
          time: {
            unit: timeUnit,
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
      type: 'line',
      options: options,
      data: {
        datasets: [
          {
            label: datasetLabel,
            borderColor: 'rgb(179, 179, 179)',
            backgroundColor: 'rgb(51, 204, 255)',
            lineTension: '0.3',
            fill: true,
          },
        ],
      },
    }
    ctx = chartCanvas.getContext('2d')
    chart = new chartjs(ctx, config)
  })

  function OnTransactionsChange(yValues, xValues, yAxisLabel, timeUnit) {
    if (chart == null) {
      return
    }
    chart.data.labels = xValues
    chart.data.datasets[0].data = yValues
    chart.options.scales.yAxes[0].scaleLabel.labelString = yAxisLabel
    chart.options.scales.xAxes[0].time.unit = timeUnit
    chart.update()
  }

  $: OnTransactionsChange(yValues, xValues, yAxisLabel, timeUnit)
</script>

<canvas
  width="200"
  height="40"
  bind:this={chartCanvas}
  id="transactionsChart"
/>
