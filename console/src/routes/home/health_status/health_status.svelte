<script>
  export let status = ''
  export let id
  import Unhealthy from './unhealthy.svelte'
  import Healthy from './healthy.svelte'
  import Loading from './loading.svelte'
</script>

<div class="field">
  {#if status == 'healthy'}
    <Healthy id={id} statusText="Healthy" />
  {:else if status == 'unhealthy'}
    <Unhealthy id={id} statusText="Unhealthy" />
  {:else}
    <Loading id={id}/>
  {/if}
</div>
