<script>
  // svelte-ignore unused-export-let
  export let location

  import SubmitDev from './submit-dev.svelte'
  import SubmitSimple from './submit-simple.svelte'
  import Notifications from 'svelte-notifications'
  import { createEventDispatcher } from 'svelte'
  import { onMount } from 'svelte'

  const dispatch = createEventDispatcher()

  onMount(() => {
    dispatch('mounted', {})
  })

  let checked = localStorage.getItem('devmode') === 'true'

  $: localStorage.setItem('devmode', checked)
</script>

<Notifications>
  <div id="devmode" class="control is-pulled-right">
    <div class="field">
      <input
        id="devButton"
        type="checkbox"
        name="devButton"
        class="switch is-rtl is-info"
        bind:checked
      />
      <label id="devLabel" class="label" for="devButton">Developer mode</label>
    </div>
  </div>

  <section class="">
    {#if checked}
      <SubmitDev />
    {:else}
      <SubmitSimple />
    {/if}
  </section>
</Notifications>

<style>
  #devmode {
    padding-right: 25px;
    padding-top: 8px;
  }

  #devLabel {
    margin-bottom: 10px;
  }
</style>
