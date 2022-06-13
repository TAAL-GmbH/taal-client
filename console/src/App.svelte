<script>
  import { Router, Route, Link } from 'svelte-routing'
  import Home from './routes/home/home.svelte'
  import Settings from './routes/settings/settings.svelte'
  import TxHistory from './routes/tx_history/transactions.svelte'
  import Submit from './routes/submit.svelte'

  let isActive = false

  const isActiveClass = 'navbar-item is-active'
  const isInactiveClass = 'navbar-item'

  const navBarItemClassesDictInit = {
    home: isInactiveClass,
    settings: isInactiveClass,
    history: isInactiveClass,
    submit: isInactiveClass,
  }

  let navBarItemClassesDict = {}

  function setActiveItem(event) {
    Object.assign(navBarItemClassesDict, navBarItemClassesDictInit)
    navBarItemClassesDict[event.detail.message] = isActiveClass
  }

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
      <div class="navbar-start is-dark">
        <Link
          to="/"
          bind:class={navBarItemClassesDict['home']}
          on:click={closeMenu}
          >Home
        </Link>

        <Link
          to="/history"
          bind:class={navBarItemClassesDict['history']}
          on:click={closeMenu}
        >
          History
        </Link>

        <Link
          to="/submit"
          bind:class={navBarItemClassesDict['submit']}
          on:click={closeMenu}
        >
          Submit data
        </Link>

        <Link
          to="/settings"
          bind:class={navBarItemClassesDict['settings']}
          on:click={closeMenu}
        >
          Settings
        </Link>
      </div>
    </div>
  </nav>

  <main class="section">
    <Route path="/"><Home on:mounted={setActiveItem} /></Route
    >
    <Route path="history"
      ><TxHistory on:mounted={setActiveItem} /></Route
    >
    <Route path="submit"
      ><Submit on:mounted={setActiveItem} /></Route
    >
    <Route path="settings"
      ><Settings on:mounted={setActiveItem} /></Route
    >
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
