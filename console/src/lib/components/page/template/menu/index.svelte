<script>
  import { useLocation, useNavigate } from 'svelte-navigator'
  import { menuLinks, menuActions } from '../../../../stores'
  import Header from '../../../header/index.svelte'
  import Footer from '../../../footer/index.svelte'
  import ContentBasic from '../../content/basic/index.svelte'

  const location = useLocation()
  const navigate = useNavigate()

  let links = []
  let actions = []

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
</script>

<Header
  showLinks={true}
  showActions={true}
  {links}
  {actions}
  on:link={onMenuItem}
  on:action={onMenuItem}
/>

<div class="content-container">
  <ContentBasic>
    <slot />
  </ContentBasic>

  <Footer />
</div>

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
