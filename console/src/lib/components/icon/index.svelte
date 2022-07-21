<script>
  import icons from './svg'

  export let testId = null

  export let name = 'menu'
  export let size = 24

  let optProps = {}
  $: {
    if (testId) {
      optProps['data-test-id'] = testId
    }
  }
</script>

<div
  class="tui-icon"
  style:--width-local={size + 'px'}
  style:--height-local={size + 'px'}
  on:click
  {...optProps}
>
  {#if icons[name]}
    {@html icons[name]}
  {/if}
</div>

<style>
  .tui-icon {
    display: flex;
    flex: 0;

    width: var(--width-local);
    height: var(--height-local);
  }
</style>
