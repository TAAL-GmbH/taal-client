<script>
  import { useLocation, useNavigate } from 'svelte-navigator'
  import { menuLinks, menuActions } from '../../../../stores'
  import { query } from '../../../../actions'
  import { medium, large } from '../../../../styles/breakpoints'
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

  $: isMedium = query(medium)
  $: isLarge = query(large)
  $: mediaSize = $isLarge ? 'large' : $isMedium ? 'medium' : 'small'

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
  {links}
  {actions}
  on:link={onMenuItem}
  on:action={onMenuItem}
  on:toggle-menu={onToggleMenu}
/>

<div class="content-container">
  <ContentMenu>
    <slot />
  </ContentMenu>

  <Footer />
</div>

{#if showSidebarMenu && mediaSize === 'small'}
  <Sidebar {links} on:link={onMenuItem} />
{/if}

<style>
  .content-container {
    position: absolute;
    top: 80px;
    width: 100%;
    height: calc(100% - 80px);
    overflow-x: hidden;
    overflow-y: auto;
  }
</style>
