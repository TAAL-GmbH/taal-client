<script>
  import { tweened } from 'svelte/motion'
  import { cubicOut } from 'svelte/easing'

  const progress = tweened(0, {
    duration: 1000,
    easing: cubicOut,
  })

  export let value = -1
  export let total = -1
  export let dangerValue = -1
  export let warnValue = -1

  let showValue = false
  let ratio = 0
  let color = '#6EC492'

  $: {
    if (value !== -1 && total !== -1 && total !== 0 && value <= total) {
      showValue = true
    }

    if (showValue) {
      ratio = value / total
    } else {
      ratio = 0
    }

    if (showValue) {
      color = '#6EC492'
      if (dangerValue !== -1 && dangerValue <= total && value <= dangerValue) {
        color = '#FF344C'
      } else if (warnValue !== -1 && warnValue <= total && value <= warnValue) {
        color = '#F5A200'
      }
    } else {
      color = '#6EC492'
    }

    progress.set(ratio)
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
    overflow: hidden;
  }

  .value {
    background-color: var(--value-color-local);
    border-radius: 20px 0px 0px 20px;
    height: 100%;
    width: var(--value-width-local);
  }
</style>
