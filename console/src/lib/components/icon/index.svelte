<script>
  import icons from './svg'

  export let testId = null

  export let name = 'menu'
  export let size = 24
  export let opacity = 1
</script>

<div
  class="tui-icon"
  data-test-id={testId}
  style:--width-local={size + 'px'}
  style:--height-local={size + 'px'}
  style:--opacity-local={opacity}
  on:click
>
  {#if icons[name]}
    {@html icons[name]}
  {/if}
</div>

<style>
  .tui-icon {
    display: flex;
    flex: 0;
    opacity: var(--opacity-local);

    width: var(--width-local);
    height: var(--height-local);
  }
</style>
