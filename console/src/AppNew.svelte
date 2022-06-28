<script>
  import { Router, Route } from 'svelte-navigator'
  import Home from './routes/home/home.svelte'
  import Settings from './routes/settings/settings.svelte'
  import TxHistory from './routes/tx_history/transactions.svelte'
  import Submit from './routes/submit.svelte'
  import GlobalStyle from './lib/styles/GlobalStyle.svelte'
  import NewUsers from './routes/new-users/index.svelte'

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
    { id: 3, path: '/settings', label: 'Settings', component: Settings },
  ]

  export let url = ''
</script>

<svelte:head>
  <link rel="preconnect" href="https://fonts.googleapis.com" />
  <link rel="preconnect" href="https://fonts.gstatic.com" crossOrigin="true" />
  <link
    href="https://fonts.googleapis.com/css2?family=Work+Sans:wght@400;500;700&display=swap"
    rel="stylesheet"
  />
</svelte:head>

<Router {url}>
  <GlobalStyle>
    <!-- <nav class="navbar is-fixed-top is-dark" aria-label="main navigation">
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

    <Footer /> -->
    <Route
      path="/"
      component={NewUsers}
      meta={{ name: 'New Users' }}
      primary={false}
    />
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
