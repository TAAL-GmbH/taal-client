<script>
  import chartjs, { CategoryScale } from 'chart.js'
  import { onMount } from 'svelte'

  export let labelFormatCallbackFunc = (label) => {
    return label
  }
  export let yAxisLabel = ''
  export let datasetLabel = ''

  export let yValues = []
  export let xValues = []

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
            callback: labelFormatCallbackFunc,
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
            backgroundColor: 'rgb(204, 255, 153)',
            fill: false,
          },
        ],
      },
    }
    ctx = chartCanvas.getContext('2d')
    chart = new chartjs(ctx, config)
  })

  function OnTransactionsChange(yValues, xValues, yAxisLabel) {
    if (chart == null) {
      return
    }
    chart.data.labels = xValues
    chart.data.datasets[0].data = yValues
    chart.options.scales.yAxes[0].scaleLabel.labelString = yAxisLabel
    chart.update()
  }

  $: OnTransactionsChange(yValues, xValues, yAxisLabel)
</script>

<canvas
  width="200"
  height="40"
  bind:this={chartCanvas}
  id="transactionsChart"
/>
