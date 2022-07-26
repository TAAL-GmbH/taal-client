<script>
  import { createEventDispatcher } from 'svelte'
  import { fade, fly } from 'svelte/transition'
  import { headerHeight } from '../../stores'
  import Button from '../button/index.svelte'

  const dispatch = createEventDispatcher()

  export let testId = null

  export let links = []
  export let dataKey = 'path'
  export let menuWidth = 385

  function onLink(item) {
    dispatch('link', item)
  }
  function onClose() {
    dispatch('close')
  }
</script>

<div
  class="tui-sidebar"
  data-test-id={testId}
  style:--top-local={$headerHeight + 'px'}
  style:--menu-width-local={menuWidth + 'px'}
>
  <div class="cover" in:fade on:mousedown|preventDefault={onClose} />
  <div class="menu" in:fly={{ x: -menuWidth, duration: 200 }}>
    <div class="content">
      {#each links as link (link[dataKey])}
        <div class="item">
          <Button
            variant="link"
            size="small"
            selected={link.selected}
            on:click={(e) => onLink(link)}>{link.label}</Button
          >
        </div>
      {/each}
    </div>
  </div>
</div>

<style>
  .tui-sidebar {
    font-family: var(--font-family);
    box-sizing: var(--box-sizing);
  }

  .cover {
    width: 100%;
    height: 100%;
    position: fixed;
    left: 0;
    top: var(--top-local);
    background-color: rgba(40, 41, 51, 0.7);
  }

  .menu {
    position: absolute;
    width: var(--menu-width-local);
    top: var(--top-local);
    left: 0;
    bottom: 0;

    background: white;
  }

  .content {
    margin: 32px 22px;
    overflow-y: auto;
    height: calc(100% - 64px);

    display: flex;
    flex-direction: column;
    gap: 24px;
  }

  .item {
    flex-grow: 0;
    align-self: flex-start;
  }
</style>
