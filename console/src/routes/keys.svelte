<script>
  import { onMount } from 'svelte'
  import Key from './key.svelte'
  import Notifications from 'svelte-notifications'

  let keys

  onMount(async () => {
    const r = await fetch(`${BASE_URL}/api/v1/apikeys`)
    const data = await r.json()
    keys = data.keys
  })
</script>

<h1>Active API Keys</h1>

{#if keys}
  <table class="table">
    <tr>
      <th>Created</th>
      <th>API Key</th>
      <th>Address</th>
      <th />
    </tr>
    {#each keys as key}
      <Notifications>
        <Key {key} />
      </Notifications>
    {/each}
  </table>
{:else}
  <p class="loading">loading...</p>
{/if}

<style>
  h1 {
    font-size: 1.4em;
    font-weight: bold;
    /* margin: 0; */
    display: block;
  }
</style>
