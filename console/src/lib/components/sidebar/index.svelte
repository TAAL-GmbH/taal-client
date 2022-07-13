<script>
  import { createEventDispatcher } from 'svelte'
  import { onMount } from 'svelte'
  import { headerHeight } from '../../stores'
  import Button from '../button/index.svelte'

  const dispatch = createEventDispatcher()

  export let links = []
  export let dataKey = 'path'

  let menuWidth = 385
  let menuLeft = -menuWidth

  onMount(() => {
    const id = setTimeout(() => {
      menuLeft = 0
    }, 0)

    return () => {
      clearTimeout(id)
      menuLeft = -menuWidth
    }
  })

  function onLink(item) {
    dispatch('link', item)
  }
</script>

<div
  class="tui-sidebar"
  style:--top-local={$headerHeight + 'px'}
  style:--menu-width-local={menuWidth + 'px'}
  style:--menu-left-local={menuLeft + 'px'}
>
  <div class="cover" />
  <div class="menu">
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
    left: var(--menu-left-local);
    top: var(--top-local);
    bottom: 0;
    transition: left 0.2s ease-in-out;

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
