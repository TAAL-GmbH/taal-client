<script>
  export let value = ''
  export let className = ''

  let props = {}

  $: {
    if (className) {
      props.class = className
    }
  }
</script>

{#if value}
  <span {...props}>{value}</span>
{:else}
  ''
{/if}
