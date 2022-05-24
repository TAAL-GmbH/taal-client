<script>
  // import { onMount } from "svelte";

  // let name;

  // onMount(async () => {
  //   name = "N/A";

  //   const response = await fetch(`${BASE_URL}/api/v1/test`);
  //   const data = await response.text();

  //   name = data;
  // });

  import { Router, Route, Link } from 'svelte-routing'
  import Home from './routes/home.svelte'
  import Settings from './routes/settings.svelte'
  import Transactions from './routes/transactions.svelte'
  import Submit from './routes/submit.svelte'
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
  <!-- <link rel="stylesheet" href="https://cdn.simplecss.org/simple.min.css" /> -->
  <link
    rel="stylesheet"
    href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"
  />
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
        <Link to="/submit" class="navbar-item" on:click={closeMenu}>Submit Transaction</Link
        >
      </div>
    </div>
  </nav>

  <main class="section">
    <Route path="/"><Home /></Route>
    <Route path="settings" component={Settings} />
    <Route path="transactions" component={Transactions} />
    <Notifications>
      <Route path="submit" component={Submit} />
    </Notifications>
   
  </main>
</Router>

<style>
  .client {
    margin-left: 5px;
  }
</style>
