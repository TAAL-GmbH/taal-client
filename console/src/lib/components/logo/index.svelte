<script>
  import logos from './svg'

  export let testId = null

  export let name = 'taal-blue'
  export let width = -1
  export let height = -1

  let hasNoSize = true

  $: {
    hasNoSize = width === -1 && height === -1
  }

  let optProps = {}
  $: {
    if (testId) {
      optProps['data-test-id'] = testId
    }
  }
</script>

<div
  class="tui-logo"
  class:w={width !== -1}
  class:h={height !== -1 || hasNoSize}
  style:--width-local={width + 'px'}
  style:--height-local={hasNoSize ? '64px' : height + 'px'}
  on:click
  {...optProps}
>
  {#if logos[name]}
    {@html logos[name]}
  {/if}
</div>

<style>
  .tui-logo {
    display: flex;
    flex: 0 0 auto;
  }

  .tui-logo.w {
    width: var(--width-local);
  }

  .tui-logo.h {
    height: var(--height-local);
  }
</style>
