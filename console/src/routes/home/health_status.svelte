<script>
  export let status = ''
  export let id
</script>

<div class="field">
  {#if status == 'healthy'}
    <button {id} class="button is-success">Healthy</button>
  {:else if status == 'unhealthy'}
    <button {id} class="button is-danger">Unhealthy</button>
  {:else}
    <button {id} class="button is-loading" />
  {/if}
</div>

<style>
  button {
    min-width: 89px;
  }
</style>
