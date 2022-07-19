<script>
  import { useLocation, useNavigate } from 'svelte-navigator'
  import {
    headerHeight,
    menuLinks,
    menuActions,
    mediaSize,
  } from '../../../../stores'
  import Header from '../../../header/index.svelte'
  import Footer from '../../../footer/index.svelte'
  import ContentMenu from '../../content/menu/index.svelte'
  import Sidebar from '../../../sidebar/index.svelte'

  const location = useLocation()
  const navigate = useNavigate()

  let links = []
  let actions = []
  let showSidebarMenu = false

  $: {
    links = $menuLinks.map((route) => ({
      ...route,
      selected: route.path === $location.pathname,
    }))
    actions = $menuActions.map((route) => ({
      ...route,
      selected: route.path === $location.pathname,
    }))
  }

  function onMenuItem(e) {
    navigate(e.detail.path)
  }

  function onToggleMenu(e) {
    showSidebarMenu = e.detail.open
  }
</script>

<Header
  showLinks={true}
  showActions={true}
  open={showSidebarMenu}
  {links}
  {actions}
  on:link={onMenuItem}
  on:action={onMenuItem}
  on:toggle-menu={onToggleMenu}
/>

<div class="content-container" style:--top-local={$headerHeight + 'px'}>
  <ContentMenu>
    <slot />
  </ContentMenu>

  <Footer />
</div>

{#if showSidebarMenu && $mediaSize === 'small'}
  <Sidebar
    {links}
    on:link={onMenuItem}
    on:close={() => (showSidebarMenu = false)}
  />
{/if}

<style>
  .content-container {
    position: absolute;
    top: var(--top-local);
    width: 100%;
    height: calc(100% - var(--top-local));
    overflow-x: hidden;
    overflow-y: auto;
  }
</style>
