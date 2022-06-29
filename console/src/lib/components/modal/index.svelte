<script>
  import { createEventDispatcher } from 'svelte'
  import { clickOutside } from '../../actions'

  const dispatch = createEventDispatcher()

  let w
  let h

  function onClose() {
    dispatch('close')
  }
</script>

<div class="cover" />

<div
  class="tui-modal"
  bind:clientWidth={w}
  bind:clientHeight={h}
  use:clickOutside
  on:outclick={onClose}
  style:--width-local="{w}px"
  style:--height-local="{h}px"
>
  <slot />
</div>

<style>
  .cover {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    background: rgba(0, 0, 0, 0.7);
  }

  .box {
    position: absolute;
    /* width: var(--width);
    height: var(--height); */
    left: calc(50% - var(--width-local) / 2);
    top: calc(50% - var(--height-local) / 2);

    /* display: flex;
    align-items: center;
    padding: 8px;
    border-radius: 4px;
    background-color: #ff3e00;
    color: #fff;
    text-align: center;
    font-weight: bold; */

    width: calc(100% - 40px);
    max-width: 880px;
  }
</style>
