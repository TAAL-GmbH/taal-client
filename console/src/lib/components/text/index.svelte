<script>
  export let testId = null

  export let size = 1
  export let color = '#232D7C'
  export let value = ''
  export let html = false

  let fontSize = '36px'
  let lineHeight = '36px'
  let fontWeight = 700
  let letterSpacing = '-0.02em'

  $: {
    switch (size) {
      case 1:
        fontSize = '36px'
        fontWeight = 700
        break
      case 2:
        fontSize = '24px'
        break
      case 3:
        fontSize = '16px'
        break
      case 4:
        fontSize = '14px'
        fontWeight = 400
        lineHeight = '18px'
        break
      case 5:
        fontSize = '12px'
        fontWeight = 500
        lineHeight = '18px'
        break
    }
  }

  let optProps = {}
  $: {
    if (testId) {
      optProps['data-test-id'] = testId
    }
  }
</script>

<div
  class="tui-text"
  style:--color-local={color}
  style:--font-size-local={fontSize}
  style:--font-weight-local={fontWeight}
  style:--letting-spacing-local={letterSpacing}
  style:--line-height-local={lineHeight}
  {...optProps}
>
  {#if html}
    {@html value}
  {:else}
    {value}
  {/if}
</div>

<style>
  .tui-text {
    font-family: var(--font-family);
    font-size: var(--font-size-local);
    font-weight: var(--font-weight-local);
    letter-spacing: var(--letting-spacing-local);
    line-height: var(--line-height-local);
    color: var(--color-local);
  }

  .tui-text :global(a) {
    color: #0094ff;
  }
</style>
