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
  <nav
    class="navbar is-fixed-top has-background-dark"
    aria-label="main navigation"
  >
    <div class="navbar-brand">
      <a class="navbar-item has-text-light" href="https://taal.com">
        <img {src} alt="Taal logo" />
        <span class="is-size-4 client">Client</span>
      </a>

      <a
        href="#/"
        role="button"
        class="navbar-burger"
        aria-label="menu"
        aria-expanded="false"
        data-target="navbarBasicExample"
      >
        <span aria-hidden="true" />
        <span aria-hidden="true" />
        <span aria-hidden="true" />
      </a>
    </div>

    <div id="navbarBasicExample" class="navbar-menu">
      <div class="navbar-start">
        <Link to="/" class="navbar-item has-text-light">Home</Link>

        <Link to="/settings" class="navbar-item has-text-light">Settings</Link>
        <Link to="/transactions" class="navbar-item has-text-light">Transactions</Link>
      </div>
    </div>
  </nav>

  <main class="section">
    <Route path="/"><Home /></Route>
    <Route path="settings" component={Settings} />
    <Route path="transactions" component={Transactions} />
  </main>
</Router>

<style>
  .client {
    margin-left: 5px;
  }
</style>
