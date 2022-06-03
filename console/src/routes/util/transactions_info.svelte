<script>
  export let transactions = []

  let filterButtonClasses = [
    'button is-success is-selected',
    'button',
    'button',
  ]

  const hours24h = 24
  const hours7d = 168
  const hours30d = 720

  let hoursBack = hours24h

  function setButtonClassIsSuccess(setIndex, buttonArray) {
    for (let index = 0; index < buttonArray.length; index++) {
      if (index == setIndex) {
        buttonArray[index] = 'button is-success'
        continue
      }
      buttonArray[index] = 'button'
    }
    return buttonArray
  }

  function click24Hours(e) {
    hoursBack = hours24h
  }
  function click7Days(e) {
    hoursBack = hours7d
  }
  function click30Days(e) {
    hoursBack = hours30d
  }
  function onChangeHoursBack(hoursBack) {
    switch (hoursBack) {
      case hours24h:
        filterButtonClasses = setButtonClassIsSuccess(0, filterButtonClasses)
        break
      case hours7d:
        filterButtonClasses = setButtonClassIsSuccess(1, filterButtonClasses)
        break
      case hours30d:
        filterButtonClasses = setButtonClassIsSuccess(2, filterButtonClasses)
      default:
        break
    }

    fetch(`${BASE_URL}/api/v1/transactions/info?hours_back=${hoursBack}`)
      .then((r) => r.json())
      .then((data) => {
        transactions = data.transactions
      })
  }
  $: onChangeHoursBack(hoursBack)
</script>

<div class="field has-addons">
  <p class="control">
    <button class={filterButtonClasses[0]} on:click={click24Hours}
      >24 hours</button
    >
  </p>
  <p class="control">
    <button class={filterButtonClasses[1]} on:click={click7Days}>
      7 days</button
    >
  </p>
  <p class="control">
    <button class={filterButtonClasses[2]} on:click={click30Days}
      >30 days</button
    >
  </p>
</div>

<style>
  button {
    min-width: 100px;
  }
</style>
