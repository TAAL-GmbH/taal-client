<script>
  import { tweened } from 'svelte/motion'
  import { cubicOut } from 'svelte/easing'

  const progress = tweened(0, {
    duration: 1000,
    easing: cubicOut,
  })

  // ratio-based progress
  export let ratio = -1
  export let dangerRatio = -1
  export let warnRatio = -1

  // value-based progress
  export let value = -1
  export let total = -1
  export let dangerValue = -1
  export let warnValue = -1

  let showValue = false
  let valueRatio = 0
  let color = '#6EC492'
  let validRatioSet = false

  $: {
    showValue = false
    validRatioSet = ratio !== -1 && ratio >= 0 && ratio <= 1

    if (validRatioSet) {
      showValue = true

      if (showValue) {
        color = '#6EC492'
        if (dangerRatio !== -1 && dangerRatio <= 1 && ratio <= dangerRatio) {
          color = '#FF344C'
        } else if (warnRatio !== -1 && warnRatio <= 1 && ratio <= warnRatio) {
          color = '#F5A200'
        }
      } else {
        color = '#6EC492'
      }
    } else {
      if (value !== -1 && total !== -1 && total !== 0 && value <= total) {
        showValue = true
      }

      if (showValue) {
        valueRatio = value / total
      } else {
        valueRatio = 0
      }

      if (showValue) {
        color = '#6EC492'
        if (
          dangerValue !== -1 &&
          dangerValue <= total &&
          value <= dangerValue
        ) {
          color = '#FF344C'
        } else if (
          warnValue !== -1 &&
          warnValue <= total &&
          value <= warnValue
        ) {
          color = '#F5A200'
        }
      } else {
        color = '#6EC492'
      }
    }

    if (showValue) {
      progress.set(validRatioSet ? ratio : valueRatio)
    } else {
      progress.set(0)
    }
  }
</script>

<div
  class="tui-progress"
  style:--value-width-local="{$progress * 100}%"
  style:--value-color-local={color}
>
  {#if showValue}
    <div class="value" />
  {/if}
</div>

<style>
  .tui-progress {
    background-color: #efefef;
    border-radius: 20px;
    height: 12px;
    width: 100%;
    overflow: hidden;
  }

  .value {
    background-color: var(--value-color-local);
    border-radius: 20px 0px 0px 20px;
    height: 100%;
    width: var(--value-width-local);
  }
</style>
