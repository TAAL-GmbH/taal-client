<script>
  import AppOld from './AppOld.svelte'
  import AppNew from './AppNew.svelte'

  let old = false
</script>

{#if old}
  <AppOld />
{:else}
  <AppNew />
{/if}
