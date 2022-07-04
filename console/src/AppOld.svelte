<script>
  import { Router, Route } from 'svelte-navigator'
  import Home from './routes/home/home.svelte'
  import SettingsOld from './routes/settings_old/settings.svelte'
  import TxHistory from './routes/tx_history/transactions.svelte'
  import Submit from './routes/submit.svelte'
  import GlobalStyle from './lib/styles/GlobalStyle.svelte'
  import Menu from './Menu.svelte'
  import Footer from './lib/components/footer/index.svelte'

  let isActive = false

  function toggleMenu() {
    isActive = !isActive
  }

  function closeMenu() {
    isActive = false
  }

  let routes = [
    { id: 0, path: '/', label: 'Home', component: Home },
    { id: 1, path: '/history', label: 'History', component: TxHistory },
    { id: 2, path: '/submit', label: 'Submit data', component: Submit },
    { id: 3, path: '/settings', label: 'Settings', component: SettingsOld },
  ]

  export let url = ''

  const src = 'taal-logo.png'
</script>

<svelte:head>
  <link rel="stylesheet" href="/css/bulma.min.css" />
  <link rel="stylesheet" href="/css/bulma-switch.min.css" />
</svelte:head>

<Router {url}>
  <GlobalStyle>
    <nav class="navbar is-fixed-top is-dark" aria-label="main navigation">
      <div class="navbar-brand">
        <span class="navbar-item has-text-light">
          <span class="is-size-4 client">TaalClient</span>
        </span>

        <a
          href="#/"
          role="button"
          class="navbar-burger {isActive ? 'is-active' : ''}"
          aria-label="menu"
          aria-expanded="false"
          data-target="navbarBasicExample"
          on:click={toggleMenu}
        >
          <span aria-hidden="true" />
          <span aria-hidden="true" />
          <span aria-hidden="true" />
        </a>
      </div>

      <div class="navbar-menu {isActive ? 'is-active' : ''}">
        <Menu {routes} on:close={closeMenu} />
      </div>
    </nav>

    <main class="section">
      {#each routes as item (item.id)}
        <Route
          path={item.path}
          component={item.component}
          meta={{ name: item.label }}
          primary={false}
        />
      {/each}
    </main>

    <Footer />
  </GlobalStyle>
</Router>

<style>
  main {
    padding-top: 20px;
  }
  .client {
    margin-left: 5px;
  }
  /**/
  .content {
  }
</style>
