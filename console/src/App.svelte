<script>
  import { Router, Route, Link } from 'svelte-routing'
  import Home from './routes/home/home.svelte'
  import Settings from './routes/settings/settings.svelte'
  import TxHistory from './routes/tx_history/transactions.svelte'
  import Submit from './routes/submit.svelte'

  let isActive = false

  function toggleMenu() {
    isActive = !isActive
  }

  function closeMenu() {
    isActive = false
  }

  export let url = ''

  const src = 'taal-logo.png'
</script>

<svelte:head>
  <link rel="stylesheet" href="/css/bulma.min.css" />
  <link rel="stylesheet" href="/css/bulma-switch.min.css" />
</svelte:head>

<Router {url}>
  <nav class="navbar is-fixed-top is-dark" aria-label="main navigation">
    <div class="navbar-brand">
      <a class="navbar-item has-text-light" href="https://taal.com">
        <img {src} alt="Taal logo" />
        <span class="is-size-4 client">Client</span>
      </a>

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
      <div class="navbar-start is-dark">
        <Link to="/" class="navbar-item" on:click={closeMenu}>Home</Link>

        <Link to="/settings" class="navbar-item" on:click={closeMenu}>
          Settings
        </Link>

        <Link to="/history" class="navbar-item" on:click={closeMenu}>
          History
        </Link>

        <Link to="/submit" class="navbar-item" on:click={closeMenu}>
          Submit data
        </Link>
      </div>
    </div>
  </nav>

  <main class="section">
    <Route path="/"><Home /></Route>
    <Route path="settings" component={Settings} />
    <Route path="history" component={TxHistory} />
    <Route path="submit" component={Submit} />
  </main>
</Router>

<style>
  main {
    padding-top: 20px;
  }
  .client {
    margin-left: 5px;
  }
</style>
