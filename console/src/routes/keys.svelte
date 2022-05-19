<script>
    import { onMount } from "svelte";
    import Key from "./key.svelte";
    
    let keys;

    onMount(async () => {
    await fetch(`${BASE_URL}/api/v1/apikeys`)
      .then(r => r.json())
      .then(data => {
        keys = data.keys;
      });
  })

</script>

<h1>API Keys</h1>

{#if keys}
  {#each keys as key }
    <ul>
      <li>    
        <Key {key} />
      </li>
    </ul>
  {/each}
{:else}
  <p class="loading">loading...</p>
{/if}
