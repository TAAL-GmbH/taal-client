<script>
  export let size = 1
  export let color = '#232D7C'

  let fontSize = '36px'
  let fontWeight = 700
  let letterSpacing = '-0.02em'

  $: {
    switch (size) {
      case 1:
        fontSize = '36px'
        fontWeight = 700
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
>
  <slot />
</div>

<style>
  .tui-heading {
    font-family: 'Work Sans';
    font-size: var(--font-size-local);
    font-weight: var(--font-weight-local);
    letter-spacing: var(--letting-spacing-local);
    color: var(--color-local);
  }
</style>
