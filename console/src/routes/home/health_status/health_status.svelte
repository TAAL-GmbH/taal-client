<script>
  export let status = ''
  export let id
  import Unhealthy from './unhealthy.svelte'
  import Healthy from './healthy.svelte'
  import Pending from './pending.svelte'
</script>

<div class="field">
  {#if status == 'healthy'}
    <Healthy id={id} statusText="Healthy" />
  {:else if status == 'unhealthy'}
    <Unhealthy id={id} statusText="Unhealthy" />
  {:else}
    <Pending id={id} statusText="Pending" />
  {/if}
</div>
