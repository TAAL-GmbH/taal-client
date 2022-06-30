<script>
  import { fade, fly } from 'svelte/transition'

  let w
  let h
  export let maxContentW = '900px'
</script>

<div class="cover" transition:fade={{ duration: 200 }} />

<div
  class="tui-modal"
  in:fly={{ y: -200, opacity: 0, duration: 200 }}
  out:fade
  bind:clientWidth={w}
  bind:clientHeight={h}
  style:--width-local="{w}px"
  style:--height-local="{h}px"
  style:--max-content-width={maxContentW}
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
    background: rgba(40, 41, 51, 0.7);
  }

  .tui-modal {
    position: absolute;
    left: calc(50% - var(--width-local) / 2);
    top: calc(50% - var(--height-local) / 2);

    display: flex;
    flex-direction: column;
    align-items: center;

    width: calc(100% - 40px);
    max-width: var(--max-content-width);
  }
</style>
