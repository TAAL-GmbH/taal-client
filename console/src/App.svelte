<script>
  import { Router, Route, Link } from 'svelte-routing'
  import Home from './routes/home/home.svelte'
  import Settings from './routes/settings/settings.svelte'
  import TxHistory from './routes/tx_history/transactions.svelte'
  import TxSubmit from './routes/submit.svelte'
  import Example from './routes/example.svelte'
  import Notifications from 'svelte-notifications'

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

    <div
      id="navbarBasicExample"
      class="navbar-menu {isActive ? 'is-active' : ''}"
    >
      <div class="navbar-start is-dark">
        <Link to="/" class="navbar-item" on:click={closeMenu}>Home</Link>

        <Link to="/settings" class="navbar-item" on:click={closeMenu}
          >Settings</Link
        >
        <Link to="/transactions" class="navbar-item" on:click={closeMenu}
          >Transactions History</Link
        >
        <Link to="/submit" class="navbar-item" on:click={closeMenu}
          >Submit Transaction</Link
        >
        <Link to="/example" class="navbar-item" on:click={closeMenu}
          >Example</Link
        >
      </div>
    </div>
  </nav>

  <main class="section">
    <Route path="/"><Home /></Route>
    <Route path="settings" component={Settings} />
    <Notifications>
      <Route path="transactions" component={TxHistory} />
      <Route path="submit" component={TxSubmit} />
      <Route path="example" component={Example} />
    </Notifications>
  </main>
</Router>

<style>
  .client {
    margin-left: 5px;
  }
</style>
