<script>
  import Modal from '../modal/index.svelte'

  // for a start.. taken from: https://svelte.dev/repl/63eabf4de9ef40108da038cf55cba8dd?version=3.48.0
  export let size = 75
  export let speed = 550
  export let color = '#232d7c'
  export let thickness = 1
  export let gap = 25
  export let radius = 10

  let dash
  $: dash = (2 * Math.PI * radius * (100 - gap)) / 100
</script>

<Modal flyContent={false} coverCol="rgba(255, 255, 255, 0.7)">
  <svg
    height={size}
    width={size}
    style="animation-duration:{speed}ms;"
    class="svelte-spinner"
    viewbox="0 0 32 32"
  >
    <circle
      role="presentation"
      cx="16"
      cy="16"
      r={radius}
      stroke={color}
      fill="none"
      stroke-width={thickness}
      stroke-dasharray="{dash},100"
      stroke-linecap="round"
    />
  </svg>
</Modal>

<style>
  .svelte-spinner {
    transition-property: transform;
    animation-name: svelte-spinner_infinite-spin;
    animation-iteration-count: infinite;
    animation-timing-function: linear;
    position: relative;
    z-index: 1000;
  }
  @keyframes svelte-spinner_infinite-spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }
</style>
