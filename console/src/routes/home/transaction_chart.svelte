<script>
  import chartjs from 'chart.js'

  let valueLabels = {}
  export let transactions
  let reverseTxs

  let chartValues = []
  let chartLabels = []

  let ctx
  let chart
  let chartCanvas
  function unique(value, index, self) {
    return self.indexOf(value) === index
  }

  export function GetDateFromISODateString(dateString) {
    var date = new Date(dateString)
    return date.toISOString().split('T')[0]
  }

  $: if (transactions.length > 0) {
    reverseTxs = transactions.reverse()
    
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
              suggestedMax: 10,
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
          },
        ],
      },
    }
    console.log(chartLabels)
    console.log(chartValues)

    const config = {
      type: 'line',
      options: options,
      data: {
        labels: chartLabels,
        datasets: [
          {
            label: 'Transactions',
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
