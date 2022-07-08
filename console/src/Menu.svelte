<script>
  import { createEventDispatcher } from 'svelte'
  import { Link, useLocation } from 'svelte-navigator'

  const dispatch = createEventDispatcher()

  function onClose() {
    dispatch('close')
  }

  const location = useLocation()

  export let routes = []

  let data = []

  $: {
    data = routes.map((route) => ({
      ...route,
      selected: route.path === $location.pathname,
    }))
  }
</script>

<div class="navbar-start is-dark" on:close>
  {#each data as item (item.id)}
    <Link
      to={item.path}
      class="navbar-item{item.selected ? ' is-active' : ''}"
      on:click={onClose}>{item.label}</Link
    >
  {/each}
</div>
