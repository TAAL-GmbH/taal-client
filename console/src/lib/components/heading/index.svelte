<script>
  import { mediaSize } from '../../stores/index'

  export let testId = null

  export let size = 1
  export let color = '#232D7C'
  export let value = ''
  export let html = false

  let fontSize = 36
  let lineHeight = 36
  let fontWeight = 700
  let letterSpacing = '0'

  $: mediaSmall = $mediaSize === 'small'

  $: {
    switch (size) {
      case 1:
        fontSize = mediaSmall ? 24 : 36
        fontWeight = 700
        break
      case 2:
        fontSize = mediaSmall ? 18 : 24
        break
      case 5:
        fontSize = 16
        break
    }
  }
</script>

<div
  class="tui-heading"
  data-test-id={testId}
  style:--color-local={color}
  style:--font-size-local={fontSize + 'px'}
  style:--font-weight-local={fontWeight}
  style:--letting-spacing-local={letterSpacing}
  style:--line-height-local={lineHeight + 'px'}
>
  {#if html}
    {@html value}
  {:else}
    {value}
  {/if}
</div>

<style>
  .tui-heading {
    font-family: var(--font-family);

    font-size: var(--font-size-local);
    font-weight: var(--font-weight-local);
    letter-spacing: var(--letting-spacing-local);
    line-height: var(--line-height-local);
    color: var(--color-local);
  }
</style>
