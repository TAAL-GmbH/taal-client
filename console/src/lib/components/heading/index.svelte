<script>
  export let size = 1
  export let color = '#232D7C'
  export let value = ''
  export let html = false

  let fontSize = '36px'
  let lineHeight = '36px'
  let fontWeight = 700
  let letterSpacing = '0'

  $: {
    switch (size) {
      case 1:
        fontSize = '36px'
        fontWeight = 700
        break
      case 2:
        fontSize = '24px'
        break
      case 5:
        fontSize = '16px'
        break
    }
  }
</script>

<div
  class="tui-heading"
  style:--color-local={color}
  style:--font-size-local={fontSize}
  style:--font-weight-local={fontWeight}
  style:--letting-spacing-local={letterSpacing}
  style:--line-height-local={lineHeight}
>
  {#if html}
    {@html value}
  {:else}
    {value}
  {/if}
</div>

<style>
  .tui-heading {
    font-family: 'Work Sans';
    font-size: var(--font-size-local);
    font-weight: var(--font-weight-local);
    letter-spacing: var(--letting-spacing-local);
    line-height: var(--line-height-local);
    color: var(--color-local);
  }
</style>
